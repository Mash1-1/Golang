package data

import (
	"errors"
	"example/task_manager/models"
	"time"
)

// Initialize sample tasks for testing
var Tasks = []models.Task{
    {ID: "1", Status: "In progress", Title: "Task 4", Description: "Backend with Go", DueDate: time.Now().AddDate(0, 0, 5)},
    {ID: "2", Status: "Pending", Title: "Design API Endpoints", Description: "Plan REST API structure and routes for task management", DueDate: time.Now().AddDate(0, 0, 2)},
    {ID: "3", Status: "Completed", Title: "Setup Database", Description: "Initialize PostgreSQL and create necessary tables", DueDate: time.Now().AddDate(0, 0, -3)},
    {ID: "4", Status: "In progress", Title: "Write Unit Tests", Description: "Cover task controller with unit tests using testify", DueDate: time.Now().AddDate(0, 0, 7)},
    {ID: "5", Status: "Pending", Title: "Dockerize App", Description: "Create Dockerfile and docker-compose for development", DueDate: time.Now().AddDate(0, 0, 10)},
}


func GetTasksService() []models.Task {
	return Tasks
}

func GetTaskByIDService(id string) (models.Task, error) {
	for _, task := range Tasks {
		if task.ID == id {
			return task, nil
		}
	}
	// If task isn't found send an error
	return models.Task{}, errors.New("task not found")
}

func CreateTaskService(new_task models.Task) {
	// Add the new task to the database
	Tasks = append(Tasks, new_task)
}

func UpdateTaskByIDService(id string, updatedTask models.Task) error {
	for ind, task := range Tasks {
		if task.ID == id {
			// Update fields that have new values sent as a request
			if updatedTask.Description != "" {
				Tasks[ind].Description = updatedTask.Description
			}
			if updatedTask.Status != "" {
				Tasks[ind].Status = updatedTask.Status
			}
			if updatedTask.Title != "" {
				Tasks[ind].Title = updatedTask.Title
			}
			return nil
		}
	} 
	// Tell controller that task wasn't found
	return errors.New("task not found")
}

func DeleteTaskService(id string) error {
	for ind, task := range Tasks{
		if task.ID == id {
			// Delete the task by slicing
			Tasks = append(Tasks[:ind], Tasks[ind+1:]...)
			return nil
		}
	}
	// Tell controller that task wasn't found
	return errors.New("task not found")
}