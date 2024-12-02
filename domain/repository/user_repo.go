package repository

import (
	"context"
	"date-me/domain/entity"
)

type UserRepository interface {
	StoreUser(ctx context.Context, user *entity.User) bool
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}
