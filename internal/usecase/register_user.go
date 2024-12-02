package usecase

import (
	"context"
	"date-me/domain/entity"
)

func (u *userUseCaseInteractor) RegisterUser(ctx context.Context, user *entity.User) bool {
	return u.userRepo.StoreUser(ctx, user)
}
