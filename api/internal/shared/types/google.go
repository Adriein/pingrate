package types

import "github.com/rotisserie/eris"

var (
	GoogleTokenNotFoundError = eris.New("google token not found")
)

type GoogleToken struct {
	Id           string
	UserEmail    string
	AccessToken  string
	TokenType    string
	RefreshToken string
	CreatedAt    string
	UpdatedAt    string
}
