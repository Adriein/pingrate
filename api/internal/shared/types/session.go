package types

import "github.com/rotisserie/eris"

var (
	SessionNotFoundError = eris.New("session not found")
)

type Session struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
