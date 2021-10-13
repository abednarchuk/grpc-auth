package main

import (
	"context"
	"fmt"

	"github.com/abednarchuk/grpc_auth/auth_backend/authpb"
)

type server struct {
	authpb.UnimplementedSignupServiceServer
}

func (*server) SignUp(ctx context.Context, req *authpb.SignupRequest) (*authpb.SignupResponse, error) {

	return &authpb.SignupResponse{}, nil
}

func main() {
	fmt.Println("test")
}
