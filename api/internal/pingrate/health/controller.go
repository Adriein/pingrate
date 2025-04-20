package health

import (
	"github.com/adriein/pingrate/internal/pingrate/shared/helper"
	"github.com/adriein/pingrate/internal/pingrate/shared/types"
	"net/http"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (h *Controller) Handler(w http.ResponseWriter, _ *http.Request) error {
	response := types.ServerResponse{Ok: true}

	if err := helper.Encode[types.ServerResponse](w, http.StatusCreated, response); err != nil {
		return err
	}

	return nil
}
