package container

import (
	"database/sql"
	"fmt"
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/adriein/pingrate/internal/shared/repository"
	"log"
	"os"
)

const DatabaseInstance = "database"
const UserRepositoryInstance = "user_repository"
const SessionRepositoryInstance = "session_repository"
const GoogleIntegrationRepositoryInstance = "google_integration_repository"

func New() map[string]interface{} {
	container := make(map[string]interface{})

	database := initDatabase()

	container[DatabaseInstance] = database
	container[UserRepositoryInstance] = repository.NewPgUserRepository(database)
	container[SessionRepositoryInstance] = repository.NewPgSessionRepository(database)
	container[GoogleIntegrationRepositoryInstance] = repository.NewPgGoogleIntegrationRepository(database)

	return container
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
