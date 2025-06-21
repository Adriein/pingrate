package user

import (
	"encoding/json"
	"errors"
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/rotisserie/eris"
	"net/http"
	"time"
)

type CreateUserRequest struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cur *CreateUserRequest) Validate() error {
	err := validation.ValidateStruct(cur,
		validation.Field(&cur.Id, validation.Required, is.UUIDv4),
		validation.Field(&cur.Email, validation.Required, is.Email),
		validation.Field(&cur.Password, validation.Required, validation.Length(8, 50)),
	)

	if err != nil {
		validateJson, marshalErr := json.Marshal(err)

		if marshalErr != nil {
			return eris.New(marshalErr.Error())
		}

		return eris.Wrap(types.ValidationError, string(validateJson))
	}

	return nil
}

type CreateUserController struct {
	service *CreateUserService
}

func NewCreateUserController(service *CreateUserService) *CreateUserController {
	return &CreateUserController{
		service: service,
	}
}

func (c *CreateUserController) Handler(ctx *types.Ctx) error {
	w, r := ctx.Res, ctx.Req

	var request CreateUserRequest

	if decodeErr := json.NewDecoder(r.Body).Decode(&request); decodeErr != nil {
		return eris.New(decodeErr.Error())
	}

	if validationErr := request.Validate(); validationErr != nil {
		if errors.Is(validationErr, types.ValidationError) {
			strJson, jsonErr := helper.ExtractJSON(validationErr.Error())

			if jsonErr != nil {
				return eris.New(jsonErr.Error())
			}

			var result map[string]string

			unMarshalErr := json.Unmarshal([]byte(strJson), &result)

			if unMarshalErr != nil {
				return eris.New(unMarshalErr.Error())
			}

			response := types.ServerResponse{
				Ok:    false,
				Error: constants.ValidationError,
				Data:  result,
			}

			if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusBadRequest, response); encodeErr != nil {
				return encodeErr
			}

			return nil
		}
	}

	user := types.User{
		Id:        request.Id,
		Email:     request.Email,
		Password:  request.Password,
		CreatedAt: time.Now().UTC().Format(time.DateTime),
		UpdatedAt: time.Now().UTC().Format(time.DateTime),
	}

	if serviceErr := c.service.Execute(&user); serviceErr != nil {
		if errors.Is(serviceErr, types.UserAlreadyExistError) {
			value := struct {
				Email string `json:"email"`
			}{
				Email: "Already in use",
			}

			response := types.ServerResponse{
				Ok:    false,
				Error: constants.ValidationError,
				Data:  value,
			}

			if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusBadRequest, response); encodeErr != nil {
				return encodeErr
			}

			return nil
		}

		return serviceErr
	}

	response := types.ServerResponse{Ok: true}

	if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusCreated, response); encodeErr != nil {
		return encodeErr
	}

	return nil
}
