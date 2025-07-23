package data

import (
	"context"

	"task_man_v3/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserDatabase *mongo.Collection

func InitializeUserDB() error {
	// Initialize client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// Handle error 
	if err != nil {
		return err
	}
	// Check if connection is established
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err 
	}

	collection := client.Database("user_db").Collection("users")
	// Clear previous data instances 
	collection.DeleteMany(context.TODO(), bson.D{{}})

	// Assign new database into a global variable
	UserDatabase = collection
	return nil 
}

func FindUserService(user_name string) bool {
	filter := bson.D{{Key: "username", Value: user_name}}
	var tmp_user models.User
	err := UserDatabase.FindOne(context.TODO(), filter).Decode(&tmp_user)
	if err == nil { // Username already in the database 
		return true
	}
	return false 
}

func UserRegisterService(user models.User, hashedPass string) error {
	// Add new user into the database
	user.Password = hashedPass
	_, err := UserDatabase.InsertOne(context.TODO(), user)
	return err
}

func UserLoginService(user models.User) (models.User, error){
	var existingUser models.User 
	// User a filter that has the email we are looking for 
	filter := bson.D{{Key: "username", Value: user.Username}}

	// Search for the user in the database
	err := UserDatabase.FindOne(context.TODO(), filter).Decode(&existingUser)
	// Handle error when the user is not found 
	if err != nil {
		return models.User{}, err
	}
	return existingUser, nil 
}