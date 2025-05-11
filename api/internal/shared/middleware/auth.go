package middleware

import (
	"context"
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rotisserie/eris"
	"net/http"
	"os"
)

const UserContextKey = "user"

func NewAuthMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, cookieErr := r.Cookie("jwt")

		if cookieErr != nil {
			http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, eris.New("unexpected signing method")
			}

			return os.Getenv(constants.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userEmail, okEmail := claims["user"].(string)
			if !okEmail {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user", userEmail)

			handler.ServeHTTP(w, r.WithContext(ctx))

			return
		}

		http.Error(w, "Invalid token structure", http.StatusUnauthorized)

		handler.ServeHTTP(w, r)
	})
}
