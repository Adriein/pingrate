package middleware

import (
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/gin-gonic/gin"
	"github.com/rotisserie/eris"
	"log/slog"
	"net/http"
)

func Error() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err

			slog.Error(eris.ToString(err, true))

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"ok":   false,
				"data": constants.ServerGenericError,
			})
		}
	}
}
