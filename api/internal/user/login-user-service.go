package user

import (
	"github.com/adriein/pingrate/internal/shared/repository"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/google/uuid"
	"time"
)

type LoginUserService struct {
	userRepository    repository.UserRepository
	sessionRepository repository.SessionRepository
}

func NewLoginUserService(
	userRepository repository.UserRepository,
	sessionRepository repository.SessionRepository,
) *LoginUserService {
	return &LoginUserService{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}

func (s *LoginUserService) Execute(email string, inputPassword string) (*types.Session, error) {
	criteria := types.NewCriteria().Equal("us_email", email)

	user, findOneErr := s.userRepository.FindOne(criteria)

	if findOneErr != nil {
		return nil, findOneErr
	}

	isCorrect := user.CheckPassword(inputPassword)

	if !isCorrect {
		return nil, types.UserIncorrectPasswordError
	}

	session := &types.Session{
		Id:        uuid.New().String(),
		Email:     user.Email,
		CreatedAt: time.Now().UTC().Format(time.DateTime),
		UpdatedAt: time.Now().UTC().Format(time.DateTime),
	}

	if sessionErr := s.sessionRepository.Save(session); sessionErr != nil {
		return nil, sessionErr
	}

	return session, nil
}
