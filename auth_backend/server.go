package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/abednarchuk/grpc_auth/auth_backend/authpb"
	"github.com/abednarchuk/grpc_auth/auth_backend/controllers"
	"github.com/abednarchuk/grpc_auth/auth_backend/models"
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
	err := s.mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Ping successful")
	ac := controllers.NewAuthController(s.mongoClient)
	user := &models.User{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	res, err := ac.CreateUser(user)
	if err != nil {
		return &authpb.SignupResponse{
			ResponseStatus: authpb.ResponseStatus_STATUS_FAIL,
			SignupErrors: []*authpb.SignupError{
				{
					ErrorMessage: err.Error(),
				},
			}}, err
	}

	return &authpb.SignupResponse{
		ResponseStatus: authpb.ResponseStatus_STATUS_SUCCESS,
		Response:       fmt.Sprintf("%s", res.InsertedID),
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
	clientOptions := options.Client().ApplyURI("mongodb://mongo")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}
