package auth

import (
	"errors"
	"github.com/adriein/pingrate/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Controller struct {
	validator *validator.Validate
	service   *Service
}

func NewController(validator *validator.Validate, service *Service) *Controller {
	return &Controller{
		validator: validator,
		service:   service,
	}
}

func (ctrl *Controller) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var json LoginRequest

		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"ok":    false,
				"error": err.Error(),
			})
			return
		}

		if err := ctrl.validator.Struct(json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"ok":    false,
				"error": err.Error(),
			})
			return
		}

		session, err := ctrl.service.CreateSession(&json)

		if err != nil {
			if errors.Is(err, user.UserIncorrectPasswordError) || errors.Is(err, user.UserNotFoundError) {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"ok":    false,
					"error": err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"ok":    false,
				"error": err.Error(),
			})
			return
		}

		ctx.SetCookie(
			"$session",
			session.Id,
			3600,
			"/",
			"localhost",
			false,
			true,
		)

		ctx.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	}
}
