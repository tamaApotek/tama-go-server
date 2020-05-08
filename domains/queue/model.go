package queue

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/tamaApotek/tama-go-server/domains/datestr"
)

// Queue represent daily patient queue
type Queue struct {
	ID *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// PatientID represent patient's id in database
	// can be empty string for patient not registered in database
	PatientID *primitive.ObjectID `json:"patient_id" bson:"patient_id"`
	// DoctorID represent doctor's id
	DoctorID *primitive.ObjectID `json:"doctor_id" bson:"doctor_id" validate:"required"`

	// Date string YYYY-MM-DD
	Date        datestr.DateStr `json:"date" bson:"date"`
	PatientName string          `json:"patient_name" bson:"patient_name"`
	CreatedBy   string          `json:"created_by" bson:"created_by"`
	UpdatedBy   string          `json:"updated_by" bson:"updated_by"`
	CreatedAt   time.Time       `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" bson:"updated_at"`
}

type Usecase interface {
	Add(ctx context.Context, queue *Queue) (string, error)
}

type Repository interface {
	Add(ctx context.Context, queue *Queue) (string, error)
}
