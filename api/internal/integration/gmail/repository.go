package gmail

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/rotisserie/eris"
	"strings"
)

type GoogleTokenRepository interface {
	FindByEmail(email string) (*GoogleToken, error)
	Save(entity *GoogleToken) error
	Update(entity *GoogleToken) error
}

type PgGoogleTokenRepository struct {
	connection *sql.DB
}

func NewPgGoogleTokenRepository(connection *sql.DB) *PgGoogleTokenRepository {
	return &PgGoogleTokenRepository{
		connection: connection,
	}
}

func (r *PgGoogleTokenRepository) FindByEmail(email string) (*GoogleToken, error) {
	statement, err := r.connection.Prepare("SELECT * FROM pi_google_integration WHERE gi_user_email = $1;")

	if err != nil {
		return nil, eris.New(err.Error())
	}

	var (
		gi_id                   string
		gi_user_email           string
		gi_google_access_token  string
		gi_google_token_type    string
		gi_google_refresh_token string
		gi_google_token_expiry  string
		gi_created_at           string
		gi_updated_at           string
	)

	if scanErr := statement.QueryRow(email).Scan(
		&gi_id,
		&gi_user_email,
		&gi_google_access_token,
		&gi_google_token_type,
		&gi_google_refresh_token,
		&gi_google_token_expiry,
		&gi_created_at,
		&gi_updated_at,
	); scanErr != nil {
		if errors.Is(scanErr, sql.ErrNoRows) {
			return nil, eris.Wrap(GoogleTokenNotFoundError, "")
		}

		return nil, eris.New(scanErr.Error())
	}

	return &GoogleToken{
		Id:           gi_id,
		UserEmail:    gi_user_email,
		AccessToken:  gi_google_access_token,
		TokenType:    gi_google_token_type,
		RefreshToken: gi_google_refresh_token,
		Expiry:       gi_google_token_expiry,
		CreatedAt:    gi_created_at,
		UpdatedAt:    gi_updated_at,
	}, nil
}

func (r *PgGoogleTokenRepository) Save(entity *GoogleToken) error {
	var query strings.Builder

	query.WriteString(`INSERT INTO pi_google_integration `)
	query.WriteString(`(gi_id, gi_user_email, gi_google_access_token, gi_google_token_type, gi_google_refresh_token, gi_google_token_expiry, gi_created_at, gi_updated_at) `)
	query.WriteString(`VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`)

	_, err := r.connection.Exec(
		query.String(),
		entity.Id,
		entity.UserEmail,
		entity.AccessToken,
		entity.TokenType,
		entity.RefreshToken,
		entity.Expiry,
		entity.CreatedAt,
		entity.UpdatedAt,
	)

	if err != nil {
		return eris.New(err.Error())
	}

	return nil
}

func (r *PgGoogleTokenRepository) Update(entity *GoogleToken) error {
	var query strings.Builder

	query.WriteString(`UPDATE pi_google_integration SET `)
	query.WriteString(`gi_user_email = $1, gi_google_access_token = $2, gi_google_token_type = $3, gi_google_refresh_token = $4, gi_google_token_expiry = $5, gi_updated_at = $6 `)
	query.WriteString(`WHERE gi_id = $7;`)

	fmt.Println(query.String())
	_, err := r.connection.Exec(
		query.String(),
		entity.UserEmail,
		entity.AccessToken,
		entity.TokenType,
		entity.RefreshToken,
		entity.Expiry,
		entity.UpdatedAt,
		entity.Id,
	)

	if err != nil {
		return eris.New(err.Error())
	}

	return nil
}
