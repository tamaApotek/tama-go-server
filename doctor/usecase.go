package doctor

import (
	"context"
	"github.com/tamaApotek/tama-go-server/models"
)

// Usecase represent Doctor's usecase contract
type Usecase interface {
	Add(ctx context.Context, doctor models.Doctor) (string, error)
}