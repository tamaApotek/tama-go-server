package doctor

import (
	"context"
)

// Repository represent doctor repository contract
type Repository interface {
	Add(ctx context.Context, doctor *Doctor) (string, error)
}
