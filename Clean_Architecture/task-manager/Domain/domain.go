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
	ID string  `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
	DueDate time.Time `json:"due_date"`
	Title string `json:"title"`
}

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	ID       string `json:"id"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
