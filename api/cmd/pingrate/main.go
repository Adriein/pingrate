package main

import (
	"database/sql"
	"fmt"
	"github.com/adriein/pingrate/internal/health"
	"github.com/adriein/pingrate/internal/server"
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/repository"
	"github.com/adriein/pingrate/internal/user"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
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
	)

	if envCheckerErr := checker.Check(); envCheckerErr != nil {
		log.Fatal(envCheckerErr.Error())
	}

	api, newServerErr := server.New(os.Getenv(constants.ServerPort))

	if newServerErr != nil {
		log.Fatal(newServerErr.Error())
	}

	databaseDsn := fmt.Sprintf(
		"postgresql://%s:%s@localhost:5432/%s?sslmode=disable",
		os.Getenv(constants.DatabaseUser),
		os.Getenv(constants.DatabasePassword),
		os.Getenv(constants.DatabaseName),
	)

	database, dbConnErr := sql.Open("postgres", databaseDsn)

	defer database.Close()

	if dbConnErr != nil {
		log.Fatal(dbConnErr.Error())
	}

	api.Route("GET /health", healthController(api))

	// USER
	api.Route("POST /users", userController(api, database))

	api.Start()
}

func healthController(api *server.PingrateApiServer) http.HandlerFunc {
	controller := health.NewController()

	return api.NewHandler(controller.Handler)
}

func userController(api *server.PingrateApiServer, database *sql.DB) http.HandlerFunc {
	userRepository := repository.NewPgUserRepository(database)

	service := user.NewCreateUserService(userRepository)

	controller := user.NewCreateUserController(service)

	return api.NewHandler(controller.Handler)
}
