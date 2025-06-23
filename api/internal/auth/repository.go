package auth

import (
	"database/sql"
	"errors"
	"github.com/rotisserie/eris"
	"strings"
)

type SessionRepository interface {
	FindByEmail(email string) (*Session, error)
	Save(entity *Session) error
	Update(entity *Session) error
}

type PgSessionRepository struct {
	connection *sql.DB
}

func NewPgSessionRepository(connection *sql.DB) *PgSessionRepository {
	return &PgSessionRepository{
		connection: connection,
	}
}

func (r *PgSessionRepository) FindByEmail(email string) (*Session, error) {
	statement, err := r.connection.Prepare("SELECT * FROM pi_session WHERE se_email = $1;")

	if err != nil {
		return nil, eris.New(err.Error())
	}

	var (
		se_id         string
		se_email      string
		se_created_at string
		se_updated_at string
	)

	if scanErr := statement.QueryRow(email).Scan(
		&se_id,
		&se_email,
		&se_created_at,
		&se_updated_at,
	); scanErr != nil {
		if errors.Is(scanErr, sql.ErrNoRows) {
			return nil, eris.Wrap(SessionNotFoundError, "")
		}

		return nil, eris.New(scanErr.Error())
	}

	return &Session{
		Id:        se_id,
		Email:     se_email,
		CreatedAt: se_created_at,
		UpdatedAt: se_updated_at,
	}, nil
}

func (r *PgSessionRepository) Save(entity *Session) error {
	var query strings.Builder

	query.WriteString(`INSERT INTO pi_session `)
	query.WriteString(`(se_id, se_email, se_created_at, se_updated_at) `)
	query.WriteString(`VALUES ($1, $2, $3, $4);`)

	_, err := r.connection.Exec(
		query.String(),
		entity.Id,
		entity.Email,
		entity.CreatedAt,
		entity.UpdatedAt,
	)

	if err != nil {
		return eris.New(err.Error())
	}

	return nil
}

func (r *PgSessionRepository) Update(_ *Session) error {
	return eris.New("Method not implemented")
}
