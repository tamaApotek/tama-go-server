package doctor

import (
	"context"
	"github.com/tamaApotek/tama-go-server/constants"
	"github.com/tamaApotek/tama-go-server/models"
	"github.com/tamaApotek/tama-go-server/user"
)

// Usecase represent Doctor's usecase contract
type Usecase interface {
	Add(ctx context.Context, doctor *Doctor) (string, error)
}

type usecase struct {
	doctorRepo Repository
	userRepo   user.Repository
}

// NewUsecase will initiate Doctor's Usecase
func NewUsecase(doctorRepo Repository, userRepo user.Repository) Usecase {
	return &usecase{doctorRepo: doctorRepo, userRepo: userRepo}
}

func (uc *usecase) Add(ctx context.Context, doctor *Doctor) (string, error) {
	u := &models.User{
		Role:     constants.UserRoleDoctor,
		FullName: doctor.FullName,
	}

	userID, err := uc.userRepo.Create(ctx, u)
	if err != nil {
		return "", err
	}

	// TODO: Check specialist exists
	// TODO: Add to doctor collection

	return userID, nil
}
