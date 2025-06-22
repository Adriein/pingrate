package healthcheck

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl *Controller) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}
}
