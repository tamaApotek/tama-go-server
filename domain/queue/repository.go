package queue

// Repository represent Queue Repository contract
type Repository interface {
	Add(ctx context.Context, queue *Queue) (string, error)
}
