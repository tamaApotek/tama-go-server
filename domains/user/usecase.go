package user

import (
	"context"
	"fmt"

	"github.com/tamaApotek/tama-go-server/domains/apperror"
)

// Usecase represent User Usecase
type Usecase interface {
	Create(ctx context.Context, user *User) (string, error)
	UpdateOne(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (user *User, err error)
	SearchText(ctx context.Context, queryString string) ([]*User, error)
}

type userUsecase struct {
	userRepo Repository
}

// NewUsecase create new user usecase
func NewUsecase(ur Repository) Usecase {
	return &userUsecase{ur}
}

func (u *userUsecase) Create(c context.Context, user *User) (string, error) {
	// TODO: insert user to auth service

	return u.userRepo.Create(c, user)
}

func (u *userUsecase) UpdateOne(c context.Context, user *User) error {

	return u.userRepo.UpdateOne(c, user)
}

func (u *userUsecase) FindByID(c context.Context, id string) (user *User, err error) {

	return u.userRepo.FindByID(c, id)
}

func (u *userUsecase) SearchText(ctx context.Context, queryString string) ([]*User, error) {
	if queryString == "" {
		return nil, fmt.Errorf("Invalid query parameter: %w", apperror.ErrInvalid)
	}

	// TODO: Search condition?
	return u.userRepo.SearchText(ctx, queryString)
}
