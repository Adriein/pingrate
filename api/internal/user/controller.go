package user

import (
	"errors"
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
		var json CreateUserRequest

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := ctrl.validator.Struct(json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := ctrl.service.CreateUser(&json); err != nil {
			if errors.Is(err, UserAlreadyExistError) {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.Status(http.StatusInternalServerError)
		}
	}
}
