package usecase

import (
	"context"

	"github.com/tamaApotek/tama-go-server/models"
	"github.com/tamaApotek/tama-go-server/user"
)

type userUsecase struct {
	userRepo user.Repository
}

// NewUserUsecase create new user usecase
func NewUserUsecase(ur user.Repository) user.Usecase {
	return &userUsecase{ur}
}

func (u *userUsecase) FindByUID(c context.Context, uid string) (user *models.User, err error) {
	ctx := context.TODO()

	return u.userRepo.FindByUID(ctx, uid)
}
