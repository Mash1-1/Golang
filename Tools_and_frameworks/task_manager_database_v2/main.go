package main

import (
	"example/task_manager_database_v2/data"
	"example/task_manager_database_v2/router"
)

func main() {
	// Initialize and run the database
	data.InitializeDB()
	// Initialize and run the router
	router.RunRouter()
}