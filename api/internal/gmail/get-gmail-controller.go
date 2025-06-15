package gmail

import (
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/middleware"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/rotisserie/eris"
	"net/http"
)

type GetGmailController struct {
	service *GetGmailService
}

func NewGetGmailController(
	service *GetGmailService,
) *GetGmailController {
	return &GetGmailController{
		service: service,
	}
}

func (h *GetGmailController) Handler(ctx *types.Ctx) error {
	w := ctx.Res

	session, ok := ctx.Data[middleware.SessionContextKey].(*types.Session)

	if !ok {
		return eris.New("Session not present in context")
	}

	err := h.service.Execute(session.Email)

	if err != nil {
		return err
	}

	response := types.ServerResponse{
		Ok: true,
	}

	if err := helper.Encode[types.ServerResponse](w, http.StatusOK, response); err != nil {
		return err
	}

	return nil
}
