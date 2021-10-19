package controllers

import (
	"context"
	"log"
	"time"

	"github.com/abednarchuk/grpc_auth/auth_backend/errors"
	"github.com/abednarchuk/grpc_auth/auth_backend/helpers"
	"github.com/abednarchuk/grpc_auth/auth_backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	mongoClient *mongo.Client
}

func NewAuthController(c *mongo.Client) *AuthController {
	return &AuthController{c}
}

// CreateUser creates user in db, and returns new userID or error
func (ac *AuthController) CreateUser(ctx context.Context, user *models.User) (*primitive.ObjectID, error) {
	usersCollection := helpers.GetUserCollection(ac.mongoClient)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	res, err := usersCollection.InsertOne(ctx, user)
	if err != nil {
		log.Println(err)
		return nil, errors.InternalServerError
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.InternalServerError
	}

	return &oid, nil
}

// CheckIfUsernameAvailable checks if username already exists in database, and returns true if it does, and false if not
func (ac *AuthController) CheckIfUsernameAvailable(ctx context.Context, username string) bool {
	userCollection := helpers.GetUserCollection(ac.mongoClient)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	res := userCollection.FindOne(ctx, bson.D{{"username", username}, {}})
	return res.Err() == mongo.ErrNoDocuments
}

// CheckIfEmailAvailable checks if email already exists in database, and returns true if it does, and false if not
func (ac *AuthController) CheckIfEmailAvailable(ctx context.Context, email string) bool {
	userCollection := helpers.GetUserCollection(ac.mongoClient)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	res := userCollection.FindOne(ctx, bson.D{{"email", email}})
	return res.Err() == mongo.ErrNoDocuments
}
