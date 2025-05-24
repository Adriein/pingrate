package repository

import "github.com/adriein/pingrate/internal/shared/types"

type GoogleIntegrationRepository interface {
	Find(criteria types.Criteria) ([]types.GoogleToken, error)
	FindOne(criteria types.Criteria) (*types.GoogleToken, error)
	Save(entity *types.GoogleToken) error
	Update(entity *types.GoogleToken) error
}
