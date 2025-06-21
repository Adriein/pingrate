package main

import (
	"database/sql"
	"github.com/adriein/pingrate/internal/gmail"
	"github.com/adriein/pingrate/internal/health"
	"github.com/adriein/pingrate/internal/pingrate"
	"github.com/adriein/pingrate/internal/session"
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/adriein/pingrate/internal/shared/container"
	"github.com/adriein/pingrate/internal/shared/external"
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/middleware"
	"github.com/adriein/pingrate/internal/shared/repository"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/adriein/pingrate/internal/user"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	dotenvErr := godotenv.Load()

	if dotenvErr != nil && os.Getenv(constants.Env) != constants.Production {
		log.Fatal("Error loading .env file")
	}

	checker := helper.NewEnvVarChecker(
		constants.DatabaseUser,
		constants.DatabasePassword,
		constants.DatabaseName,
		constants.ServerPort,
		constants.JwtSecret,
		constants.GoogleClientId,
		constants.GoogleClientSecret,
	)

	if envCheckerErr := checker.Check(); envCheckerErr != nil {
		log.Fatal(envCheckerErr.Error())
	}

	app := pingrate.New(os.Getenv(constants.ServerPort))

	api := app.GetHttpServer()

	database := app.Get(container.DatabaseInstanceKey).(*sql.DB)

	defer database.Close()

	api.Route("GET /health", healthController())

	// SESSION
	api.Route("POST /sessions", createSessionController())

	// USER
	api.Route("POST /users", createUserController(app))
	api.Route("POST /users/login", loginUserController(app))

	// INTEGRATIONS
	api.Route("GET /integrations/gmail", googleIntegrationController(), middleware.Auth())
	api.Route("GET /integrations/gmail/oauth-callback", googleOauthCallbackController(app))

	// GMAIL
	api.Route("GET /gmail", googleGmailController(app), middleware.Auth())

	api.Start()
}

func healthController() types.PingrateHttpHandler {
	controller := health.NewController()

	return controller.Handler
}

func createSessionController() types.PingrateHttpHandler {
	service := session.NewCreateSessionService()

	controller := session.NewCreateUserController(service)

	return controller.Handler
}

func loginUserController(app *pingrate.Pingrate) types.PingrateHttpHandler {
	userRepository := app.Get(container.UserRepositoryInstanceKey).(repository.UserRepository)
	sessionRepository := app.Get(container.SessionRepositoryInstanceKey).(repository.SessionRepository)

	service := user.NewLoginUserService(userRepository, sessionRepository)

	controller := user.NewLoginUserController(service)

	return controller.Handler
}

func createUserController(app *pingrate.Pingrate) types.PingrateHttpHandler {
	userRepository := app.Get(container.UserRepositoryInstanceKey).(repository.UserRepository)

	service := user.NewCreateUserService(userRepository)

	controller := user.NewCreateUserController(service)

	return controller.Handler
}

func googleIntegrationController() types.PingrateHttpHandler {
	service := gmail.NewGoogleOauthService(external.NewGoogleApi())

	controller := gmail.NewGoogleOauthController(service)

	return controller.Handler
}

func googleOauthCallbackController(app *pingrate.Pingrate) types.PingrateHttpHandler {
	userRepository := app.Get(container.UserRepositoryInstanceKey).(repository.UserRepository)
	googleIntegrationRepository := app.Get(container.GoogleIntegrationRepositoryInstanceKey).(repository.GoogleIntegrationRepository)

	service := gmail.NewGoogleOauthCallbackService(userRepository, googleIntegrationRepository, external.NewGoogleApi())

	controller := gmail.NewGoogleOauthCallbackController(service)

	return controller.Handler
}

func googleGmailController(app *pingrate.Pingrate) types.PingrateHttpHandler {
	googleIntegrationRepository := app.Get(container.GoogleIntegrationRepositoryInstanceKey).(repository.GoogleIntegrationRepository)

	service := gmail.NewGetGmailService(external.NewGoogleApi(), googleIntegrationRepository)

	controller := gmail.NewGetGmailController(service)

	return controller.Handler
}
