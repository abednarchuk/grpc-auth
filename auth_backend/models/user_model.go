package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserName   string             `bson:"username,omitempty"`
	Email      string             `bson:"email,omitempty"`
	Password   string             `bson:"password,omitempty"`
	CreatedAt  time.Time          `bson:"createdAt"`
	ModifiedAt time.Time          `bson:"modifiedAt"`
}
