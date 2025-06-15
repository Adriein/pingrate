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
	userRepository repository.UserRepository,
	googleRepository repository.GoogleIntegrationRepository,
	googleApi *external.GoogleApi,
) *GoogleOauthCallbackService {
	return &GoogleOauthCallbackService{
		userRepository:   userRepository,
		googleRepository: googleRepository,
		googleApi:        googleApi,
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

	currentToken, tokenFindOneErr := s.googleRepository.FindOne(
		types.NewCriteria().Equal("gi_user_email", userEmail),
	)

	if tokenFindOneErr != nil {
		return tokenFindOneErr
	}

	if currentToken != nil {
		var refreshToken = ""

		if googleToken.RefreshToken != "" {
			refreshToken = googleToken.RefreshToken
		} else {
			refreshToken = currentToken.RefreshToken
		}

		mergedToken := &types.GoogleToken{
			Id:           currentToken.Id,
			UserEmail:    currentToken.UserEmail,
			AccessToken:  googleToken.AccessToken,
			TokenType:    currentToken.TokenType,
			RefreshToken: refreshToken,
			CreatedAt:    currentToken.CreatedAt,
			UpdatedAt:    googleToken.UpdatedAt,
		}

		if updateErr := s.googleRepository.Update(mergedToken); updateErr != nil {
			return updateErr
		}

		return nil
	}

	if saveErr := s.googleRepository.Save(googleToken); saveErr != nil {
		return saveErr
	}

	return nil
}
