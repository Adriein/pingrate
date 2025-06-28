package auth

import "github.com/rotisserie/eris"

type LoginRequest struct {
	Email    string `json:"email" validate:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var (
	SessionNotFoundError = eris.New("Session not found")
)

type Session struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
