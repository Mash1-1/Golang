package models

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	ID       string `json:"id"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
