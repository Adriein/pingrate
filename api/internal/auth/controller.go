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
	return func(c *gin.Context) {
		var json LoginRequest

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := ctrl.validator.Struct(json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		session, err := ctrl.service.CreateSession(&json)

		if err != nil {
			if errors.Is(err, user.UserIncorrectPasswordError) || errors.Is(err, user.UserNotFoundError) {
				c.Status(http.StatusUnauthorized)
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.SetCookie(
			"$session",
			session.Id,
			3600,
			"/",
			"localhost",
			false,
			true,
		)

		c.Status(http.StatusOK)
	}
}
