package controllers

import (
	"net/http"
	"task_man_v3/data"
	"task_man_v3/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Initialize JWT secret for signing the jwt token
var Jwt_secret = []byte("golang_is_amazing") 

func UserRegisterController(c *gin.Context) {
	var new_user models.User
	// Get user information from the input
	if err := c.ShouldBindJSON(&new_user); err != nil {
		// Handle binding errors
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return 
	}

	// Check if the username is unique
	if data.FindUserService(new_user.Username) {
		// If username already in use
		c.JSON(http.StatusConflict, gin.H{"error" : "username already exists"})
		return 
	}

	// Encrypt password using bcrypt
	hashedPassword, err:= bcrypt.GenerateFromPassword([]byte(new_user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "error hashing using bcrypt"})
	}
	// Insert the new_user information into the database
	err = data.UserRegisterService(new_user, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "database insertion error"})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"message" : "user registered successfully"})
}

func GetUserProfileController(c *gin.Context) {
	c.JSON(200, gin.H{"message" : "Only logged in users can see this"})
}

func GetAdminPageController(c *gin.Context) {
	c.JSON(200, gin.H{"message" : "Hello, welcome to the admin page!"})
}

func UserLoginController(c *gin.Context) {
	var user models.User 

	// Accept user information in json format from user & handle binding error
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	// Check if user is in the database 
	existingUser, err := data.UserLoginService(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "Invalid email or password"})
		return 
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role" : existingUser.Role,
		"username" : existingUser.Username,
	})

	// Generate JWT token 
	jwtToken, err := token.SignedString(Jwt_secret)

	// Handle error while signing 
	if err != nil {
		c.JSON(500, gin.H{"error" : "Internal server error"})
		return 
	}

	// Give the token to the client for the next session use
	c.JSON(http.StatusOK, gin.H{"message" : "user logged in successfully", "token" : jwtToken})
}

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