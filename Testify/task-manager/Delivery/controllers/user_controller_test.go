package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"task_manager_ca/Domain"
	"task_manager_ca/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type UserControllerSuite struct {
	suite.Suite
	useCase *mocks.UserUseCaseI
	UsrCtrl UserController
	testServer *httptest.Server
}

var test_user = Domain.User{
	ID: "1",
	Username: "Mash",
	Password: "123",
	Role: "Admin",
	Name: "Mahder Ashenafi",
}

func (suite *UserControllerSuite) SetupSuite() {
	// Initialize dependencies
	usecase := new(mocks.UserUseCaseI)
	uctr := NewUserController(usecase)

	// Initialize a test server using all endpoints for users
	router := gin.Default()
	router.POST("/register", suite.UsrCtrl.RegisterController)
	router.GET("/user_profile", suite.UsrCtrl.UserProfileController)
	router.GET("/admin_page", suite.UsrCtrl.AdminPageController)
	router.POST("/login", suite.UsrCtrl.LoginController)
	testServer := httptest.NewServer(router)

	// Assign to fields
	suite.UsrCtrl = uctr
	suite.useCase = usecase
	suite.testServer = testServer 
}

func (suite *UserControllerSuite) TearDownSuite() {
	// Stop running test server
	suite.testServer.Close()
}

func (suite *UserControllerSuite) TestRegisterController_Positive() {
	// Setup expected responses from mocks
	suite.useCase.On("Register", test_user).Return(nil)

	// Marshaling and some assertion
	requestBody, err := json.Marshal(&test_user)
	suite.NoError(err, "No error expected when marshaling valid user")

	// Send the request to the server (actual test)
	response, err := http.Post(suite.testServer.URL + "/register", "application/json", bytes.NewBuffer(requestBody))

	// Assert expectations
	suite.NoError(err, "No error expected when sending valid register request")
	suite.Equal(http.StatusCreated, response.StatusCode)
	suite.useCase.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestRegisterController_ExistingUsername_Negative() {
	invalid_user := test_user
	invalid_user.Username = "exists"
	// Setup expected responses from mocks
	suite.useCase.On("Register", invalid_user).Return(errors.New("username already in use"))

	// Marshaling and some assertion
	requestBody, err := json.Marshal(&invalid_user)
	suite.NoError(err, "No error expected when marshaling valid user")

	// Send the request to the server (actual test)
	response, err := http.Post(suite.testServer.URL + "/register", "application/json", bytes.NewBuffer(requestBody))

	// Assert expectations
	suite.NoError(err, "No error expected when sending valid register request")
	suite.Equal(http.StatusInternalServerError, response.StatusCode)
	suite.useCase.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestUserProfileController_Positive() {
	// Send request (check actual function)
	resp, err := http.Get(suite.testServer.URL+"/user_profile")

	// Assert
	suite.NoError(err, "No error expected when getting user profile")
	suite.Equal(http.StatusOK, resp.StatusCode)
}

func (suite *UserControllerSuite) TestAdminPageController_Positive() {
	// Send request (check actual function)
	resp, err := http.Get(suite.testServer.URL+"/admin_page")

	// Assert
	suite.NoError(err, "No error expected when getting user profile")
	suite.Equal(http.StatusOK, resp.StatusCode)
}

func (suite *UserControllerSuite) TestLoginController_Positive() {
	// Setup dependency responses for mock
	suite.useCase.On("Login", test_user).Return("test_token", nil)

	// Marshaling and some assertion
	requestBody, err := json.Marshal(&test_user) 
	suite.NoError(err, "No error expected when marshaling valid user to json")

	// Send request
	response, err := http.Post(suite.testServer.URL+"/login", "application/json", bytes.NewBuffer(requestBody))
	
	//Assert expectations
	suite.Equal(http.StatusCreated, response.StatusCode)
	suite.NoError(err, "No error expected when sending valid login request with post")
	suite.useCase.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestLoginController_NonExistingUsername_Negative() {
	invalid_user := test_user
	invalid_user.Username = "doesnt exist"
	// Setup dependency responses for mock
	suite.useCase.On("Login", invalid_user).Return("", errors.New("Username not found"))

	// Marshaling and some assertion
	requestBody, err := json.Marshal(&invalid_user) 
	suite.NoError(err, "No error expected when marshaling valid user to json")

	// Send request
	response, err := http.Post(suite.testServer.URL+"/login", "application/json", bytes.NewBuffer(requestBody))
	
	//Assert expectations
	suite.Equal(http.StatusUnauthorized, response.StatusCode)
	suite.NoError(err, "No error expected when sending valid login request with post")
	suite.useCase.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestLoginController_WrongPassword_Negative() {
	invalid_user := test_user
	invalid_user.Password = "111"
	// Setup dependency responses for mock
	suite.useCase.On("Login", invalid_user).Return("", errors.New("invalid username or password"))

	// Marshaling and some assertion
	requestBody, err := json.Marshal(&invalid_user) 
	suite.NoError(err, "No error expected when marshaling valid user to json")

	// Send request
	response, err := http.Post(suite.testServer.URL+"/login", "application/json", bytes.NewBuffer(requestBody))
	
	//Assert expectations
	suite.Equal(http.StatusUnauthorized, response.StatusCode)
	suite.NoError(err, "No error expected when sending valid login request with post")
	suite.useCase.AssertExpectations(suite.T())
}

func TestUserControllerSuite(t *testing.T) {
	// Run the suite
	suite.Run(t, new(UserControllerSuite))
}