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

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (lur *LoginUserRequest) Validate() error {
	err := validation.ValidateStruct(lur,
		validation.Field(&lur.Email, validation.Required, is.Email),
		validation.Field(&lur.Password, validation.Required, validation.Length(8, 50)),
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

type LoginUserController struct {
	service *LoginUserService
}

func NewLoginUserController(service *LoginUserService) *LoginUserController {
	return &LoginUserController{
		service: service,
	}
}

func (c *LoginUserController) Handler(w http.ResponseWriter, r *http.Request) error {
	var request LoginUserRequest

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

	session, serviceErr := c.service.Execute(request.Email, request.Password)

	if serviceErr != nil {
		if errors.Is(serviceErr, types.UserNotFoundError) || errors.Is(serviceErr, types.UserIncorrectPasswordError) {
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

	http.SetCookie(w, &http.Cookie{
		Name:     "$session",
		Value:    session.Id,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false, // Set to true in production (requires HTTPS)
		SameSite: http.SameSiteLaxMode,
	})

	response := types.ServerResponse{Ok: true}

	if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusOK, response); encodeErr != nil {
		return encodeErr
	}

	return nil
}
