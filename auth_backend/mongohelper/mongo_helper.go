package mongohelper

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoHelper struct {
	mongoClient    *mongo.Client
	grpcDatabase   *mongo.Database
	userCollection *mongo.Collection
}

var mongoStruct *mongoHelper = &mongoHelper{}

func InitializeMongoHelper() {
	mongoStruct.mongoClient = getMongoClient()
	mongoStruct.grpcDatabase = mongoStruct.mongoClient.Database("grpc-auth")
	mongoStruct.userCollection = mongoStruct.grpcDatabase.Collection("users")
}

func GetUsersCollection() *mongo.Collection {
	return mongoStruct.userCollection
}

// GetMongoClient returns a pointer to mongo client
func getMongoClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://root:secret@mongo")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	return client
}
