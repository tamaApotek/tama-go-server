package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represent User model
// AuthID used to store uid from auth service
// ID used to store user profile document id in database
type User struct {
	AuthID      string              `json:"auth_id" bson:"auth_id"`
	ID          *primitive.ObjectID `json:"id" bson:"_id"`
	Role        string              `json:"role" bson:"role"`
	FullName    string              `json:"full_name" bson:"full_name" validate:"required"`
	Email       string              `json:"email" bson:"email"`
	PhoneNumber string              `json:"phone_number" bson:"phone_number"`
}
