package routers

import (
	"task_manager_ca/Delivery/controllers"
	infrastructure "task_manager_ca/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(UsrCtrl controllers.UserController, TaskCtrl controllers.TaskController, middleware infrastructure.AuthMiddleWareI) {
	router := gin.Default()

	// Handle path routes
	router.GET("/tasks", TaskCtrl.GetAllTasks)
	router.GET("/tasks/:id", TaskCtrl.GetTaskByID)
	router.PUT("/tasks/:id", middleware.Validate_token(),TaskCtrl.UpdateTaskByID)
	router.POST("/tasks", middleware.Validate_token() , TaskCtrl.CreateTaskController)
	router.POST("/register", UsrCtrl.RegisterController)
	router.POST("/login", UsrCtrl.LoginController)
	router.GET("/admin_page", middleware.Validate_token(), middleware.Validate_role(), UsrCtrl.AdminPageController)
	router.DELETE("/tasks/:id", middleware.Validate_token(),TaskCtrl.DeleteTaskController) 
	router.GET("/user_profile", middleware.Validate_token(), UsrCtrl.UserProfileController)

	// Start router
	router.Run()
}