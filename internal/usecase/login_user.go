package usecase

import (
	"context"
	"date-me/domain/dto"
	"date-me/domain/entity"
)

func (u *userUseCaseInteractor) LoginUser(ctx context.Context, userDto *dto.User) (*entity.User, error) {
	eUser, err := u.userRepo.GetUserByEmail(ctx, userDto.Email)
	if err != nil {
		return nil, err
	}
	if err = eUser.CheckPassword(userDto.Password); err != nil {
		return nil, err
	}
	return eUser, nil
}
