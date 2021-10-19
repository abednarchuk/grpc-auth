package models

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ScopeAuthentication = "authentication"
)

// Token is the type for authentication tokens
type Token struct {
	PlainText string             `bson:"plainText"`
	UserID    primitive.ObjectID `bson:"userID"`
	Hash      [32]byte           `bson:"hash"`
	Expiry    time.Time          `bson:"expiry"`
	Scope     string             `bson:"scope"`
}

// GenerateToken generates token that lasts for ttl, and returns it
func GenerateToken(userID primitive.ObjectID, ttl time.Duration, scope string) (*Token, error) {
	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	token.PlainText = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	token.Hash = sha256.Sum256([]byte(token.PlainText))
	return token, nil
}
