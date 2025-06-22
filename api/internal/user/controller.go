package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type CreateUserRequest struct {
	Id       string `json:"id" validate:"uuid4" binding:"required"`
	Email    string `json:"email" validate:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Controller struct {
	validator *validator.Validate
}

func NewController(validator *validator.Validate) *Controller {
	return &Controller{
		validator: validator,
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
		}
	}
}
