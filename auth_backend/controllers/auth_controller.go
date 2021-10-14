package controllers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/abednarchuk/grpc_auth/auth_backend/errors"
	"github.com/abednarchuk/grpc_auth/auth_backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthController struct {
	client *mongo.Client
}

func NewAuthController(c *mongo.Client) *AuthController {
	return &AuthController{c}
}

func (ac *AuthController) SignUp(ctx context.Context, user *models.User) (*primitive.ObjectID, error) {
	usersCollection := ac.client.Database("grpc-auth").Collection("users")
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	res, err := usersCollection.InsertOne(ctx, user)
	if err != nil {
		log.Println(err)
		return nil, errors.InternalServerError
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot convert to OID"))
	}

	return &oid, nil
}
