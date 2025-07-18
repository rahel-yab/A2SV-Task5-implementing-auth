package controllers

import (
	"net/http"
	"task-management-rest-api/data"
	"task-management-rest-api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func GetTasks(c *gin.Context, client *mongo.Client) {
	tasks, err := data.GetAllTasks(c.Request.Context(), client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context, client *mongo.Client) {
	id := c.Param("id")
	task, err := data.GetTaskByID(c.Request.Context(), client, id)
	if err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, task)
}

func RemoveTask(c *gin.Context, client *mongo.Client) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(403, gin.H{"error": "Admin access required"})
		return
	}
	jwtClaims, ok := claims.(jwt.MapClaims)
	role, ok := jwtClaims["role"].(string)
	if !ok || role != "admin" {
		c.JSON(403, gin.H{"error": "Admin access required"})
		return
	}
	id := c.Param("id")
	err := data.RemoveTask(c.Request.Context(), client, id)
	if err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task removed"})
}

func UpdateTask(c *gin.Context, client *mongo.Client) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(403, gin.H{"error": "Admin access required"})
		return
	}
	jwtClaims, ok := claims.(jwt.MapClaims)
	role, ok := jwtClaims["role"].(string)
	if !ok || role != "admin" {
		c.JSON(403, gin.H{"error": "Admin access required"})
		return
	}
	var updatedTask models.Task
	id := c.Param("id")

	err := c.BindJSON(&updatedTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = data.UpdateTask(c.Request.Context(), client, id, updatedTask)
	if err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func AddTask(c *gin.Context, client *mongo.Client) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(403, gin.H{"error": "Admin access required"})
		return
	}
	jwtClaims, ok := claims.(jwt.MapClaims)
	role, ok := jwtClaims["role"].(string)
	if !ok || role != "admin" {
		c.JSON(403, gin.H{"error": "Admin access required"})
		return
	}
	var newTask models.Task
	err := c.BindJSON(&newTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = data.AddTask(c.Request.Context(), client, newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

func RegisterUser(c *gin.Context, client *mongo.Client) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	emailExists, err := data.UserExistsByEmail(c.Request.Context(), client, req.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	if emailExists {
		c.JSON(400, gin.H{"error": "Email already registered"})
		return
	}

	usernameExists, err := data.UserExistsByUsername(c.Request.Context(), client, req.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	if usernameExists {
		c.JSON(400, gin.H{"error": "Username already taken"})
		return
	}

	isEmpty, err := data.IsUsersCollectionEmpty(c.Request.Context(), client)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	role := "user"
	if isEmpty {
		role = "admin"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     role,
	}

	if err := data.AddUser(c.Request.Context(), client, user); err != nil {
		c.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully", "role": role})
}

func LoginUser(c *gin.Context, client *mongo.Client) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	var user *models.User
	var err error

	if req.Email != "" {
		user, err = data.GetUserByEmail(c.Request.Context(), client, req.Email)
	} else if req.Username != "" {
		user, err = data.GetUserByUsername(c.Request.Context(), client, req.Username)
	} else {
		c.JSON(400, gin.H{"error": "Email or username is required"})
		return
	}

	if err != nil || user == nil {
		c.JSON(401, gin.H{"error": "Invalid email/username or password"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		c.JSON(401, gin.H{"error": "Invalid email/username or password"})
		return
	}

	jwtSecret := []byte("your_jwt_secret") // TODO: move to config/env
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role,
	})

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{
		"message": "User logged in successfully",
		"token":   jwtToken,
		"role":    user.Role,
	})
}

func PromoteUser(c *gin.Context, client *mongo.Client) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	jwtClaims := claims.(jwt.MapClaims)
	role, ok := jwtClaims["role"].(string)
	if !ok || role != "admin" {
		c.JSON(403, gin.H{"error": "Admin access required"})
		return
	}

	var req struct {
		Identifier string `json:"identifier"` // username or email
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Identifier == "" {
		c.JSON(400, gin.H{"error": "Username or email is required"})
		return
	}

	err := data.PromoteUserToAdmin(c.Request.Context(), client, req.Identifier)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Failed to promote user"})
		return
	}

	c.JSON(200, gin.H{"message": "User promoted to admin"})
}