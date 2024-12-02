package usecase

import (
	"context"
	"date-me/domain/dto"
	"date-me/domain/entity"
)

type IUserUseCase interface {
	RegisterUser(ctx context.Context, user *entity.User) bool
	LoginUser(ctx context.Context, userDto *dto.User) (*entity.User, error)
}
