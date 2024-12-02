package controller

import "date-me/domain/usecase"

type Controller struct {
	userUsecase usecase.IUserUseCase
}

// NewController will return Controller object and implementing
func NewController(uc usecase.Wrapper) *Controller {
	return &Controller{
		userUsecase: uc.UserUseCase,
	}
}
