package user

import (
	"github.com/adriein/pingrate/internal/shared/repository"
	"github.com/adriein/pingrate/internal/shared/types"
)

type LoginUserService struct {
	repository repository.UserRepository
}

func NewLoginUserService(repository repository.UserRepository) *LoginUserService {
	return &LoginUserService{
		repository: repository,
	}
}

func (s *LoginUserService) Execute(email string, inputPassword string) error {
	criteria := types.NewCriteria().Equal("us_email", email)

	user, findOneErr := s.repository.FindOne(criteria)

	if findOneErr != nil {
		return findOneErr
	}

	isCorrect := user.CheckPassword(inputPassword)

	if !isCorrect {
		return types.UserIncorrectPasswordError
	}

	return nil
}
