package health

import (
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/types"
	"net/http"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Handler(ctx *types.Ctx) error {
	w := ctx.Res

	response := types.ServerResponse{Ok: true}

	if err := helper.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
