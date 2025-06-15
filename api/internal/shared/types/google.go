package types

import "github.com/rotisserie/eris"

var (
	GoogleTokenNotFoundError = eris.New("google token not found")
)

type GoogleIntegration struct {
	Id           string
	UserEmail    string
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiry       string
	CreatedAt    string
	UpdatedAt    string
}
