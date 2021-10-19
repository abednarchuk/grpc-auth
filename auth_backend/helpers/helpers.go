package helpers

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/abednarchuk/grpc_auth/auth_backend/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bcryptCost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	if err != nil {
		return "", errors.InternalServerError
	}
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", errors.InternalServerError
	}
	return string(res), nil
}

func GetMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://root:secret@mongo")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func GetUserCollection(c *mongo.Client) *mongo.Collection {
	return c.Database("grpc-auth").Collection("users")
}
