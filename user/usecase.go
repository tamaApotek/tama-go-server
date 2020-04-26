package user

import (
	"context"

	"github.com/tamaApotek/tama-go-server/models"
)

// Usecase represent User Usecase
type Usecase interface {
	Create(ctx context.Context, user *models.User) (string, error)
	UpdateOne(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id string) (user *models.User, err error)
	SearchText(ctx context.Context, queryString string) ([]*models.User, error)
}
