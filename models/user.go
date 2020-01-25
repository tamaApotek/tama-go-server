package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represent User model
type User struct {
	ID *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// UID represent user's auth uid
	UID         string `json:"uid" bson:"uid"`
	FullName    string `json:"full_name" bson:"full_name" validate:"required"`
	Role        string `json:"role" bson:"role"`
	Email       string `json:"email" bson:"email"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
}
