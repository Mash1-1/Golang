package models

import (
	"time"
)

type Task struct {
	ID string  `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
	DueDate time.Time `json:"due_date"`
	Title string `json:"title"`
}
