package schedule

import "context"

type Usecase interface {
	Create(ctx context.Context, sch *Schedule) (string, error)
	FindByDoctorID(ctx context.Context, doctorID string) ([]*Schedule, error)
}
