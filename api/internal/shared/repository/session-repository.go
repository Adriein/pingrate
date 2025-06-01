package repository

import "github.com/adriein/pingrate/internal/shared/types"

type SessionRepository interface {
	Find(criteria types.Criteria) ([]types.Session, error)
	FindOne(criteria types.Criteria) (*types.Session, error)
	Save(entity *types.Session) error
	Update(entity *types.Session) error
}
