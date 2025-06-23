package pingrate

import (
	"database/sql"
	"fmt"
	"github.com/adriein/pingrate/internal/auth"
	"github.com/adriein/pingrate/internal/health"
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/adriein/pingrate/internal/shared/middleware"
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
	//HEALTH CHECK
	p.router.GET("/ping", health.NewController().Get())

	//AUTH
	p.router.POST("/auth", p.createSession())

	//USERS
	p.router.POST("/users", p.createUser())

	//Integrations
	p.router.GET("/integrations/gmail", p.auth(), health.NewController().Get())
}

func (p *Pingrate) createUser() gin.HandlerFunc {
	repository := user.NewPgUserRepository(p.database)
	service := user.NewService(repository)

	return user.NewController(
		p.validator,
		service,
	).Post()
}

func (p *Pingrate) createSession() gin.HandlerFunc {
	sessionRepository := auth.NewPgSessionRepository(p.database)
	userRepository := user.NewPgUserRepository(p.database)

	service := auth.NewService(sessionRepository, userRepository)

	return auth.NewController(
		p.validator,
		service,
	).Post()
}

func (p *Pingrate) auth() gin.HandlerFunc {
	sessionRepository := auth.NewPgSessionRepository(p.database)

	return middleware.Auth(sessionRepository)
}
