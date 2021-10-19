package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	"github.com/abednarchuk/grpc_auth/auth_backend/authpb"
	"github.com/abednarchuk/grpc_auth/auth_backend/controllers"
	"github.com/abednarchuk/grpc_auth/auth_backend/helpers"
	"github.com/abednarchuk/grpc_auth/auth_backend/models"
	"github.com/abednarchuk/grpc_auth/auth_backend/validators"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	mongoClient *mongo.Client
	authpb.UnimplementedSignupServiceServer
}

func (s *server) IsUsernameAvailable(ctx context.Context, req *authpb.UsernameAvailabilityRequest) (*authpb.AvailabilityResponse, error) {
	log.Println("IsUsernameAvailable func was invoked with req: ", req)
	username := req.GetUsername()
	ac := controllers.NewAuthController(s.mongoClient)
	r := ac.CheckIfUsernameAvailable(ctx, username)
	return &authpb.AvailabilityResponse{Available: r}, nil
}
func (s *server) IsEmailAvailable(ctx context.Context, req *authpb.EmailAvailabilityRequest) (*authpb.AvailabilityResponse, error) {
	log.Println("IsEmailAvailable func was invoked with req: ", req)
	email := req.GetEmail()
	ac := controllers.NewAuthController(s.mongoClient)
	r := ac.CheckIfEmailAvailable(ctx, email)
	return &authpb.AvailabilityResponse{Available: r}, nil
}

func (s *server) SignUp(ctx context.Context, req *authpb.SignupRequest) (*authpb.SignupResponse, error) {
	log.Println("SignUp func was invoked with req: ", req)
	user := req.GetUser()
	newUser := &models.User{
		UserName:   strings.ToLower(user.GetUserName()),
		Email:      strings.ToLower(user.GetEmail()),
		Password:   user.GetPassword(),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}
	// Validation
	err := validators.ValidateUser(newUser)
	if err != nil {
		return nil, err
	}

	// TODO: Check if fields are available in database
	ac := controllers.NewAuthController(s.mongoClient)
	emailAvailable := ac.CheckIfEmailAvailable(ctx, newUser.Email)
	if !emailAvailable {
		return nil, status.Error(codes.AlreadyExists, "Email already in use")
	}
	usernameAvailable := ac.CheckIfUsernameAvailable(ctx, newUser.UserName)
	if !usernameAvailable {
		return nil, status.Error(codes.AlreadyExists, "Username already in use")
	}
	// TODO: Encrypt password
	encryptedPassword, err := helpers.HashPassword(newUser.Password)
	if err != nil {
		return nil, err
	}
	newUser.Password = encryptedPassword

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
		mongoClient: helpers.GetMongoClient(),
	}

	authpb.RegisterSignupServiceServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
