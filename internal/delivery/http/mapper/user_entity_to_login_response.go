package mapper

import (
	"date-me/domain/entity"
	"date-me/internal/delivery/http/response"
	"date-me/pkg/common"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	"time"
)

func (m *Mapper) UserEntityToLoginResponse(p *entity.User) *response.LoginResponse {
	duration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRY"))
	token, err := common.EncodeJwtToken(&common.StandardClaim{
		UserId: int(p.GetId()),
		From:   "api",
		Email:  p.GetEmail(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "date-me",
			Subject:   "date-me-service",
			Audience:  []string{"xavier"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(duration) * time.Minute)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	if err != nil {
		return nil
	}
	return &response.LoginResponse{Token: token}
}
