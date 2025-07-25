package infrastructure

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthRoleMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if ok && role == "admin" {
			c.Next()
			return 
		} 
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "unauthorized user", "message" : "Only admins allowed!"})
		c.Abort()
	}
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement Middleware logic 
		authHeader:= c.GetHeader("Authorization")
		if authHeader == "" {
			// Handle missing authorization headers
			c.JSON(401, gin.H{"error" : "Authorization header is required.", "message" : "You have to be registered to view use this feature."})
			c.Abort()
			return 
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "bearer"{
			// Handle invalid authorization format
			c.JSON(401, gin.H{"error": "Invalid Authorization header"})
			c.Abort()
			return  
		}
		// Parse the JWT token
		token, err := jwt.Parse(authParts[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return Jwt_secret, nil
		})

		// Check if the JWT is valid and has the type MapClaims 
		if claims, ok := token.Claims.(jwt.MapClaims); err == nil && ok && token.Valid {
			// Get role and store it for the next handlers to authorize role
			c.Set("role", claims["role"].(string))
		} else {
			c.JSON(401, gin.H{"error" : "Invalid JWT"})
			c.Abort()
			return 
		}
		c.Next()
	}
}