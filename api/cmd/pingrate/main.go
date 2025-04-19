package main

import (
	"github.com/adriein/pingrate/internal/pingrate/server"
	"github.com/adriein/pingrate/internal/pingrate/shared/constants"
	"github.com/adriein/pingrate/internal/pingrate/shared/helper"
	"github.com/joho/godotenv"
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
	)

	if envCheckerErr := checker.Check(); envCheckerErr != nil {
		log.Fatal(envCheckerErr.Error())
	}

	_, newServerErr := server.New(os.Getenv(constants.ServerPort))

	if newServerErr != nil {
		log.Fatal(newServerErr.Error())
	}
}
