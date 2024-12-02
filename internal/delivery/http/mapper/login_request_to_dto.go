package mapper

import (
	"date-me/domain/dto"
	"date-me/internal/delivery/http/request"
)

func (m *Mapper) LoginRequestToDTO(p request.UserLoginRequest) *dto.User {
	return &dto.User{
		Email:    p.Email,
		Password: p.Password,
	}
}
