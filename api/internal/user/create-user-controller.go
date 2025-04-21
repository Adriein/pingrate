package user

import (
	"encoding/json"
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/rotisserie/eris"
	"net/http"
)

type CreateUserController struct {
	service *CreateUserService
}

func NewCreateUserController(service *CreateUserService) *CreateUserController {
	return &CreateUserController{
		service: service,
	}
}

func (c *CreateUserController) Handler(w http.ResponseWriter, r *http.Request) error {
	var request types.User

	if decodeErr := json.NewDecoder(r.Body).Decode(&request); decodeErr != nil {
		return eris.New(decodeErr.Error())
	}

	if serviceErr := c.service.Execute(&request); serviceErr != nil {
		return serviceErr
	}

	response := types.ServerResponse{Ok: true}

	if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusCreated, response); encodeErr != nil {
		return encodeErr
	}

	return nil
}
