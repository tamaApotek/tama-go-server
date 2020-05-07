package doctor

import "go.mongodb.org/mongo-driver/bson/primitive"

// Doctor represent doctor model structure
type Doctor struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID     string             `json:"user_id" bson:"user_id"`
	FullName   string             `json:"full_name" bson:"full_name" validate:"required"`
	Specialist string             `json:"specialist" bson:"specialist" validate:"required"`
}
