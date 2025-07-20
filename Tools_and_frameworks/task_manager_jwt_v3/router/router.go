package router

import (
	"task_man_v3/controllers"
	"task_man_v3/middleware"

	"github.com/gin-gonic/gin"
)


func RunRouter() {
	router := gin.Default()

	// Handle path routes
	router.GET("/tasks", controllers.GetTasksController)
	router.GET("/tasks/:id", controllers.GetTaskByIDController)
	router.PUT("/tasks/:id", middleware.AuthMiddleWare(),controllers.UpdateTaskByIDController)
	router.POST("/tasks", middleware.AuthMiddleWare() ,controllers.CreateTaskController)
	router.POST("/register", controllers.UserRegisterController)
	router.POST("/login", controllers.UserLoginController)
	router.GET("/admin_page", middleware.AuthMiddleWare(), middleware.AuthRoleMiddleWare(), controllers.GetAdminPageController)
	router.DELETE("/tasks/:id", middleware.AuthMiddleWare(),controllers.DeleteTaskController) 
	router.GET("/user_profile", middleware.AuthMiddleWare(), controllers.GetUserProfileController)

	// Start router
	router.Run()
}