package router

import (
	"task-management-rest-api/controllers"
	"task-management-rest-api/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(mongoClient *mongo.Client) *gin.Engine{
	router := gin.Default()

	jwtSecret := []byte("your_jwt_secret") // TODO: move to config/env

	taskGroup := router.Group("/tasks", middleware.AuthMiddleware(jwtSecret))
	{
		taskGroup.GET("", func(c *gin.Context) { controllers.GetTasks(c, mongoClient) })
		taskGroup.GET(":id", func(c *gin.Context) { controllers.GetTask(c, mongoClient) })
		taskGroup.DELETE(":id", func(c *gin.Context) { controllers.RemoveTask(c, mongoClient) })
		taskGroup.PUT(":id", func(c *gin.Context) { controllers.UpdateTask(c, mongoClient) })
		taskGroup.POST("", func(c *gin.Context) { controllers.AddTask(c, mongoClient) })
	}

	router.POST("/register", func(c *gin.Context) { controllers.RegisterUser(c, mongoClient) })
	router.POST("/login", func(c *gin.Context) { controllers.LoginUser(c, mongoClient) })

	// Protected route for promoting users
	router.POST("/promote", middleware.AuthMiddleware(jwtSecret), func(c *gin.Context) {
		controllers.PromoteUser(c, mongoClient)
	})

	return router
}