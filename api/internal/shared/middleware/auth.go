package middleware

import (
	"errors"
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/adriein/pingrate/internal/shared/container"
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/repository"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/rotisserie/eris"
	"net/http"
)

const SessionContextKey = "session"

func Auth() types.Middleware {
	return func(next types.PingrateHttpHandler) types.PingrateHttpHandler {
		return func(ctx *types.Ctx) error {
			r, w := ctx.Req, ctx.Res

			con, ok := ctx.Data[container.ContainerInstanceKey].(container.Container)

			if !ok {
				return eris.New("Container not found")
			}

			sessionRepository, ok := con.Get(container.SessionRepositoryInstanceKey).(repository.SessionRepository)

			if !ok {
				return eris.New("Session repository not found")
			}

			cookie, cookieErr := r.Cookie("$session")

			if cookieErr != nil {
				response := types.ServerResponse{
					Ok:    false,
					Error: constants.MissingSessionCookie,
				}

				if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusUnauthorized, response); encodeErr != nil {
					http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
				}

				return nil
			}

			sessionId := cookie.Value

			criteria := types.NewCriteria().Equal("se_id", sessionId)

			session, sessionRepoErr := sessionRepository.FindOne(criteria)

			if sessionRepoErr != nil {
				if errors.Is(sessionRepoErr, types.SessionNotFoundError) {
					response := types.ServerResponse{
						Ok:    false,
						Error: constants.InvalidSession,
					}

					if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusUnauthorized, response); encodeErr != nil {
						http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
					}

					return nil
				}

				return sessionRepoErr
			}

			ctx.Data[SessionContextKey] = session

			return next(ctx)
		}
	}
}
