package gmail

import (
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/types"
	"net/http"
)

type GoogleOauthCallbackController struct {
	service *GoogleOauthCallbackService
}

func NewGoogleOauthCallbackController(
	service *GoogleOauthCallbackService,
) *GoogleOauthCallbackController {
	return &GoogleOauthCallbackController{
		service: service,
	}
}

func (h *GoogleOauthCallbackController) Handler(w http.ResponseWriter, r *http.Request) error {
	response := types.ServerResponse{
		Ok: true,
	}

	if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusOK, response); encodeErr != nil {
		return encodeErr
	}

	return nil
}
