package usecases

import (
	"task_manager_ca/Domain"
	"task_manager_ca/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserUsecaseSuite struct {
	suite.Suite 
	repository *mocks.UserRepository 
	UsrUc UserUseCase
	Password_service *mocks.PasswordService 
	Jwt_serv *mocks.JwtService
}

func (suite *UserUsecaseSuite) SetupSuite() {
	// Initialize the fields and assign to the suite
	ur := new(mocks.UserRepository)
	ps := new(mocks.PasswordService)
	js := new(mocks.JwtService)
	uc := NewUserUseCase(ur, ps, js)
	suite.repository = ur 
	suite.UsrUc = uc 
	suite.Password_service = ps 
	suite.Jwt_serv = js 
}

var test_user = Domain.User{
	ID: "1",
	Username: "Mash",
	Password: "123",
	Role: "Admin",
	Name: "Mahder Ashenafi",
}

func (suite *UserUsecaseSuite) TestRegister_Positive() {
	// Steup expected responses from mocks
	suite.repository.On("FindUserRepository", test_user.Username).Return(false)
	suite.Password_service.On("EncryptPassword", test_user.Password).Return([]uint8(test_user.Password), nil)
	suite.repository.On("Create", test_user).Return(nil)
	
	// Check the actual function
	err := suite.UsrUc.Register(test_user)

	// Assert expectations
	suite.NoError(err, "No error expected when registering with valid credentials")
	suite.repository.AssertExpectations(suite.T())
	suite.Password_service.AssertExpectations(suite.T())
	suite.Jwt_serv.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestLogin_Positive() {
	// Setup expected responses from mocks
	suite.repository.On("Login", test_user).Return(test_user, nil)
	suite.Password_service.On("CheckPasswordHash", test_user.Password, test_user.Password).Return(true)
	suite.Jwt_serv.On("CreateJwtToken", test_user).Return("token", nil)

	// Check actual function
	_, err := suite.UsrUc.Login(test_user)

	// Assert Expectations
	suite.NoError(err, "No errors expected when logging in with valid credentials.")
	suite.repository.AssertExpectations(suite.T())
	suite.Password_service.AssertExpectations(suite.T())
	suite.Jwt_serv.AssertExpectations(suite.T())
}

func TestUserUseCaseSuite(t *testing.T) {
	// Run the suite
	suite.Run(t, new(UserUsecaseSuite))
}