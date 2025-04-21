package user

import (
	"github.com/adriein/pingrate/internal/shared/repository"
	"github.com/adriein/pingrate/internal/shared/types"
)

type CreateUserService struct {
	repository repository.UserRepository
}

func NewCreateUserService(repository repository.UserRepository) *CreateUserService {
	return &CreateUserService{
		repository: repository,
	}
}

func (s *CreateUserService) Execute(user *types.User) error {
	if securePassErr := user.SecurePassword(); securePassErr != nil {
		return securePassErr
	}

	if err := s.repository.Save(user); err != nil {
		return err
	}

	return nil
}
