package router

import (
	"task-management-rest-api/controllers"

	"github.com/gin-gonic/gin"
)
	
func SetupRouter() *gin.Engine{
	router := gin.Default()

	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTask)
	router.DELETE("/tasks/:id", controllers.RemoveTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.POST("/tasks", controllers.AddTask)
	return router
}