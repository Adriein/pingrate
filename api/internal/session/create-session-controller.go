package session

import (
	"encoding/json"
	"errors"
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/types"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/rotisserie/eris"
	"net/http"
	"time"
)

type CreateSessionRequest struct {
	Email string `json:"email"`
}

func (r *CreateSessionRequest) Validate() error {
	err := validation.ValidateStruct(r,
		validation.Field(&r.Email, validation.Required, is.Email),
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

type CreateSessionController struct {
	service *CreateSessionService
}

func NewCreateUserController(service *CreateSessionService) *CreateSessionController {
	return &CreateSessionController{
		service: service,
	}
}

func (c *CreateSessionController) Handler(w http.ResponseWriter, r *http.Request) error {
	var request CreateSessionRequest

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

	sessionId, _ := c.service.Execute()

	http.SetCookie(w, &http.Cookie{
		Name:     "$session",
		Value:    sessionId,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false, // Set to true in production (requires HTTPS)
		SameSite: http.SameSiteLaxMode,
	})

	response := types.ServerResponse{Ok: true}

	if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusCreated, response); encodeErr != nil {
		return encodeErr
	}

	return nil
}
