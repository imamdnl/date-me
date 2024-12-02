package usecase

import (
	"date-me/domain/repository"
	"date-me/domain/usecase"
	"go.uber.org/zap"
)

func Init(repo *repository.Wrapper, logger *zap.Logger) *usecase.Wrapper {
	return &usecase.Wrapper{
		UserUseCase: NewUserUseCase(repo, logger),
	}
}

type userUseCaseInteractor struct {
	userRepo repository.UserRepository
	logger   *zap.Logger
}

func NewUserUseCase(w *repository.Wrapper, logger *zap.Logger) usecase.IUserUseCase {
	return &userUseCaseInteractor{
		userRepo: w.UserRepo,
		logger:   logger,
	}
}
