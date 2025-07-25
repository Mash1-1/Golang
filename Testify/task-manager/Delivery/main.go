package main

import (
	"task_manager_ca/Delivery/controllers"
	"task_manager_ca/Delivery/routers"
	infrastructure "task_manager_ca/Infrastructure"
	repositories "task_manager_ca/Repositories"
	usecases "task_manager_ca/Usecases"
)

func main() {
	// Initialize the user Database
	Usrdb := repositories.InitializeUserDB()
	
	// Initialize the Task Database
	Taskdb := repositories.InitializeDB()

	// Create the TaskRepository
	tr := repositories.NewTaskRepo(Taskdb) 

	// Create the TaskUsecase
	tc := usecases.NewTaskUseCase(&tr)

	// Create the Task Controller
	TaskCtrl := controllers.NewTaskController(&tc)

	// Create the UserRepository and Password Service for the UserUseCase
	ur := repositories.NewUserRepo(Usrdb) 
	ps := infrastructure.BcryptPasswordService{}
	jserv := infrastructure.JwtImplementation{}

	// Create the UserUsecase for the User Controller
	uc := usecases.NewUserUseCase(&ur, ps, jserv)

	// Create the UserController for the routers to use
	UsrCtrl := controllers.NewUserController(&uc)

	// Initialize and run the router
	routers.SetupRouter(UsrCtrl, TaskCtrl)
}