package controllers

import "go.mongodb.org/mongo-driver/mongo"

type AuthController struct {
	client *mongo.Client
}

func NewAuthController(c *mongo.Client) *AuthController {
	return &AuthController{c}
}

func (ac *AuthController) CreateUser() {

}
