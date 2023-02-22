package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FullName     string             `bson:"full_name"`
	PhoneNumber  string             `bson:"phone_number"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"password_hash"`
	Salt         string             `bson:"salt"`
	USDBalance   float64            `bson:"usd_balance"`
	NGNBalance   float64            `bson:"ngn_balance"`
	CreatedAt    primitive.DateTime `bson:"created_at"`
}

type UserRequest struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}
