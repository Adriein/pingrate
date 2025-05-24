package gmail

import (
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/rotisserie/eris"
	"net/http"
	"net/url"
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

func (c *GoogleOauthCallbackController) Handler(w http.ResponseWriter, r *http.Request) error {
	parsedUrl, parseUrlErr := url.Parse(r.RequestURI)

	if parseUrlErr != nil {
		return eris.New(parseUrlErr.Error())
	}

	state := parsedUrl.Query().Get("state")
	code := parsedUrl.Query().Get("code")

	c.service.Execute(state, code)

	response := types.ServerResponse{
		Ok: true,
	}

	if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusOK, response); encodeErr != nil {
		return encodeErr
	}

	return nil
}
