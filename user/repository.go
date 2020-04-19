package user

import (
	"context"

	"github.com/tamaApotek/tama-go-server/models"
)

// Repository represent User Repository contract
type Repository interface {
	Create(ctx context.Context, user *models.User) (string, error)
	UpdateByUID(ctx context.Context, user *models.User) error
	FindByUID(ctx context.Context, uid string) (res *models.User, err error)
}
