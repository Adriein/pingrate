package user

import (
	"database/sql"
	"errors"
	"github.com/rotisserie/eris"
	"strings"
)

type UserRepository interface {
	FindByEmail(email string) (*User, error)
	Save(entity *User) error
	Update(entity *User) error
}

type PgUserRepository struct {
	connection *sql.DB
}

func NewPgUserRepository(connection *sql.DB) *PgUserRepository {
	return &PgUserRepository{
		connection: connection,
	}
}

func (r *PgUserRepository) FindByEmail(email string) (*User, error) {
	query := "SELECT * FROM pi_user WHERE us_email = $1;"

	var (
		us_id         string
		us_email      string
		us_password   string
		us_created_at string
		us_updated_at string
	)

	if scanErr := r.connection.QueryRow(query).Scan(
		&us_id,
		&us_email,
		&us_password,
		&us_created_at,
		&us_updated_at,
	); scanErr != nil {
		if errors.Is(scanErr, sql.ErrNoRows) {
			return nil, eris.Wrap(UserNotFoundError, "")
		}

		return nil, eris.New(scanErr.Error())
	}

	return &User{
		Id:        us_id,
		Email:     us_email,
		Password:  us_password,
		CreatedAt: us_created_at,
		UpdatedAt: us_updated_at,
	}, nil
}

func (r *PgUserRepository) Save(entity *User) error {
	var query strings.Builder

	query.WriteString(`INSERT INTO pi_user `)
	query.WriteString(`(us_id, us_email, us_password, us_created_at, us_updated_at) `)
	query.WriteString(`VALUES ($1, $2, $3, $4, $5);`)

	_, err := r.connection.Exec(
		query.String(),
		entity.Id,
		entity.Email,
		entity.Password,
		entity.CreatedAt,
		entity.UpdatedAt,
	)

	if err != nil {
		return eris.New(err.Error())
	}

	return nil
}

func (r *PgUserRepository) Update(_ *User) error {
	return eris.New("Method not implemented")
}
