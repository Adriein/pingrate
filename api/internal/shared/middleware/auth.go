package middleware

import (
	"context"
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

const UserContextKey = "user"

func NewAuthMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, cookieErr := r.Cookie("jwt")

		if cookieErr != nil {
			response := types.ServerResponse{
				Ok:    false,
				Error: constants.MissingJwt,
			}

			if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusUnauthorized, response); encodeErr != nil {
				http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
			}

			return
		}

		tokenStr := cookie.Value

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv(constants.JwtSecret)), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		if err != nil || !token.Valid {
			response := types.ServerResponse{
				Ok:    false,
				Error: constants.InvalidJwt,
			}

			if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusUnauthorized, response); encodeErr != nil {
				http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
			}

			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userEmail, okEmail := claims["user"].(string)
			if !okEmail {
				response := types.ServerResponse{
					Ok:    false,
					Error: constants.InvalidStructureJwt,
				}

				if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusUnauthorized, response); encodeErr != nil {
					http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
				}

				return
			}

			ctx := context.WithValue(r.Context(), "user", userEmail)

			handler.ServeHTTP(w, r.WithContext(ctx))

			return
		}

		response := types.ServerResponse{
			Ok:    false,
			Error: constants.InvalidStructureJwt,
		}

		if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusUnauthorized, response); encodeErr != nil {
			http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
		}

		handler.ServeHTTP(w, r)
	})
}
