package repository

import (
	"database/sql"
	"errors"
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/rotisserie/eris"
	"strings"
)

type PgUserRepository struct {
	connection  *sql.DB
	transformer *helper.CriteriaToSqlService
}

func NewPgUserRepository(connection *sql.DB) *PgUserRepository {
	transformer, _ := helper.NewCriteriaToSqlService(&types.User{})

	return &PgUserRepository{
		connection:  connection,
		transformer: transformer,
	}
}

func (r *PgUserRepository) Find(criteria types.Criteria) ([]types.User, error) {
	query, err := r.transformer.Transform(criteria)

	if err != nil {
		return nil, eris.New(err.Error())
	}

	rows, queryErr := r.connection.Query(query)

	if queryErr != nil {
		return nil, eris.New(queryErr.Error())
	}

	defer rows.Close()

	var (
		us_id         string
		us_email      string
		us_password   string
		us_created_at string
		us_updated_at string
	)

	var results []types.User

	for rows.Next() {
		if scanErr := rows.Scan(
			&us_id,
			&us_email,
			&us_password,
			&us_created_at,
			&us_updated_at,
		); scanErr != nil {
			return nil, eris.New(scanErr.Error())
		}

		results = append(results, types.User{
			Id:        us_id,
			Email:     us_email,
			Password:  us_password,
			CreatedAt: us_created_at,
			UpdatedAt: us_updated_at,
		})
	}

	return results, nil
}

func (r *PgUserRepository) FindOne(criteria types.Criteria) (types.User, error) {
	query, err := r.transformer.Transform(criteria)

	if err != nil {
		return types.User{}, eris.New(err.Error())
	}

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
			return types.User{}, eris.Wrap(types.UserNotFoundError, "")
		}

		return types.User{}, eris.New(scanErr.Error())
	}

	return types.User{
		Id:        us_id,
		Email:     us_email,
		Password:  us_password,
		CreatedAt: us_created_at,
		UpdatedAt: us_updated_at,
	}, nil
}

func (r *PgUserRepository) Save(entity *types.User) error {
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

func (r *PgUserRepository) Update(_ *types.User) error {
	return eris.New("Method not implemented")
}
