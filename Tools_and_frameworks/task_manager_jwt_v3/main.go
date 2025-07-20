package main

import (
	"task_man_v3/data"
	"task_man_v3/router"
)

func main() {
	// Initialize and run the user database
	data.InitializeDB()
	
	// Initialize and connect the user database
	data.InitializeUserDB()

	// Initialize and run the router
	router.RunRouter()
}