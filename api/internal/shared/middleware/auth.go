package middleware

import (
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rotisserie/eris"
	"net/http"
	"os"
	"strings"
)

func NewAuthMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(constants.HTTP_AUTH_HEADER)

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

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

		// Extract user ID and set it in the request context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			_, ok := claims["user"].(string)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}
			// Attach to context
			// ctx := context.WithValue(r.Context(), "userID", userID)
			// next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		http.Error(w, "Invalid token structure", http.StatusUnauthorized)

		handler.ServeHTTP(w, r)
	})
}
