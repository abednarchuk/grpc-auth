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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthController struct {
	mongoClient *mongo.Client
}

func NewAuthController(c *mongo.Client) *AuthController {
	return &AuthController{c}
}

func (ac *AuthController) SignUp(ctx context.Context, user *models.User) (*primitive.ObjectID, error) {
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
		return nil, status.Errorf(codes.Internal, "Cannot convert to OID")
	}

	return &oid, nil
}

func (ac *AuthController) CheckIfUsernameAvailable(ctx context.Context, username string) bool {
	userCollection := helpers.GetUserCollection(ac.mongoClient)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	res := userCollection.FindOne(ctx, bson.D{{"username", username}, {}})
	return res.Err() == mongo.ErrNoDocuments
}

func (ac *AuthController) CheckIfEmailAvailable(ctx context.Context, email string) bool {
	userCollection := helpers.GetUserCollection(ac.mongoClient)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	res := userCollection.FindOne(ctx, bson.D{{"email", email}})
	return res.Err() == mongo.ErrNoDocuments
}
