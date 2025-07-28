package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type PasswordServiceSuite struct {
	suite.Suite 
	passServ *BcryptPasswordService 
}

func (suite *PasswordServiceSuite) SetupSuite() {
	// Initialize password service
	suite.passServ = new(BcryptPasswordService)
}

var test_password = "this is a test password"

func (suite *PasswordServiceSuite) TestEncryptPassword_Positive() {
	_, err := suite.passServ.EncryptPassword(test_password)
	suite.NoError(err, "no error expected when hashing valid password")
}

func (suite *PasswordServiceSuite) TestCheckPasswordHash_Positive() {
	// Generate hashed password to check with
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(test_password), bcrypt.DefaultCost)
	suite.NoError(err, "no error expected when hashing valid passowrd")
	found := suite.passServ.CheckPasswordHash(test_password, string(hashedPass))
	suite.True(found)
}

func TestPasswordServiceSuite(t *testing.T) {
	// Run the suite
	suite.Run(t, new(PasswordServiceSuite))
}