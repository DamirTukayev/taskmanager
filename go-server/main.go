package main

import (
	"tasksmanager/database"
	"tasksmanager/router"
)

func main() {
	database.ConnectDatabase()
	r := router.SetupRouter()
	r.Run(":8080")
}
