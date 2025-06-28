package auth

import (
	"github.com/adriein/pingrate/internal/user"
	"github.com/google/uuid"
	"time"
)

type Service struct {
	sessionRepository SessionRepository
	userRepository    user.UserRepository
}

func NewService(sessionRepository SessionRepository, userRepository user.UserRepository) *Service {
	return &Service{
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
	}
}

func (s *Service) CreateSession(request *LoginRequest) (*Session, error) {
	registeredUser, findUserErr := s.userRepository.FindByEmail(request.Email)

	if findUserErr != nil {
		return nil, findUserErr
	}

	isCorrect := registeredUser.CheckPassword(request.Password)

	if !isCorrect {
		return nil, user.UserIncorrectPasswordError
	}

	session := &Session{
		Id:        uuid.New().String(),
		Email:     registeredUser.Email,
		CreatedAt: time.Now().UTC().Format(time.DateTime),
		UpdatedAt: time.Now().UTC().Format(time.DateTime),
	}

	if sessionErr := s.sessionRepository.Save(session); sessionErr != nil {
		return nil, sessionErr
	}

	return session, nil
}
