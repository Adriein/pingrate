package repository

import "github.com/adriein/pingrate/internal/shared/types"

type UserRepository interface {
	Find(criteria types.Criteria) ([]types.User, error)
	FindOne(criteria types.Criteria) (*types.User, error)
	Save(entity *types.User) error
	Update(entity *types.User) error
}
