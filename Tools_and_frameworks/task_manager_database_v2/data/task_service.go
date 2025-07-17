package data

import (
	"context"
	"errors"
	"example/task_manager_database_v2/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Collection 

func InitializeDB() error{
	// Initialize collection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err 
	}
	collection := client.Database("Task_DB").Collection("tasks")
	// Clear previous usage leftover data
	collection.DeleteMany(context.TODO(), bson.D{{}})

	// Initialize sample tasks for testing
	var Tasks = []interface{}{
    models.Task{ID: "1", Status: "In progress", Title: "Task 4", Description: "Backend with Go", DueDate: time.Now().AddDate(0, 0, 5)},
    models.Task{ID: "2", Status: "Pending", Title: "Design API Endpoints", Description: "Plan REST API structure and routes for task management", DueDate: time.Now().AddDate(0, 0, 2)},
    models.Task{ID: "3", Status: "Completed", Title: "Setup Database", Description: "Initialize PostgreSQL and create necessary tables", DueDate: time.Now().AddDate(0, 0, -3)},
    models.Task{ID: "4", Status: "In progress", Title: "Write Unit Tests", Description: "Cover task controller with unit tests using testify", DueDate: time.Now().AddDate(0, 0, 7)},
    models.Task{ID: "5", Status: "Pending", Title: "Dockerize App", Description: "Create Dockerfile and docker-compose for development", DueDate: time.Now().AddDate(0, 0, 10)},
}
	// Insert sample data into the database
	collection.InsertMany(context.TODO(), Tasks)
	database = collection
	return nil
}

func DisconnectDB() error {
	err := database.Database().Client().Disconnect(context.TODO())
	return err 
}

func getElementsInDB() ([]models.Task, error) {
	// Gets all elements in the database using the find method and returns a tasks slice 
	var result []models.Task
	cursor, err := database.Find(context.TODO(), bson.D{{}})

	if err != nil {
		return result, err 
	}
	for cursor.Next(context.TODO()) {
		// Write the elements in the database into structs one by one
		var elem models.Task 
		err = cursor.Decode(&elem)
		//add the elements into the result
		result = append(result, elem)
	}
	if cursor.Err() != nil {
		return []models.Task{}, err 
	}
	return result, nil
}

func GetTasksService() ([]models.Task, error) {
	return getElementsInDB()
}

func GetTaskByIDService(id string) (models.Task, error) {
	Tasks, err := getElementsInDB()
	if err != nil {
		// Incase of database failure
		return models.Task{}, err 
	}
	for _, task := range Tasks {
		if task.ID == id {
			return task, nil
		}
	}
	// If task isn't found send an error
	return models.Task{}, errors.New("task not found")
}

func CreateTaskService(new_task models.Task) error {
	// Add the new task to the database
	_, err := database.InsertOne(context.TODO(), new_task)
	return err 
}

func UpdateTaskByIDService(id string, updatedTask models.Task) error {
	filter := bson.D{{Key: "id", Value: id}}
	updateBSON := bson.M{}

	if updatedTask.Description != "" {
		updateBSON["description"] = updatedTask.Description
	}
	if updatedTask.Status != "" {
		updateBSON["status"] = updatedTask.Status
	}
	if updatedTask.Title != "" {
		updateBSON["title"] = updatedTask.Title
	}

	update := bson.M{"$set" : updateBSON}
	updateResult, err := database.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err 
	}
	if updateResult.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

func DeleteTaskService(id string) error {
	filter := bson.D{{Key: "id", Value :id}}
	deleteResult, err := database.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 0 {
		// Tell controller that task wasn't found
		return errors.New("task not found")
	}
	return nil 
}