package infrastructure

import (
	"task_manager_ca/Domain"
	"testing"

	"github.com/stretchr/testify/suite"
)

type JwtServiceSuite struct {
	suite.Suite
	jwtServ *JwtImplementation
}

// Initialize testing data
var test_user2 = Domain.User{
	ID: "1",
	Name: "Mahder Ashenafi",
	Password: "123",
	Role: "admin",
	Username: "machine1569",
}

func (suite *JwtServiceSuite) SetupSuite() {
	// Initialize the jwt service
	suite.jwtServ = new(JwtImplementation)
}

func (suite *JwtServiceSuite) TestCreatJwtToken_Positive() {
	_, err := suite.jwtServ.CreateJwtToken(test_user2)
	suite.NoError(err, "no error expected when creating token from valid user")
}

func TestJwtServiceSuite(t *testing.T) {
	// Run the suite
	suite.Run(t, new(JwtServiceSuite))
}