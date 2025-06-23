package gmail

import (
	"github.com/adriein/pingrate/internal/shared/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) GetOauthLink() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, ok := ctx.Get(middleware.SessionContextKey)

		if !ok {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		ctrl.service.GetGmailOauthLink(email.(string))
	}
}

func (ctrl *Controller) PostGoogleOauthCallback() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
