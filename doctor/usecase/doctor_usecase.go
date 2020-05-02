package usecase

import (
	"context"

	"github.com/tamaApotek/tama-go-server/constants"
	"github.com/tamaApotek/tama-go-server/doctor"
	"github.com/tamaApotek/tama-go-server/models"
	"github.com/tamaApotek/tama-go-server/user"
)

type doctorUsecase struct {
	userRepo user.Repository
}

func NewDoctorUsecase(userRepo user.Repository) doctor.Usecase {
	return &doctorUsecase{userRepo}
}

func (du *doctorUsecase) Add(ctx context.Context, doctor models.Doctor) (string, error) {
	user := &models.User{
		Role:     constants.UserRoleDoctor,
		FullName: doctor.FullName,
	}

	userID, err := du.userRepo.Create(ctx, user)
	if err != nil {
		return "", err
	}

	// TODO: Check specialist exists
	// TODO: Add to doctor collection

	return userID, nil
}
