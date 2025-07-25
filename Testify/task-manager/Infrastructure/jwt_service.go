package infrastructure

import (
	"task_manager_ca/Domain"

	"github.com/dgrijalva/jwt-go"
)

type JwtImplementation struct{}

// Initialize jwt secret 
var Jwt_secret = []byte("clean architecture")

func (jwtImp JwtImplementation) CreateJwtToken(existingUser Domain.User) (string, error) {
	// Generate JWT
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role" : existingUser.Role,
		"username" : existingUser.Username,
	})

	// Generate JWT token 
	jwtToken, err := token.SignedString(Jwt_secret)

	// Handle error while signing 
	if err != nil {
		return "", err
	}

	// Give the token to the client for the next session use
	return jwtToken, nil
}