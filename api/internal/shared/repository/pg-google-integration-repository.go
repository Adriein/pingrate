package repository

import (
	"database/sql"
	"errors"
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/rotisserie/eris"
	"strings"
)

type PgGoogleIntegrationRepository struct {
	connection  *sql.DB
	transformer *helper.CriteriaToSqlService
}

func NewPgGoogleIntegrationRepository(connection *sql.DB) *PgGoogleIntegrationRepository {
	transformer, _ := helper.NewCriteriaToSqlService(&types.GoogleToken{})

	return &PgGoogleIntegrationRepository{
		connection:  connection,
		transformer: transformer,
	}
}

func (r *PgGoogleIntegrationRepository) Find(criteria types.Criteria) ([]types.GoogleToken, error) {
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
		gi_id                   string
		gi_user_email           string
		gi_google_access_token  string
		gi_google_token_type    string
		gi_google_refresh_token string
		gi_created_at           string
		gi_updated_at           string
	)

	var results []types.GoogleToken

	for rows.Next() {
		if scanErr := rows.Scan(
			&gi_id,
			&gi_user_email,
			&gi_google_access_token,
			&gi_google_token_type,
			&gi_google_refresh_token,
			&gi_created_at,
			&gi_updated_at,
		); scanErr != nil {
			return nil, eris.New(scanErr.Error())
		}

		results = append(results, types.GoogleToken{
			Id:           gi_id,
			UserEmail:    gi_user_email,
			AccessToken:  gi_google_access_token,
			TokenType:    gi_google_token_type,
			RefreshToken: gi_google_refresh_token,
			CreatedAt:    gi_created_at,
			UpdatedAt:    gi_updated_at,
		})
	}

	return results, nil
}

func (r *PgGoogleIntegrationRepository) FindOne(criteria types.Criteria) (*types.GoogleToken, error) {
	query, err := r.transformer.Transform(criteria)

	if err != nil {
		return nil, eris.New(err.Error())
	}

	var (
		gi_id                   string
		gi_user_email           string
		gi_google_access_token  string
		gi_google_token_type    string
		gi_google_refresh_token string
		gi_created_at           string
		gi_updated_at           string
	)

	if scanErr := r.connection.QueryRow(query).Scan(
		&gi_id,
		&gi_user_email,
		&gi_google_access_token,
		&gi_google_token_type,
		&gi_google_refresh_token,
		&gi_created_at,
		&gi_updated_at,
	); scanErr != nil {
		if errors.Is(scanErr, sql.ErrNoRows) {
			return nil, eris.Wrap(types.GoogleTokenNotFoundError, "")
		}

		return nil, eris.New(scanErr.Error())
	}

	return &types.GoogleToken{
		Id:           gi_id,
		UserEmail:    gi_user_email,
		AccessToken:  gi_google_access_token,
		TokenType:    gi_google_token_type,
		RefreshToken: gi_google_refresh_token,
		CreatedAt:    gi_created_at,
		UpdatedAt:    gi_updated_at,
	}, nil
}

func (r *PgGoogleIntegrationRepository) Save(entity *types.GoogleToken) error {
	var query strings.Builder

	query.WriteString(`INSERT INTO pi_google_integration `)
	query.WriteString(`(gi_id, gi_user_email, gi_google_access_token, gi_google_token_type, gi_google_refresh_token, gi_created_at, gi_updated_at) `)
	query.WriteString(`VALUES ($1, $2, $3, $4, $5);`)

	_, err := r.connection.Exec(
		query.String(),
		entity.Id,
		entity.UserEmail,
		entity.AccessToken,
		entity.TokenType,
		entity.RefreshToken,
		entity.CreatedAt,
		entity.UpdatedAt,
	)

	if err != nil {
		return eris.New(err.Error())
	}

	return nil
}

func (r *PgGoogleIntegrationRepository) Update(_ *types.GoogleToken) error {
	return eris.New("Method not implemented")
}
