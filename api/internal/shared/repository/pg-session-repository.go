package repository

import (
	"database/sql"
	"errors"
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/rotisserie/eris"
	"strings"
)

type PgSessionRepository struct {
	connection  *sql.DB
	transformer *helper.CriteriaToSqlService
}

func NewPgSessionRepository(connection *sql.DB) *PgSessionRepository {
	transformer, _ := helper.NewCriteriaToSqlService(&types.Session{})

	return &PgSessionRepository{
		connection:  connection,
		transformer: transformer,
	}
}

func (r *PgSessionRepository) Find(criteria types.Criteria) ([]types.Session, error) {
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
		se_id         string
		se_email      string
		se_created_at string
		se_updated_at string
	)

	var results []types.Session

	for rows.Next() {
		if scanErr := rows.Scan(
			&se_id,
			&se_email,
			&se_created_at,
			&se_updated_at,
		); scanErr != nil {
			return nil, eris.New(scanErr.Error())
		}

		results = append(results, types.Session{
			Id:        se_id,
			Email:     se_email,
			CreatedAt: se_created_at,
			UpdatedAt: se_updated_at,
		})
	}

	return results, nil
}

func (r *PgSessionRepository) FindOne(criteria types.Criteria) (*types.Session, error) {
	query, err := r.transformer.Transform(criteria)

	if err != nil {
		return nil, eris.New(err.Error())
	}

	var (
		se_id         string
		se_email      string
		se_created_at string
		se_updated_at string
	)

	if scanErr := r.connection.QueryRow(query).Scan(
		&se_id,
		&se_email,
		&se_created_at,
		&se_updated_at,
	); scanErr != nil {
		if errors.Is(scanErr, sql.ErrNoRows) {
			return nil, eris.Wrap(types.SessionNotFoundError, "")
		}

		return nil, eris.New(scanErr.Error())
	}

	return &types.Session{
		Id:        se_id,
		Email:     se_email,
		CreatedAt: se_created_at,
		UpdatedAt: se_updated_at,
	}, nil
}

func (r *PgSessionRepository) Save(entity *types.Session) error {
	var query strings.Builder

	query.WriteString(`INSERT INTO pi_session `)
	query.WriteString(`(se_id, se_email, se_created_at, se_updated_at) `)
	query.WriteString(`VALUES ($1, $2, $3, $4, $5);`)

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

func (r *PgSessionRepository) Update(_ *types.Session) error {
	return eris.New("Method not implemented")
}
