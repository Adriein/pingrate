package user

import (
	"errors"
	"time"
)

type Service struct {
	repository UserRepository
}

func NewService(repository UserRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateUser(request CreateUserRequest) error {
	user := &User{
		Id:        request.Id,
		Email:     request.Email,
		Password:  request.Password,
		CreatedAt: time.Now().UTC().Format(time.DateTime),
		UpdatedAt: time.Now().UTC().Format(time.DateTime),
	}

	if securePassErr := user.SecurePassword(); securePassErr != nil {
		return securePassErr
	}

	_, findOneErr := s.repository.FindByEmail(user.Email)

	if findOneErr == nil {
		return UserAlreadyExistError
	}

	if !errors.Is(findOneErr, UserNotFoundError) {
		return findOneErr
	}

	if err := s.repository.Save(user); err != nil {
		return err
	}

	return nil
}
