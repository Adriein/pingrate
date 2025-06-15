package repository

import "github.com/adriein/pingrate/internal/shared/types"

type GoogleIntegrationRepository interface {
	Find(criteria types.Criteria) ([]types.GoogleIntegration, error)
	FindOne(criteria types.Criteria) (*types.GoogleIntegration, error)
	Save(entity *types.GoogleIntegration) error
	Update(entity *types.GoogleIntegration) error
}
