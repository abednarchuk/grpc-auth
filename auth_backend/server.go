package main

import (
	"context"
	"log"
	"net"
	"strings"

	"github.com/abednarchuk/grpc_auth/auth_backend/authpb"
	"github.com/abednarchuk/grpc_auth/auth_backend/controllers"
	"github.com/abednarchuk/grpc_auth/auth_backend/helpers"
	"github.com/abednarchuk/grpc_auth/auth_backend/models"
	"github.com/abednarchuk/grpc_auth/auth_backend/validators"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type server struct {
	mongoClient *mongo.Client
	authpb.UnimplementedSignupServiceServer
}

func (s *server) SignUp(ctx context.Context, req *authpb.SignupRequest) (*authpb.SignupResponse, error) {
	log.Println("SignUp func was invoked with req: ", req)
	user := req.GetUser()
	newUser := &models.User{
		UserName: strings.ToLower(user.GetUserName()),
		Email:    strings.ToLower(user.GetEmail()),
		Password: user.GetPassword(),
	}
	// Validation
	err := validators.ValidateUser(newUser)
	if err != nil {
		return nil, err
	}

	// TODO: Check if fields are available in database

	// TODO: Encrypt password
	encryptedPassword, err := helpers.HashPassword(newUser.Password)
	if err != nil {
		return nil, err
	}
	newUser.Password = encryptedPassword

	ac := controllers.NewAuthController(s.mongoClient)
	oid, err := ac.SignUp(ctx, newUser)
	if err != nil {
		return nil, err
	}
	newUser.ID = *oid

	return &authpb.SignupResponse{
		User: &authpb.User{
			Id:       newUser.ID.Hex(),
			UserName: newUser.UserName,
			Email:    newUser.Email,
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	server := &server{
		mongoClient: getMongoClient(),
	}
	authpb.RegisterSignupServiceServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

func getMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://root:secret@mongo")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}
