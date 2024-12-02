package common

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type StandardClaim struct {
	UserId int    `json:"id,omitempty"`
	From   string `json:"from,omitempty"`
	Email  string `json:"email,omitempty"`

	jwt.RegisteredClaims
}

func EncodeJwtToken(customClaim *StandardClaim) (string, error) {
	// default claim data
	var claim *StandardClaim
	if customClaim != nil {
		claim = customClaim
	} else {
		duration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRY"))
		claim = &StandardClaim{
			UserId: 0,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "date-me",
				Subject:   "date-me-service",
				Audience:  []string{"xavier"},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(duration) * time.Minute)),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		}
	}

	// generate signed jwt key, to generate jwt key first you
	// must define secret key on configuration file.
	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claim)
	token, err := sign.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		panic(err)
	}

	return token, nil
}
