package router

import (
	"task-management-rest-api/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)
	
func SetupRouter(mongoClient *mongo.Client) *gin.Engine{
	router := gin.Default()

	router.GET("/tasks", func(c *gin.Context) { controllers.GetTasks(c, mongoClient) })
	router.GET("/tasks/:id", func(c *gin.Context) { controllers.GetTask(c, mongoClient) })
	router.DELETE("/tasks/:id", func(c *gin.Context) { controllers.RemoveTask(c, mongoClient) })
	router.PUT("/tasks/:id", func(c *gin.Context) { controllers.UpdateTask(c, mongoClient) })
	router.POST("/tasks", func(c *gin.Context) { controllers.AddTask(c, mongoClient) })
	return router
}