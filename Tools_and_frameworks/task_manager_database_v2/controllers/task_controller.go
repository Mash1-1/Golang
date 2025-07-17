package controllers

import (
	"example/task_manager_database_v2/data"
	"example/task_manager_database_v2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTaskByIDController(c *gin.Context) {
	id := c.Param("id")
	task, err := data.GetTaskByIDService(id)

	if err != nil && err.Error() == "task not found" {
		// Handle task not found
		c.JSON(http.StatusNotFound, gin.H{"message" : err.Error()})
		return
	}
	if err != nil {
		// Handle database failure 
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Task" : task})
}

func GetTasksController(c *gin.Context) {
	tasks, err := data.GetTasksService() 
	// Added error handling incase of database failure
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Database failure"})
		return 
	}
	
	c.JSON(http.StatusOK, gin.H{"Tasks" : tasks})
}

func CreateTaskController(c *gin.Context) {
	var new_task models.Task 
	if err := c.ShouldBindJSON(&new_task); err != nil {
		// Handle invalid inputs or binding errors
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return 
	}
	err := data.CreateTaskService(new_task)
	if err != nil {
		// Handle database failure
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Database failure"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message" : "Task created!"})
}

func UpdateTaskByIDController(c *gin.Context) {
	id := c.Param("id")

	var updatedTask models.Task 
	
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		// Handle invalid inputs or binding errors
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return 
	}
	err := data.UpdateTaskByIDService(id, updatedTask)
	if err != nil && err.Error() == "task not found" {
		// Handle Task not found
		c.JSON(http.StatusNotFound, gin.H{"message" : err.Error()})
		return
	}
	if err != nil {
		// Handle database failure
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Database failure"})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{"message" : "Task updated Successfully!"})
}

func DeleteTaskController(c *gin.Context) {
	id := c.Param("id")

	err := data.DeleteTaskService(id)
	if err != nil && err.Error() == "task not found" {
		// Handle task not found
		c.JSON(http.StatusNotFound, gin.H{"message" : "Task not found!"})
		return 
	}
	if err != nil {
		// Handle database failure
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Database failure"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Task deleted!"})
}