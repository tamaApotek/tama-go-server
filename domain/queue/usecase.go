package queue

import (
	"context"
	"fmt"
	"time"

	"github.com/tamaApotek/tama-go-server/domain/doctor"
)

const dateFormat = "2006-01-02"

type usecase struct {
	tz *time.Location

	queueRepo  Repository
	doctorRepo doctor.Repository
}

// Usecase represent Queue Usecase contract
type Usecase interface {
	Add(ctx context.Context, queue *Queue) (string, error)
}

// NewUsecase initiate queue usecase
func NewUsecase(tz *time.Location, queueRepo Repository, doctorRepo doctor.Repository) Usecase {
	return &usecase{tz, queueRepo, doctorRepo}
}

func (uc *usecase) Add(ctx context.Context, que *Queue) (string, error) {
	todayDate := time.Now().In(uc.tz).Format(dateFormat)

	if que.Date < todayDate {
		return "", ErrInvalidDate
	}

	doctor, err := uc.doctorRepo.FindByID(ctx, *que.DoctorID)
	if err != nil {
		return "", err
	}

	if doctor == nil {
		return "", fmt.Errorf("%w. Doctor not found", ErrInvalidDoctor)
	}

	return uc.queueRepo.Add(ctx, que)
}
