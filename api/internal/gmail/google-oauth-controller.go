package gmail

import (
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/middleware"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/rotisserie/eris"
	"net/http"
)

type GoogleOauthController struct {
	service *GoogleOauthCallbackService
}

func NewGoogleAuthController(
	service *GoogleOauthCallbackService,
) *GoogleOauthController {
	return &GoogleOauthController{
		service: service,
	}
}

func (h *GoogleOauthController) Handler(w http.ResponseWriter, r *http.Request) error {
	userEmail, ok := r.Context().Value(middleware.UserContextKey).(string)

	if !ok {
		return eris.New("User key inside the context is not a string")
	}

	googleAuthUrl := h.service.Execute(userEmail)

	response := types.ServerResponse{
		Ok:   true,
		Data: googleAuthUrl,
	}

	if err := helper.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
