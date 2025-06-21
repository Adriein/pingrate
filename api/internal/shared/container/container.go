package container

import (
	"database/sql"
	"fmt"
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/adriein/pingrate/internal/shared/repository"
	"log"
	"os"
)

const (
	ContainerInstanceKey = "container"
)

const (
	DatabaseInstanceKey                    = "database"
	UserRepositoryInstanceKey              = "user_repository"
	SessionRepositoryInstanceKey           = "session_repository"
	GoogleIntegrationRepositoryInstanceKey = "google_integration_repository"
)

type Container map[string]interface{}

func New() Container {
	container := make(map[string]interface{})

	database := initDatabase()

	container[DatabaseInstanceKey] = database
	container[UserRepositoryInstanceKey] = repository.NewPgUserRepository(database)
	container[SessionRepositoryInstanceKey] = repository.NewPgSessionRepository(database)
	container[GoogleIntegrationRepositoryInstanceKey] = repository.NewPgGoogleIntegrationRepository(database)

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

func (c Container) Get(instance string) interface{} {
	return c[instance]
}
