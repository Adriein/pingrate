package user

import (
	"github.com/adriein/pingrate/internal/shared/repository"
	"github.com/adriein/pingrate/internal/shared/types"
)

type WhoAmIService struct {
	repository repository.UserRepository
}

func NewWhoAmIService(repository repository.UserRepository) *WhoAmIService {
	return &WhoAmIService{
		repository: repository,
	}
}

func (s *WhoAmIService) Execute(email string) error {
	criteria := types.NewCriteria().Equal("us_email", email)

	_, findOneErr := s.repository.FindOne(criteria)

	if findOneErr != nil {
		return findOneErr
	}

	return nil
}
