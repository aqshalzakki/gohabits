package user

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (user *User, err error)
	FindByID(ctx context.Context, id uint) (user *User, err error)
}
