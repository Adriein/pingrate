package oauth

import (
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/types"
	"net/http"
)

type GoogleOauthController struct {
	service *GoogleOauthControllerService
}

func NewGoogleAuthHandler(
	service *GoogleOauthControllerService,
) *GoogleOauthController {
	return &GoogleOauthController{
		service: service,
	}
}

func (h *GoogleOauthController) Handler(w http.ResponseWriter, r *http.Request) error {
	r.c

	googleAuthUrl := h.service.Execute(userId)

	response := types.ServerResponse{
		Ok:   true,
		Data: googleAuthUrl,
	}

	if err := helper.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
