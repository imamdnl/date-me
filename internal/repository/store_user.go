package repository

import (
	"context"
	"date-me/domain/entity"
	"go.uber.org/zap"
)

func (u *userRepositoryInteractor) StoreUser(ctx context.Context, user *entity.User) bool {
	q := "INSERT INTO users (name, email, password, age, gender, bio, profile_picture_url) VALUES ($1,$2,$3,$4,$5,$6,$7)"

	var params []interface{}
	params = append(params, user.GetName(), user.GetEmail(), user.GetPassword(), user.GetAge(),
		user.GetGender().GetValue(), user.GetBio(), user.GetPhoto())
	_, err := u.base.PostgresSql.Exec(ctx, q, params...)
	if err != nil {
		u.logger.Error("failed register user", zap.Error(err))
		return false
	}
	return true
}
