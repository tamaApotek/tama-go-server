package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Doctor represent Doctor model
type Doctor struct {
	ID *primitive.ObjectID `json:"id" bson:"_id"`
	UserID string `json:"user_id" bson:"user_id"`
	FullName string `json:"full_name" bson:"full_name" validate:"required"`
}