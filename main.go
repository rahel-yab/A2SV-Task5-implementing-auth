package main

import (
	"task-management-rest-api/data"
	"task-management-rest-api/router"
)

func main(){
	mongoClient := data.ConnectMongo()
	router := router.SetupRouter(mongoClient)
	router.Run()
}

