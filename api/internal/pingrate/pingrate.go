package pingrate

import (
	"github.com/adriein/pingrate/internal/healthcheck"
	"github.com/adriein/pingrate/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rotisserie/eris"
	"log"
	"log/slog"
)

type Pingrate struct {
	router    *gin.RouterGroup
	validator *validator.Validate
}

func New(port string) *Pingrate {
	engine := gin.Default()
	router := engine.Group("/api/v1")

	app := &Pingrate{
		router:    router,
		validator: validator.New(),
	}

	app.routeSetup()

	if ginErr := engine.Run(port); ginErr != nil {
		pingrateErr := eris.Wrap(ginErr, "Error starting HTTP server")

		log.Fatal(eris.ToString(pingrateErr, true))
	}

	slog.Info("Starting the PingrateApiServer at " + port)

	return app
}

func (p *Pingrate) routeSetup() {
	p.router.GET("/ping", healthcheck.NewController().Get())

	p.router.POST("/users", user.NewController(p.validator).Post())
}
