package routers

import (
	"task_manager_ca/Delivery/controllers"
	infrastructure "task_manager_ca/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(UsrCtrl controllers.UserController, TaskCtrl controllers.TaskController) {
	router := gin.Default()

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