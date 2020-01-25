package user

import (
	"context"
	"github.com/tamaApotek/tama-go-server/models"
)

// Repository represent User Repository contract
type Repository interface {
	FindByUID(ctx context.Context, uid string) (res *models.User, err error)
}
