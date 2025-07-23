package controllers

import (
	"net/http"
	"task_manager_ca/Domain"
	usecases "task_manager_ca/Usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase usecases.UserUseCase 
}

type TaskController struct {
	TaskUseCase usecases.TaskUseCase
}

func NewUserController(uc *usecases.UserUseCase) (UserController) {
	return UserController{
		UserUseCase: *uc,
	}
}

func NewTaskController(tc *usecases.TaskUseCase) (TaskController) {
	return TaskController{
		TaskUseCase: *tc,
	}
}

func (UsrCtrl *UserController) RegisterController(c *gin.Context) {
	var new_user Domain.User
	// Get user information from the input
	if err := c.ShouldBindJSON(&new_user); err != nil {
		// Handle binding errors
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return 
	}

	// Call the usecase for user registration
	err := UsrCtrl.UserUseCase.Register(new_user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message" : "user registered successfully"})
}

func (UsrCtrl *UserController) UserProfileController(c *gin.Context) {
	c.JSON(200, gin.H{"message" : "Only logged in users can see this"})
}

func (UsrCtrl *UserController) AdminPageController(c *gin.Context) {
	c.JSON(200, gin.H{"message" : "Hello, welcome to the admin page!"})
}

func (UsrCtrl *UserController) LoginController(c *gin.Context) {
	var user Domain.User 

	// Accept user information in json format from user & handle binding error
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	// Call the UserUseCase for login
	token, err := UsrCtrl.UserUseCase.Login(user)

	if err != nil {
		// Check the type of error and generate a valid response
		if err.Error() == "error while generating jwt token" {
			c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
			return 
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "invalid username or password"})
		return 
	}

	// Give token for the user to use for upcoming sessions.
	c.JSON(http.StatusOK, gin.H{"message" : "Logged in successfully!", "token" : token})
}

func (TaskCtrl *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	task, err := TaskCtrl.TaskUseCase.GetElementByID(id)

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

func (TaskCtrl *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := TaskCtrl.TaskUseCase.GetAllElements()
	// Added error handling incase of database failure
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Database failure"})
		return 
	}
	
	c.JSON(http.StatusOK, gin.H{"Tasks" : tasks})
}

func (TaskCtrl *TaskController) CreateTaskController(c *gin.Context) {
	var new_task Domain.Task 
	if err := c.ShouldBindJSON(&new_task); err != nil {
		// Handle invalid inputs or binding errors
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return 
	}
	err := TaskCtrl.TaskUseCase.CreateTask(new_task)
	if err != nil {
		// Handle database failure
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Database failure"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message" : "Task created!"})
}

func (TaskCtrl *TaskController) UpdateTaskByID(c *gin.Context) {
	id := c.Param("id")

	var updatedTask Domain.Task 
	
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		// Handle invalid inputs or binding errors
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return 
	}

	err := TaskCtrl.TaskUseCase.UpdateTask(id, updatedTask)
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

func (TaskCtrl *TaskController) DeleteTaskController(c *gin.Context) {
	id := c.Param("id")

	err := TaskCtrl.TaskUseCase.DeleteTask(id)
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