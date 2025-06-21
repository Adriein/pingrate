package main

import (
	"database/sql"
	"github.com/adriein/pingrate/internal/auth"
	"github.com/adriein/pingrate/internal/health"
	gmail2 "github.com/adriein/pingrate/internal/integration/gmail"
	"github.com/adriein/pingrate/internal/pingrate"
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

	// AUTH
	api.Route("POST /auth", loginController(app))

	// USER
	api.Route("POST /users", createUserController(app))

	// INTEGRATIONS
	api.Route("GET /integration/gmail/oauth", googleIntegrationController(), middleware.Auth())
	api.Route("GET /integration/gmail/oauth-callback", googleOauthCallbackController(app))

	//api.Route("GET /integration/gmail", googleIntegrationController(), middleware.Auth())

	// GMAIL
	api.Route("GET /gmail", googleGmailController(app), middleware.Auth())

	api.Start()
}

func healthController() types.PingrateHttpHandler {
	controller := health.NewController()

	return controller.Handler
}

func loginController(app *pingrate.Pingrate) types.PingrateHttpHandler {
	userRepository := app.Get(container.UserRepositoryInstanceKey).(repository.UserRepository)
	sessionRepository := app.Get(container.SessionRepositoryInstanceKey).(repository.SessionRepository)

	service := auth.NewLoginService(userRepository, sessionRepository)

	controller := auth.NewLoginController(service)

	return controller.Handler
}

func createUserController(app *pingrate.Pingrate) types.PingrateHttpHandler {
	userRepository := app.Get(container.UserRepositoryInstanceKey).(repository.UserRepository)

	service := user.NewCreateUserService(userRepository)

	controller := user.NewCreateUserController(service)

	return controller.Handler
}

func googleIntegrationController() types.PingrateHttpHandler {
	service := gmail2.NewGoogleOauthService(external.NewGoogleApi())

	controller := gmail2.NewGoogleOauthController(service)

	return controller.Handler
}

func googleOauthCallbackController(app *pingrate.Pingrate) types.PingrateHttpHandler {
	userRepository := app.Get(container.UserRepositoryInstanceKey).(repository.UserRepository)
	googleIntegrationRepository := app.Get(container.GoogleIntegrationRepositoryInstanceKey).(repository.GoogleIntegrationRepository)

	service := gmail2.NewGoogleOauthCallbackService(userRepository, googleIntegrationRepository, external.NewGoogleApi())

	controller := gmail2.NewGoogleOauthCallbackController(service)

	return controller.Handler
}

func googleGmailController(app *pingrate.Pingrate) types.PingrateHttpHandler {
	googleIntegrationRepository := app.Get(container.GoogleIntegrationRepositoryInstanceKey).(repository.GoogleIntegrationRepository)

	service := gmail2.NewGetGmailService(external.NewGoogleApi(), googleIntegrationRepository)

	controller := gmail2.NewGetGmailController(service)

	return controller.Handler
}
