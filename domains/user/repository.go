package user

import (
	"context"
)

// Repository represent User Repository contract
type Repository interface {
	Create(ctx context.Context, user *User) (string, error)
	UpdateOne(ctx context.Context, user *User) error
	FindByID(ctx context.Context, uid string) (res *User, err error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	SearchText(ctx context.Context, queryString string) ([]*User, error)
}
