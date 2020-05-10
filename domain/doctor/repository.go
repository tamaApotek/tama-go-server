package doctor

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Repository represent doctor repository contract
type Repository interface {
	Add(ctx context.Context, doctor *Doctor) (string, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*Doctor, error)
}
