package controllers

import (
	"net/http"
	"task-management-rest-api/data"
	"task-management-rest-api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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