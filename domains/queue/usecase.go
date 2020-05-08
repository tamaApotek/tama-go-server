package queue

import (
	"context"

	"github.com/tamaApotek/tama-go-server/domains/doctor"
)

type usecase struct {
	queueRepo  Repository
	doctorRepo doctor.Repository
}

func NewUsecase(queueRepo Repository, doctorRepo doctor.Repository) Usecase {
	return &usecase{queueRepo, doctorRepo}
}

func (uc *usecase) Add(ctx context.Context, que *Queue) (string, error) {
	isValid := que.Date.IsValid()
	if !isValid {
		return "", ErrInvalidDate
	}

	// TODO: Validate date >= today

	if que.DoctorID == nil {
		return "", ErrInvalidDoctor
	}

	doctor, err := uc.doctorRepo.FindByID(ctx, *que.DoctorID)
	if err != nil {
		return "", err
	}

	if doctor == nil {
		return "", ErrInvalidDoctor
	}

	return uc.queueRepo.Add(ctx, que)
}
