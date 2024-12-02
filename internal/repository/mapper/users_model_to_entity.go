package mapper

import (
	"date-me/domain/dto"
	"date-me/domain/entity"
	"date-me/internal/repository/model"
)

func (mp *Mapper) UsersModelToEntity(m *model.Users) (*entity.User, error) {
	uEntity, err := entity.NewUser(&dto.User{
		Id:         m.Id,
		Name:       m.Name,
		Email:      m.Email,
		Password:   m.Password,
		Age:        m.Age,
		Gender:     m.Gender,
		Bio:        m.Bio,
		Photo:      m.ProfilePictureUrl,
		IsVerified: m.IsVerified,
		IsPremium:  m.IsPremium,
	})
	if err != nil {
		return nil, err
	}
	return uEntity, nil
}
