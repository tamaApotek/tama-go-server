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

func (u *userUsecase) Create(c context.Context, user *models.User) (string, error) {
	// TODO: insert user to auth service

	return u.userRepo.Create(c, user)
}

func (u *userUsecase) UpdateOne(c context.Context, user *models.User) error {

	return u.userRepo.UpdateOne(c, user)
}

func (u *userUsecase) FindByID(c context.Context, id string) (user *models.User, err error) {
	ctx := context.TODO()

	return u.userRepo.FindByID(ctx, id)
}