package models

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/abednarchuk/grpc_auth/auth_backend/errors"
	"github.com/abednarchuk/grpc_auth/auth_backend/mongohelper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserName   string             `bson:"username,omitempty"`
	Email      string             `bson:"email,omitempty"`
	Password   string             `bson:"password,omitempty"`
	Tokens     []Token            `bson:"tokens,omitempty"`
	CreatedAt  time.Time          `bson:"createdAt"`
	ModifiedAt time.Time          `bson:"modifiedAt"`
}

func (u *User) CreateUser(ctx context.Context) (primitive.ObjectID, error) {
	usersCollection := mongohelper.GetUsersCollection()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	res, err := usersCollection.InsertOne(ctx, u)
	if err != nil {
		log.Println(err)
		return primitive.NilObjectID, errors.InternalServerError
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.InternalServerError
	}

	return oid, nil
}

func (u *User) UpdateUser(ctx context.Context) error {
	usersCollection := mongohelper.GetUsersCollection()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	u.ModifiedAt = time.Now()
	_, err := usersCollection.UpdateByID(ctx, u.ID, u)
	if err != nil {
		return errors.InternalServerError
	}

	return nil
}

// HashPassword hashes password and returns hash or error
func (u *User) HashPassword() error {
	bcryptCost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	if err != nil {
		return errors.InternalServerError
	}
	res, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcryptCost)
	if err != nil {
		return errors.InternalServerError
	}
	u.Password = string(res)
	return nil
}

func (u *User) InsertToken(ctx context.Context, t *Token) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	u.Tokens = append(u.Tokens, *t)
	if err := u.UpdateUser(ctx); err != nil {
		return errors.InternalServerError
	}
	return nil
}
