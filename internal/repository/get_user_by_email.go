package repository

import (
	"context"
	"date-me/domain/entity"
	"date-me/internal/repository/mapper"
	"date-me/internal/repository/model"
	"go.uber.org/zap"
)

func (u *userRepositoryInteractor) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	m := new(model.Users)

	query := "SELECT id, email, name, password FROM users WHERE email = $1"

	err := u.base.PostgresSql.QueryRow(ctx, query, email).Scan(&m.Id, &m.Email, &m.Name, &m.Password)
	if err != nil {
		u.logger.Error("error get user by email", zap.Error(err))
		return nil, err
	}

	return mapper.NewMapper().UsersModelToEntity(m)
}
