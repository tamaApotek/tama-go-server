package user

import (
	"context"
	"github.com/tamaApotek/tama-go-server/models"
)

// Usecase represent User Usecase
type Usecase interface {
	FindByUID(ctx context.Context, uid string) (user *models.User, err error)
}
