package mapper

import (
	"date-me/domain/dto"
	"date-me/domain/entity"
	"date-me/internal/delivery/http/request"
)

func (m *Mapper) RegisterRequestToEntity(p request.UserRegisterRequest) (*entity.User, error) {
	eUser, err := entity.NewUser(&dto.User{
		Name:       p.Name,
		Email:      p.Email,
		Password:   p.Password,
		Age:        p.Age,
		Gender:     p.Gender,
		Bio:        p.Bio,
		Photo:      p.Photo,
		IsVerified: false,
		IsPremium:  false,
	})
	if err != nil {
		return nil, err
	}
	eUser.SetPasswordToHash()

	return eUser, nil
}
