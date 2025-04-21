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
		id         string
		email      string
		password   string
		created_at string
		updated_at string
	)

	var results []types.User

	for rows.Next() {
		if scanErr := rows.Scan(
			&id,
			&email,
			&password,
			&password,
			&created_at,
			&updated_at,
		); scanErr != nil {
			return nil, eris.New(scanErr.Error())
		}

		results = append(results, types.User{
			Id:        id,
			Email:     email,
			Password:  password,
			CreatedAt: created_at,
			UpdatedAt: updated_at,
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
		id         string
		email      string
		password   string
		created_at string
		updated_at string
	)

	if scanErr := r.connection.QueryRow(query).Scan(
		&id,
		&email,
		&password,
		&password,
		&created_at,
		&updated_at,
	); scanErr != nil {
		if errors.Is(scanErr, sql.ErrNoRows) {
			return types.User{}, eris.New("Entity Business not found")
		}

		return types.User{}, eris.New(scanErr.Error())
	}

	return types.User{
		Id:        id,
		Email:     email,
		Password:  password,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}, nil
}

func (r *PgUserRepository) Save(entity *types.User) error {
	var query strings.Builder

	query.WriteString(`INSERT INTO user `)
	query.WriteString(`(id, email, password, created_at, updated_at) `)
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
