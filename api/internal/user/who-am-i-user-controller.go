package user

import (
	"errors"
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/middleware"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/rotisserie/eris"
	"net/http"
)

type WhoAmIController struct {
	service *WhoAmIService
}

func NewWhoAmIController(service *WhoAmIService) *WhoAmIController {
	return &WhoAmIController{
		service: service,
	}
}

func (c *WhoAmIController) Handler(w http.ResponseWriter, r *http.Request) error {
	userEmail, ok := r.Context().Value(middleware.UserContextKey).(string)

	if !ok {
		return eris.New("User key inside the context is not a string")
	}

	if serviceErr := c.service.Execute(userEmail); serviceErr != nil {
		if errors.Is(serviceErr, types.UserNotFoundError) {
			response := types.ServerResponse{
				Ok: true,
			}

			if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusUnauthorized, response); encodeErr != nil {
				return encodeErr
			}

			return nil
		}

		return serviceErr
	}

	response := types.ServerResponse{Ok: true}

	if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusOK, response); encodeErr != nil {
		return encodeErr
	}

	return nil
}
