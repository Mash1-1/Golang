package infrastructure

import (
	"net/http"
	"net/http/httptest"
	"task_manager_ca/Domain"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AuthMiddlewareSuite struct {
	suite.Suite 
	middleware AuthMiddleWare
	testServer *httptest.Server
	jwt_helper Domain.JwtService
}

// Initialize testing data
var test_user = Domain.User{
	ID: "1",
	Name: "Mahder Ashenafi",
	Password: "123",
	Role: "admin",
	Username: "machine1569",
}

func (suite *AuthMiddlewareSuite) SetupSuite() {
	// Initialize test server using test endpoint 
	test_router := gin.Default()

	test_router.GET("/protected", suite.middleware.Validate_token(), func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message" : "test successful"})
	})
	test_router.GET("/admin_page", suite.middleware.Validate_token(), suite.middleware.Validate_role(), func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{"message" : "admin passed test"})
	})

	suite.jwt_helper = new(JwtImplementation)
	testServer := httptest.NewServer(test_router)
	suite.testServer = testServer
}

func (suite *AuthMiddlewareSuite) TestValidateToken_Positive() {
	// Set up request
	request, err := http.NewRequest(http.MethodGet, suite.testServer.URL+"/protected", nil)
	suite.NoError(err, "no error expected when setting up valid request")
	
	token, _ := suite.jwt_helper.CreateJwtToken(test_user)
	request.Header.Set("Authorization", "bearer " + token)

	// Send Request 
	resp, err := http.DefaultClient.Do(request)
	suite.NoError(err, "no error expected when sending valid request")
	suite.Equal(http.StatusOK, resp.StatusCode)
}

func (suite *AuthMiddlewareSuite) TestValidateToken_InvalidToken_Negative() {
	// Set up request
	request, err := http.NewRequest(http.MethodGet, suite.testServer.URL+"/protected", nil)
	suite.NoError(err, "no error expected when setting up valid request")
	
	token := "random token"
	request.Header.Set("Authorization", "bearer " + token)

	// Send Request 
	resp, err := http.DefaultClient.Do(request)
	suite.NoError(err, "no error expected when sending valid request")
	suite.Equal(http.StatusUnauthorized, resp.StatusCode)
}

func (suite *AuthMiddlewareSuite) TestValidateRole_Positive() {
	// Set up request 
	request, err := http.NewRequest(http.MethodGet, suite.testServer.URL+"/admin_page", nil)
	suite.NoError(err, "no error expected when setting up valid request")
	token, _ := suite.jwt_helper.CreateJwtToken(test_user)
	request.Header.Set("Authorization", "bearer " + token)

	//send request
	resp, err := http.DefaultClient.Do(request)
	suite.NoError(err, "no error expected when sending valid request")
	suite.Equal(http.StatusOK , resp.StatusCode)
}

func (suite *AuthMiddlewareSuite) TestValidateRole_UserRole_Negative() {
	// Set up request 
	var user = test_user
	user.Role = "user"
	request, err := http.NewRequest(http.MethodGet, suite.testServer.URL+"/admin_page", nil)
	suite.NoError(err, "no error expected when setting up valid request")
	token, _ := suite.jwt_helper.CreateJwtToken(user)
	request.Header.Set("Authorization", "bearer " + token)

	//send request
	resp, err := http.DefaultClient.Do(request)
	suite.NoError(err, "no error expected when sending valid request")
	// Unauthorized expected because role is not admin
	suite.Equal(http.StatusUnauthorized , resp.StatusCode)
}

func TestAuthMiddlewareSuite(t *testing.T) {
	// Run the suite
	suite.Run(t, new(AuthMiddlewareSuite))
}