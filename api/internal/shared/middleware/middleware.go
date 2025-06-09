package middleware

import (
	"fmt"
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/adriein/pingrate/internal/shared/container"
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/repository"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

type Ctx struct {
	Res  http.ResponseWriter
	Req  *http.Request
	Data map[string]interface{}
}

type PingrateHttpHandler func(ctx *Ctx) error

type Middleware func(next PingrateHttpHandler) PingrateHttpHandler

func Auth() Middleware {
	return func(next PingrateHttpHandler) PingrateHttpHandler {
		return func(ctx *Ctx) error {
			r, w := ctx.Req, ctx.Res

			userRepository, ok := ctx.Data[container.UserRepositoryInstance].(repository.UserRepository)

			if !ok {
				return fmt.Errorf("user repository not found")
			}

			cookie, cookieErr := r.Cookie("$session")

			if cookieErr != nil {
				response := types.ServerResponse{
					Ok:    false,
					Error: constants.MissingJwt,
				}

				if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusUnauthorized, response); encodeErr != nil {
					http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
				}

				return nil
			}

			sessionId := cookie.Value

			userRepository.FindOne()

			ctx := context.WithValue(r.Context(), "user", userEmail)

			handler.ServeHTTP(w, r.WithContext(ctx))

			return

			response := types.ServerResponse{
				Ok:    false,
				Error: constants.InvalidStructureJwt,
			}

			if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusUnauthorized, response); encodeErr != nil {
				http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
			}

			handler.ServeHTTP(w, r)
		}
	}
}
