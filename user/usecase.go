package user

import (
	"context"

	"github.com/tamaApotek/tama-go-server/models"
)

// Usecase represent User Usecase
type Usecase interface {
	Create(ctx context.Context, user *models.User) (string, error)
	UpdateByUID(ctx context.Context, user *models.User) error
	FindByUID(ctx context.Context, uid string) (user *models.User, err error)
}
