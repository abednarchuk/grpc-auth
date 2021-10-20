package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	"github.com/abednarchuk/grpc_auth/auth_backend/authpb"
	"github.com/abednarchuk/grpc_auth/auth_backend/errors"
	"github.com/abednarchuk/grpc_auth/auth_backend/helpers"
	"github.com/abednarchuk/grpc_auth/auth_backend/models"
	"github.com/abednarchuk/grpc_auth/auth_backend/mongohelper"
	"github.com/abednarchuk/grpc_auth/auth_backend/validators"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	authpb.UnimplementedSignupServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	server := &server{}
	mongohelper.InitializeMongoHelper()

	authpb.RegisterSignupServiceServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

func (s *server) IsUsernameAvailable(ctx context.Context, req *authpb.UsernameAvailabilityRequest) (*authpb.AvailabilityResponse, error) {
	username := req.GetUsername()
	r := helpers.CheckIfUsernameAvailable(ctx, username)
	return &authpb.AvailabilityResponse{Available: r}, nil
}
func (s *server) IsEmailAvailable(ctx context.Context, req *authpb.EmailAvailabilityRequest) (*authpb.AvailabilityResponse, error) {
	email := req.GetEmail()
	r := helpers.CheckIfEmailAvailable(ctx, email)
	return &authpb.AvailabilityResponse{Available: r}, nil
}

func (s *server) SignUp(ctx context.Context, req *authpb.SignupRequest) (*authpb.SignupResponse, error) {
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
	emailAvailable := helpers.CheckIfEmailAvailable(ctx, newUser.Email)
	if !emailAvailable {
		return nil, status.Error(codes.AlreadyExists, "Email already in use")
	}
	usernameAvailable := helpers.CheckIfUsernameAvailable(ctx, newUser.UserName)
	if !usernameAvailable {
		return nil, status.Error(codes.AlreadyExists, "Username already in use")
	}
	// TODO: Encrypt password
	err = newUser.HashPassword()
	if err != nil {
		return nil, err
	}

	oid, err := newUser.CreateUser(ctx)
	if err != nil {
		return nil, err
	}
	newUser.ID = oid
	// Generate token
	token, err := models.GenerateToken(oid, time.Hour*168, models.ScopeAuthentication)
	if err != nil {
		return nil, errors.InternalServerError
	}
	userWithToken := &models.User{
		ID:     newUser.ID,
		Tokens: append(newUser.Tokens, *token),
	}
	userWithToken.UpdateUser(ctx)

	return &authpb.SignupResponse{
		User: &authpb.User{
			Id:       newUser.ID.Hex(),
			UserName: newUser.UserName,
			Email:    newUser.Email,
		},
		Token: &authpb.Token{
			PlainText: token.PlainText,
			UserId:    token.UserID.String(),
			Hash:      token.Hash[:],
			Expiry:    token.Expiry.Unix(),
			Scope:     models.ScopeAuthentication,
		},
	}, nil
}
