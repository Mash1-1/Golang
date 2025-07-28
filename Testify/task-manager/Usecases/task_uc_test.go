package usecases

import (
	"errors"
	"task_manager_ca/Domain"
	"task_manager_ca/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TaskUseCaseSuite struct {
	suite.Suite 
	TskUseCase TaskUseCase 
	repository *mocks.TaskRepository
}

// Initialize slice of tasks for testing
var Tasks = []Domain.Task{
	{ID: "1", Status: "In progress", Title: "Task 4", Description: "Backend with Go", DueDate: time.Now().AddDate(0, 0, 5)},
	{ID: "2", Status: "Pending", Title: "Design API Endpoints", Description: "Plan REST API structure and routes for task management", DueDate: time.Now().AddDate(0, 0, 2)},
	{ID: "3", Status: "Completed", Title: "Setup Database", Description: "Initialize PostgreSQL and create necessary tables", DueDate: time.Now().AddDate(0, 0, -3)},
	{ID: "4", Status: "In progress", Title: "Write Unit Tests", Description: "Cover task controller with unit tests using testify", DueDate: time.Now().AddDate(0, 0, 7)},
	{ID: "5", Status: "Pending", Title: "Dockerize App", Description: "Create Dockerfile and docker-compose for development", DueDate: time.Now().AddDate(0, 0, 10)},
}


// Initialize a task for testing
var test_task = Domain.Task{
	ID: "1", Status: "In progress", Title: "Task 4", Description: "Backend with Go", DueDate: time.Now().AddDate(0, 0, 5),
}

func (suite *TaskUseCaseSuite) SetupSuite() {
	// Create the mock repo and the usecase and add to fields
	repo := new(mocks.TaskRepository)
	uc := NewTaskUseCase(repo)
	suite.repository = repo
	suite.TskUseCase = uc 
}

func (suite *TaskUseCaseSuite) TestGetElementByID_Positive() {
	suite.repository.On("GetTaskByID", test_task.ID).Return(test_task, nil)
	// Check operation
	task, err := suite.TskUseCase.GetElementByID(test_task.ID)
	suite.NoError(err, "no error expected with valid id input")
	suite.Equal(test_task, task)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestGetElementByID_NonExistingID_Negative() {
	suite.repository.On("GetTaskByID", "1000").Return(Domain.Task{}, errors.New("Task not found"))
	// Check operation
	task, err := suite.TskUseCase.GetElementByID("1000")
	suite.Error(err, "error expected with invalid id input")
	suite.Equal(Domain.Task{}, task)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestGetAllElements_Positive() {
	suite.repository.On("GetAllElements").Return(Tasks, nil)
	// Check operation
	tasks, err := suite.TskUseCase.GetAllElements() 
	suite.Equal(Tasks, tasks, "didn't return all tasks in repo")
	suite.NoError(err, "No error expected when fetching tasks from task repo")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestCreateTask_Positive() {
	suite.repository.On("CreateTask", test_task).Return(nil)
	// Check function
	err := suite.TskUseCase.CreateTask(test_task)
	suite.NoError(err, "expected no error when creating a task with valid inputs")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestCreateTask_InvalidTask_Negative() {
	suite.repository.On("CreateTask", Domain.Task{}).Return(errors.New("nil Task"))
	// Check function
	err := suite.TskUseCase.CreateTask(Domain.Task{})
	suite.Error(err, "expected error when creating a task with invalid inputs")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestUpdateTask_Positive() {
	suite.repository.On("UpdateTaskByID", test_task.ID, test_task).Return(nil)
	// Check function
	err := suite.TskUseCase.UpdateTask(test_task.ID, test_task)
	suite.NoError(err, "expected no error when updating a task with valid inputs")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestUpdateTask_NonExistingID_Negative() {
	suite.repository.On("UpdateTaskByID", "1000", test_task).Return(errors.New("Task not found"))
	// Check function
	err := suite.TskUseCase.UpdateTask("1000", test_task)
	suite.Error(err, "expected error when updating a task with invalid inputs")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestDeleteTask_Positive() {
	suite.repository.On("DeleteTask", test_task.ID).Return(nil)
	// Check function
	err := suite.TskUseCase.DeleteTask(test_task.ID)
	suite.NoError(err, "no error expected when deleting using valid inputs")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseSuite) TestDeleteTask_NonExistingID_Negative() {
	suite.repository.On("DeleteTask", "1000").Return(errors.New("Task not found"))
	// Check function
	err := suite.TskUseCase.DeleteTask("1000")
	suite.Error(err, "error expected when deleting using invalid inputs")
	suite.repository.AssertExpectations(suite.T())
}

func TestTaskUsecaseSuite(t *testing.T) {
	// Run the suite
	suite.Run(t, new(TaskUseCaseSuite))
}