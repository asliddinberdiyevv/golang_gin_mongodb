package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"first_name"     validate:"required,min=3,max=20"`
	LastName     *string            `json:"last_name"      validate:"required,min=3,max=20"`
	Password     *string            `json:"password"       validate:"required,min=6"`
	Email        *string            `json:"email"          validate:"email,required"`
	Phone        *string            `json:"phone"          validate:"required"`
	UserType     *string            `json:"user_type"      validate:"required,eq=ADMIN|eq=USER"`
	AccessToken  *string            `json:"access_token"`
	RefreshToken *string            `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserID       string             `json:"user_id"`
}

type UserResponse struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserID    string             `json:"user_id"`
	FirstName *string            `json:"first_name"`
	LastName  *string            `json:"last_name"`
	Email     *string            `json:"email"`
	Phone     *string            `json:"phone"`
	UserType  *string            `json:"user_type"`
	Token     UserToken          `json:"token"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type UserToken struct {
	AccessToken  *string `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
}
