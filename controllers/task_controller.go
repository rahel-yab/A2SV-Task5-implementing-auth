package controllers

import (
	"net/http"
	"task-management-rest-api/data"
	"task-management-rest-api/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context){
c.JSON(http.StatusOK , data.Tasks)
}

func GetTask(c *gin.Context){
	id := c.Param("id")
	for _ , t := range data.Tasks{
		if t.ID == id{
			c.JSON(http.StatusOK , t)
			return
		}
	}
	c.JSON(http.StatusNotFound , gin.H{"message":"task not found"})
}

func RemoveTask(c *gin.Context){
	id := c.Param("id")

	for i , t := range data.Tasks{
		if t.ID == id{
			data.Tasks = append(data.Tasks[:i],data.Tasks[i+1:]... )
			c.JSON(http.StatusOK , gin.H{"message": "Task removed"})
			return
		}
	}
	c.JSON(http.StatusNotFound , gin.H{"message":"task not found"})
}


func UpdateTask(c *gin.Context){
	var updatedTask models.Task
	id := c.Param("id")

	err := c.BindJSON(&updatedTask)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	for i , t := range data.Tasks{
		if t.ID == id{
			data.Tasks[i] = updatedTask
		c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
            return
}}}


func AddTask(c *gin.Context){
	var newTask models.Task
	err := c.BindJSON(&newTask)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	data.Tasks = append(data.Tasks, newTask)
	c.JSON(http.StatusCreated , gin.H{"message": "Task created"})
}