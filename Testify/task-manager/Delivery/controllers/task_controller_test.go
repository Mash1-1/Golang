package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task_manager_ca/Domain"
	"task_manager_ca/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type TaskControllerSuite struct {
	suite.Suite
	useCase *mocks.TaskUseCaseI
	TskCtrl TaskController 
	testServer *httptest.Server
}

// Initialize a task for testing with
var test_task = Domain.Task{
	ID: "1", Status: "In progress", Title: "Task 4", Description: "Backend with Go",
}

func (suite *TaskControllerSuite) SetupSuite() {
	// Initialize depenencies
	uc := new(mocks.TaskUseCaseI)
	router := gin.Default()
	router.GET("/tasks/:id", suite.TskCtrl.GetTaskByID)
	router.GET("/tasks", suite.TskCtrl.GetAllTasks)
	router.POST("/tasks", suite.TskCtrl.CreateTaskController)
	router.PUT("/tasks/:id", suite.TskCtrl.UpdateTaskByID)
	router.DELETE("/tasks/:id", suite.TskCtrl.DeleteTaskController)

	testServer := httptest.NewServer(router)

	// Assign to fields
	suite.TskCtrl = NewTaskController(uc)
	suite.testServer = testServer
	suite.useCase = uc 
}

func (suite *TaskControllerSuite) TestGetTaskByID_Positive() {
	// Setup what to return for mocks
	suite.useCase.On("GetElementByID", test_task.ID).Return(test_task, nil)

	// Send request
	resp, err := http.Get(suite.testServer.URL+"/tasks/"+test_task.ID)
	
	//Assert expectations
	suite.NoError(err, "no error expected when sending valid get request")	
	suite.Equal(http.StatusOK, resp.StatusCode)
	suite.useCase.AssertExpectations(suite.T())
}

func (suite *TaskControllerSuite) TestGetAllTasks_Positive() {
	// Setup dependencies/mocks
	suite.useCase.On("GetAllElements").Return([]Domain.Task{test_task}, nil)

	// Send request 
	resp, err := http.Get(suite.testServer.URL+"/tasks")
	
	//Assert expectations
	suite.NoError(err, "no error expected when getting all elements in taskdb")
	suite.Equal(http.StatusOK, resp.StatusCode)
	suite.useCase.AssertExpectations(suite.T())
}

func (suite *TaskControllerSuite) TestCreateTaskController_Positive() {
	// Setup Dependecy/mocks responses
	suite.useCase.On("CreateTask", test_task).Return(nil)

	// Marshalling and some assertion
	requestBody, err := json.Marshal(test_task)
	suite.NoError(err, "no error expected when marshaling json")


	// Send request 
	resp, err := http.Post(suite.testServer.URL+"/tasks", "application/json", bytes.NewBuffer(requestBody))
	
	//Assert Expectations
	suite.NoError(err, "no error expected when sending valid create task request")
	suite.Equal(http.StatusCreated, resp.StatusCode)
	suite.useCase.AssertExpectations(suite.T())
}

func (suite *TaskControllerSuite) TestUpdateTaskByID_Positive() {
	// Setup Dependencies/mocks
	suite.useCase.On("UpdateTask", test_task.ID, test_task).Return(nil)

	// Marshalling and some assertion
	requestBody, err := json.Marshal(test_task)
	suite.NoError(err, "no error expected when marshaling valid task")

	//setup put request 
	request, err := http.NewRequest(http.MethodPut, suite.testServer.URL+"/tasks/"+test_task.ID, bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error expected when setting up valid update request")
	request.Header.Set("Content-Type", "application/json")
	
	//send request
	resp, err := http.DefaultClient.Do(request)
	suite.NoError(err, "no error expected when ssending valid update request")
	suite.Equal(http.StatusCreated, resp.StatusCode)
	suite.useCase.AssertExpectations(suite.T())
}

func (suite *TaskControllerSuite) TestDeleteTaskController_Positive() {
	// Setup mocks/dependencies
	suite.useCase.On("DeleteTask", test_task.ID).Return(nil)

	// Setup request
	request, err := http.NewRequest(http.MethodDelete, suite.testServer.URL+"/tasks/" + test_task.ID, bytes.NewBuffer([]byte{}))
	suite.NoError(err, "no error expected when setting up valid delete request")
	
	// Send request
	resp, err := http.DefaultClient.Do(request)
	suite.NoError(err, "no error expected when sending valid delete request")
	suite.Equal(http.StatusOK, resp.StatusCode)
	suite.useCase.AssertExpectations(suite.T())
}

func TestTaskControllerSuite(t *testing.T) {
	// Run the suite
	suite.Run(t, new(TaskControllerSuite))
}
