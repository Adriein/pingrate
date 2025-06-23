package pingrate

import (
	"database/sql"
	"fmt"
	"github.com/adriein/pingrate/internal/health"
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/adriein/pingrate/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rotisserie/eris"
	"log"
	"log/slog"
	"os"
)

type Pingrate struct {
	database  *sql.DB
	router    *gin.RouterGroup
	validator *validator.Validate
}

func New(port string) *Pingrate {
	engine := gin.Default()
	router := engine.Group("/api/v1")

	app := &Pingrate{
		database:  initDatabase(),
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

func initDatabase() *sql.DB {
	databaseDsn := fmt.Sprintf(
		"postgresql://%s:%s@localhost:5432/%s?sslmode=disable",
		os.Getenv(constants.DatabaseUser),
		os.Getenv(constants.DatabasePassword),
		os.Getenv(constants.DatabaseName),
	)

	database, dbConnErr := sql.Open("postgres", databaseDsn)

	if dbConnErr != nil {
		log.Fatal(dbConnErr.Error())
	}

	return database
}

func (p *Pingrate) routeSetup() {
	p.router.GET("/ping", health.NewController().Get())

	p.router.POST("/users", p.createUser())
}

func (p *Pingrate) createUser() gin.HandlerFunc {
	repository := user.NewPgUserRepository(p.database)
	service := user.NewService(repository)

	return user.NewController(
		p.validator,
		service,
	).Post()
}
