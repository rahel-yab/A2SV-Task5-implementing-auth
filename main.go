package main

import (
	"task-management-rest-api/router"
)

func main(){
	router := router.SetupRouter()
	router.Run()
}

