package repository

import (
	"date-me/domain/repository"
	"date-me/pkg/common"
	"go.uber.org/zap"
)

func Init(conn *common.BaseRepository, logger *zap.Logger) *repository.Wrapper {
	return &repository.Wrapper{
		UserRepo: NewUserRepository(conn, logger),
	}
}

type userRepositoryInteractor struct {
	base   *common.BaseRepository
	logger *zap.Logger
}

func NewUserRepository(connDB *common.BaseRepository, logger *zap.Logger) repository.UserRepository {
	return &userRepositoryInteractor{
		base:   connDB,
		logger: logger,
	}
}
