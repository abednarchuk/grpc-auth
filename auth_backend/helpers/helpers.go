package helpers

import (
	"context"
	"time"

	"github.com/abednarchuk/grpc_auth/auth_backend/mongohelper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckIfUsernameAvailable(ctx context.Context, username string) bool {
	userCollection := mongohelper.GetUsersCollection()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	res := userCollection.FindOne(ctx, bson.D{{"username", username}, {}})
	return res.Err() == mongo.ErrNoDocuments
}

// CheckIfEmailAvailable checks if email already exists in database, and returns true if it does, and false if not
func CheckIfEmailAvailable(ctx context.Context, email string) bool {
	userCollection := mongohelper.GetUsersCollection()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	res := userCollection.FindOne(ctx, bson.D{{"email", email}})
	return res.Err() == mongo.ErrNoDocuments
}
