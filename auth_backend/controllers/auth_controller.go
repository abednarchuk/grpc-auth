package controllers

import (
	"context"

	"github.com/abednarchuk/grpc_auth/auth_backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	client *mongo.Client
}

func NewAuthController(c *mongo.Client) *AuthController {
	return &AuthController{c}
}

func (ac *AuthController) CreateUser(user *models.User) (*mongo.InsertOneResult, error) {
	usersCollection := ac.client.Database("grpc-auth").Collection("users")

	data, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return data, nil
}
