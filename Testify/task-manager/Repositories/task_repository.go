package repositories

import (
	"context"
	"errors"
	"task_manager_ca/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepo struct {
	TaskDatabase *mongo.Collection
}

func NewTaskRepo(Tskdb *mongo.Collection) (TaskRepo){
	return TaskRepo{
		TaskDatabase: Tskdb,
	}
}

func InitializeDB() (*mongo.Collection) {
	// Initialize collection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return &mongo.Collection{}
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return &mongo.Collection{}
	}
	collection := client.Database("Task_DB").Collection("tasks")
	// Clear previous usage leftover data
	collection.DeleteMany(context.TODO(), bson.D{{}})

	return collection
}

func (tr *TaskRepo) GetAllElements() ([]Domain.Task, error) {
	// Gets all elements in the Database using the find method and returns a tasks slice 
	var result []Domain.Task
	cursor, err := tr.TaskDatabase.Find(context.TODO(), bson.D{{}})

	if err != nil {
		return result, err 
	}
	for cursor.Next(context.TODO()) {
		// Write the elements in the Database into structs one by one
		var elem Domain.Task 
		err = cursor.Decode(&elem)
		//add the elements into the result
		result = append(result, elem)
	}
	if cursor.Err() != nil {
		return []Domain.Task{}, err 
	}
	return result, nil
}

func (tr *TaskRepo) GetAllTasks() ([]Domain.Task, error) {
	return tr.GetAllElements()
}

func (tr *TaskRepo) GetTaskByID(id string) (Domain.Task, error) {
	Tasks, err := tr.GetAllElements()
	if err != nil {
		// Incase of Database failure
		return Domain.Task{}, err 
	}
	for _, task := range Tasks {
		if task.ID == id {
			return task, nil
		}
	}
	// If task isn't found send an error
	return Domain.Task{}, errors.New("task not found")
}

func (tr *TaskRepo) CreateTask(new_task Domain.Task) error {
	// Add the new task to the Database
	tmp := Domain.Task{}
	if new_task != tmp{
		_, err := tr.TaskDatabase.InsertOne(context.TODO(), new_task)
		return err
	}
	return errors.New("empty task")
}

func (tr *TaskRepo) UpdateTaskByID(id string, updatedTask Domain.Task) error {
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
	updateResult, err := tr.TaskDatabase.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err 
	}
	if updateResult.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (tr *TaskRepo) DeleteTask(id string) error {
	filter := bson.D{{Key: "id", Value :id}}
	deleteResult, err := tr.TaskDatabase.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 0 {
		// Tell controller that task wasn't found
		return errors.New("task not found")
	}
	return nil 
}