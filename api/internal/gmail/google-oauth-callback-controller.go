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

func (c *GoogleOauthCallbackController) Handler(ctx *types.Ctx) error {
	w, r := ctx.Res, ctx.Req

	parsedUrl, parseUrlErr := url.Parse(r.RequestURI)

	if parseUrlErr != nil {
		return eris.New(parseUrlErr.Error())
	}

	state := parsedUrl.Query().Get("state")
	code := parsedUrl.Query().Get("code")

	if err := c.service.Execute(state, code); err != nil {
		return err
	}

	response := types.ServerResponse{
		Ok: true,
	}

	if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusOK, response); encodeErr != nil {
		return encodeErr
	}

	return nil
}
