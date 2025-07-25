package repositories

import (
	"context"
	"task_manager_ca/Domain"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositorySuite struct {
	suite.Suite 
	usrRepo *UserRepo 
	db *mongo.Collection
}

var test_user = Domain.User{
	ID: "1",
	Name: "Mahder Ashenafi",
	Password: "123",
	Role: "admin",
	Username: "machine1569",
}

func (suite *UserRepositorySuite) SetupSuite() {
	// Initialize Database before the suite runs
	db := InitializeUserDB()

	// Inititalize user repo 
	ur := NewUserRepo(db)

	// Add fields to  suite 
	suite.db = db 
	suite.usrRepo = &ur
}

func (suite *UserRepositorySuite) TearDownTest() {
	// Clear database after each test to ensure test independence 
	suite.db.Database().Collection("users").Drop(context.TODO())
}

func (suite *UserRepositorySuite) TestCreate_Positive() {
	// Test the function 
	err := suite.usrRepo.Create(test_user)

	//assertion for the result of our test
	suite.Assert().NoError(err, "no error expected when creating a new user with valid inputs.")
}

func (suite *UserRepositorySuite) TestLogin_Positive() {
	// Add the user into the database and assertion
	_, err := suite.db.InsertOne(context.TODO(), test_user)
	suite.NoError(err, "can not insert into database to test login")

	// Check the function
	_, err = suite.usrRepo.Login(test_user)
	suite.NoError(err, "no error expected when logging in with a valid user credential.")
}

func (suite *UserRepositorySuite) TestFindUserRepository_Positive() {
	// Add the user into the database and assertion
	_, err := suite.db.InsertOne(context.TODO(), test_user)
	suite.NoError(err, "can not insert into database to test FindUserService")

	// Check the function
	ok := suite.usrRepo.FindUserRepository(test_user.Username)
	suite.True(ok, "expected true while looking for user but found false")
}

func TestUserRepositoryTestSuite(t *testing.T) {
	// Run the suite 
	suite.Run(t, new(UserRepositorySuite))
}