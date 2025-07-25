package Domain

import (
	"time"
)

type PasswordService interface {
	EncryptPassword(password string) ([]byte, error)
	CheckPasswordHash(password, hashedPassword string) (bool)
}

type UserRepository interface {
	Create(user User) error
	Login(user User) (User, error)
	FindUserRepository(string) (bool)
}

type TaskRepository interface {
	GetAllElements() ([]Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	CreateTask(new_task Task) error
	UpdateTaskByID(id string, updatedTask Task) error
	DeleteTask(id string) error
}


type JwtService interface {
	CreateJwtToken(User) (string, error)
}

type Task struct {
	ID string  
	Description string 
	Status string 
	DueDate time.Time 
	Title string 
}

type User struct {
	Name     string 
	Username string 
	ID       string 
	Role     string 
	Password string 
}
