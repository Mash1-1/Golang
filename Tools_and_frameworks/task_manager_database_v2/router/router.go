package router

import (
	"example/task_manager_database_v2/controllers"

	"github.com/gin-gonic/gin"
)


func RunRouter() {
	router := gin.Default()

	// Handle path routes
	router.GET("/tasks", controllers.GetTasksController)
	router.GET("/tasks/:id", controllers.GetTaskByIDController)
	router.PUT("/tasks/:id", controllers.UpdateTaskByIDController)
	router.POST("/tasks", controllers.CreateTaskController)
	router.DELETE("/tasks/:id", controllers.DeleteTaskController) 

	// Start router
	router.Run()
}