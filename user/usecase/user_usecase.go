package usecase

import (
	"context"
	"fmt"

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

	existing, err := u.userRepo.FindByUID(c, user.UID)
	if err != nil {
		return "", models.ErrInternal
	}

	if existing != nil {
		return "", fmt.Errorf("User %v, already exists: %w", user.Email, models.ErrConflict)
	}

	return u.userRepo.Create(c, user)
}

func (u *userUsecase) UpdateByUID(c context.Context, user *models.User) error {

	return u.userRepo.UpdateByUID(c, user)

}

func (u *userUsecase) FindByUID(c context.Context, uid string) (user *models.User, err error) {
	ctx := context.TODO()

	return u.userRepo.FindByUID(ctx, uid)
}
