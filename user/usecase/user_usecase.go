package usecase

import (
	"context"
	"github.com/tamaApotek/tama-go-server/query"

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

	return u.userRepo.FindByID(c, id)
}

func (u *userUsecase) SearchText(ctx context.Context, queryString string) ([]*models.User, error) {
	if queryString == "" {
		return nil, query.NewErrorQuery("Invalid search parameter", query.ErrorEnum.Invalid, nil)
	}

	// TODO: Search condition?

	return u.userRepo.SearchText(ctx, queryString)
}
