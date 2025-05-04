package types

import (
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rotisserie/eris"
	"os"
	"time"
)

type JWT struct {
	User  string `json:"user"`
	Token string `json:"token"`
}

func NewJwt(email string) (*JWT, error) {
	claims := jwt.MapClaims{
		"user": email,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // 1 day expiration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	stringToken, err := token.SignedString([]byte(os.Getenv(constants.JwtSecret)))

	if err != nil {
		return nil, eris.New(err.Error())
	}

	return &JWT{
		User:  email,
		Token: stringToken,
	}, nil
}
