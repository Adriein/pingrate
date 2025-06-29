package gmail

import (
	"github.com/adriein/pingrate/internal/shared/constants"
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
		email, ok := ctx.Get(constants.SessionContextKey)

		if !ok {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		link := ctrl.service.GetGmailOauthLink(email.(string))

		ctx.JSON(http.StatusOK, gin.H{
			"data": link,
		})
	}
}

func (ctrl *Controller) PostGoogleOauthCallback() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state := ctx.Query("state")
		code := ctx.Query("code")

		if err := ctrl.service.ExchangeGoogleToken(state, code); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Redirect(http.StatusFound, "http://localhost:5173/dashboard")
	}
}

func (ctrl *Controller) GetGmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, ok := ctx.Get(constants.SessionContextKey)

		if !ok {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		emails, err := ctrl.service.GetGmailInbox(email.(string))

		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": emails,
		})
	}
}
