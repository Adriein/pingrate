package gmail

import (
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/middleware"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/rotisserie/eris"
	"net/http"
)

type GoogleOauthController struct {
	service *GoogleOauthService
}

func NewGoogleOauthController(
	service *GoogleOauthService,
) *GoogleOauthController {
	return &GoogleOauthController{
		service: service,
	}
}

func (h *GoogleOauthController) Handler(ctx *types.Ctx) error {
	w := ctx.Res

	session, ok := ctx.Data[middleware.SessionContextKey].(*types.Session)

	if !ok {
		return eris.New("Session not present in context")
	}

	googleAuthUrl := h.service.Execute(session.Email)

	response := types.ServerResponse{
		Ok:   true,
		Data: googleAuthUrl,
	}

	if err := helper.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
