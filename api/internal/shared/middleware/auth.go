package middleware

import (
	"errors"
	"github.com/adriein/pingrate/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

const SessionContextKey = "session"

func Auth(repository auth.SessionRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, cookieErr := ctx.Cookie("$session")

		if cookieErr != nil {
			if errors.Is(cookieErr, http.ErrNoCookie) {
				ctx.Status(http.StatusUnauthorized)
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": cookieErr.Error()})
			return
		}

		session, findByIdErr := repository.FindById(value)

		if findByIdErr != nil {
			if errors.Is(findByIdErr, auth.SessionNotFoundError) {
				ctx.Status(http.StatusUnauthorized)
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": cookieErr.Error()})
			return
		}

		ctx.Set(SessionContextKey, session.Email)

		ctx.Next()
	}
}
