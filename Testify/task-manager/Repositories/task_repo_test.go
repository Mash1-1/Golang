package repositories

import (
	"context"
	"task_manager_ca/Domain"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepositorySuite struct {
	suite.Suite 
	TskRepo TaskRepo 
	db *mongo.Collection
}

// Initialize sample tasks for testing
var Tasks = []interface{}{
	Domain.Task{ID: "1", Status: "In progress", Title: "Task 4", Description: "Backend with Go", DueDate: time.Now().AddDate(0, 0, 5)},
	Domain.Task{ID: "2", Status: "Pending", Title: "Design API Endpoints", Description: "Plan REST API structure and routes for task management", DueDate: time.Now().AddDate(0, 0, 2)},
	Domain.Task{ID: "3", Status: "Completed", Title: "Setup Database", Description: "Initialize PostgreSQL and create necessary tables", DueDate: time.Now().AddDate(0, 0, -3)},
	Domain.Task{ID: "4", Status: "In progress", Title: "Write Unit Tests", Description: "Cover task controller with unit tests using testify", DueDate: time.Now().AddDate(0, 0, 7)},
	Domain.Task{ID: "5", Status: "Pending", Title: "Dockerize App", Description: "Create Dockerfile and docker-compose for development", DueDate: time.Now().AddDate(0, 0, 10)},
}

func (suite *TaskRepositorySuite) SetupSuite() {
	// Initialize Database 
	db := InitializeDB()

	// Initialize task Repo
	tr := NewTaskRepo(db)

	// Insert values into suite fields
	suite.TskRepo = tr 
	suite.db = db 
}

func (suite *TaskRepositorySuite) TearDownTest() {
	// Clear database after each test to ensure test independence 
	suite.db.Database().Collection("tasks").Drop(context.TODO())
}

func (suite *TaskRepositorySuite) TestGetAllElements_Positive() {
	// Get all tasks returns the value from get all elements so this test also tests get all tasks implicitly
	// Insert sample data into the Database
	suite.db.InsertMany(context.TODO(), Tasks)
	
	// Check function 
	result, err := suite.TskRepo.GetAllElements() 

	// Handle error while reading from db
	suite.NoError(err, "error while reading from database")

	//Check if the result has equal length to pre inserted tasks in InitializeDB() 
	suite.Equal(len(Tasks), len(result), "didn't find all elements from result")
}

func (suite *TaskRepositorySuite) TestGetTaskByID_Positive() {
	// Add Tasks into the database for testing
	suite.db.InsertMany(context.TODO(), Tasks)

	// take a test task from the inserted tasks
	test_task := Tasks[0].(Domain.Task)

	// Check function
	result, err := suite.TskRepo.GetTaskByID(test_task.ID)

	// Check that err is nil
	suite.NoError(err, "error while getting task from db")

	// Check that the returned task is the same as the task whose id was sent
	suite.Equal(test_task.ID, result.ID, "wrong task returned")
}
// Negative for get task by ID - Non existing id -> "task not found"

func (suite *TaskRepositorySuite) TestGetTaskByID_NonExistingID_Negative() {
	// Add Tasks into the database for testing
	suite.db.InsertMany(context.TODO(), Tasks)
	non_existing_id := "100"

	// Check function
	_, err := suite.TskRepo.GetTaskByID(non_existing_id)

	// Check that err is not nil
	suite.Error(err, "error expected when getting a non existing task task from db")
}

func (suite *TaskRepositorySuite) TestCreateTask_Positive() {
	// Get a test task to test with
	test_task := Tasks[0].(Domain.Task)

	// Check Function
	err := suite.TskRepo.CreateTask(test_task)

	// Assert no error
	suite.NoError(err, "could not create task properly")
}
// Negative for create task, nil task
func (suite *TaskRepositorySuite) TestCreateTask_NilTask_Negative() {
	// Get a test task to test with
	invalid_task := Domain.Task{}

	// Check Function
	err := suite.TskRepo.CreateTask(invalid_task)

	// Assert error not equal to nil
	suite.Error(err, "expected error when creating nil task")
}

func (suite *TaskRepositorySuite) TestUpdateTaskByID_Positive() {
	// Get a test task for testing
	test_task := Tasks[0].(Domain.Task)

	// Add task into database and assertion
	_, err := suite.db.InsertOne(context.TODO(), test_task)
	suite.NoError(err, "error while adding into database")
	
	// Check function
	test_task.Description = "done"
	err = suite.TskRepo.UpdateTaskByID(test_task.ID, test_task)
	suite.NoError(err, "error while updating task")

	// Get task and check equality
	result, err := suite.TskRepo.GetTaskByID(test_task.ID)
	suite.NoError(err, "error while getting task to check update")
	suite.Equal("done", result.Description, "task not updated properly")
}
// Negative for update task by id - Non existent id -> task not found 

func (suite *TaskRepositorySuite) TestDeleteTaskByID_Positive() {
	test_task := Tasks[0].(Domain.Task)
	_, err := suite.db.InsertOne(context.TODO(), test_task)
	suite.NoError(err, "error while inserting into database")

	err = suite.TskRepo.DeleteTask(test_task.ID)
	suite.NoError(err, "no error expected when deleting with valid id")

	// Check for the task in database
	_, err = suite.TskRepo.GetTaskByID(test_task.ID)
	suite.NotNil(err, "error expected when trying to find element in database after deletion")
}
// Negative for DeleteTaskByID - Non existent id -> task not found

func TestTaskRepositorySuite(t *testing.T) {
	// Run the suite 
	suite.Run(t, new(TaskRepositorySuite))
}