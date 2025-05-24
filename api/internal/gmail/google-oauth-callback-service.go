package gmail

import (
	"github.com/adriein/pingrate/internal/shared/external"
	"github.com/adriein/pingrate/internal/shared/repository"
	"github.com/adriein/pingrate/internal/shared/types"
)

type GoogleOauthCallbackService struct {
	userRepository   repository.UserRepository
	googleRepository repository.GoogleIntegrationRepository
	googleApi        *external.GoogleApi
}

func NewGoogleOauthCallbackService(
	repository repository.UserRepository,
	googleRepository repository.GoogleIntegrationRepository,
	googleApi *external.GoogleApi,
) *GoogleOauthCallbackService {
	return &GoogleOauthCallbackService{
		googleApi: googleApi,
	}
}

func (s *GoogleOauthCallbackService) Execute(userEmail string, code string) error {
	_, findOneErr := s.userRepository.FindOne(types.NewCriteria().Equal("us_email", userEmail))

	if findOneErr != nil {
		return findOneErr
	}

	googleToken, exchangeTokenErr := s.googleApi.ExchangeToken(userEmail, code)

	if exchangeTokenErr != nil {
		return exchangeTokenErr
	}

	if saveErr := s.googleRepository.Save(googleToken); saveErr != nil {
		return saveErr
	}

	return nil
}
