package routers

import (
	"task_manager_ca/Delivery/controllers"
	infrastructure "task_manager_ca/Infrastructure"
	repositories "task_manager_ca/Repositories"
	usecases "task_manager_ca/Usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(Usrdb , Taskdb *mongo.Collection) {
	router := gin.Default()

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

	// Handle path routes
	router.GET("/tasks", TaskCtrl.GetAllTasks)
	router.GET("/tasks/:id", TaskCtrl.GetTaskByID)
	router.PUT("/tasks/:id", infrastructure.AuthMiddleWare(),TaskCtrl.UpdateTaskByID)
	router.POST("/tasks", infrastructure.AuthMiddleWare() , TaskCtrl.CreateTaskController)
	router.POST("/register", UsrCtrl.RegisterController)
	router.POST("/login", UsrCtrl.LoginController)
	router.GET("/admin_page", infrastructure.AuthMiddleWare(), infrastructure.AuthRoleMiddleWare(), UsrCtrl.AdminPageController)
	router.DELETE("/tasks/:id", infrastructure.AuthMiddleWare(),TaskCtrl.DeleteTaskController) 
	router.GET("/user_profile", infrastructure.AuthMiddleWare(), UsrCtrl.UserProfileController)

	// Start router
	router.Run()
}