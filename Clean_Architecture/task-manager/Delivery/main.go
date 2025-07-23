package main

import (
	"task_manager_ca/Delivery/routers"
	repositories "task_manager_ca/Repositories"
)

func main() {
	// Initialize the user Database
	Usrdb := repositories.InitializeUserDB()
	
	// Initialize the Task Database
	Taskdb := repositories.InitializeDB()

	// Initialize and run the router
	routers.SetupRouter(Usrdb, Taskdb)
}