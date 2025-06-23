package gmail

import "errors"

type Service struct {
	repository GoogleTokenRepository
	googleApi  GoogleApi
}

func NewService(repository GoogleTokenRepository, googleApi GoogleApi) *Service {
	return &Service{
		repository: repository,
		googleApi:  googleApi,
	}
}

func (s *Service) GetGmailOauthLink(email string) string {
	return s.googleApi.GetAuthCodeUrlForUser(email)
}

func (s *Service) ExchangeGoogleToken(email string, code string) error {
	googleToken, exchangeTokenErr := s.googleApi.ExchangeToken(email, code)

	if exchangeTokenErr != nil {
		return exchangeTokenErr
	}

	currentToken, tokenFindOneErr := s.repository.FindByEmail(email)

	if errors.Is(tokenFindOneErr, GoogleTokenNotFoundError) {
		if saveErr := s.repository.Save(googleToken); saveErr != nil {
			return saveErr
		}

		return nil
	}

	if tokenFindOneErr != nil {
		return tokenFindOneErr
	}

	var refreshToken = ""

	if googleToken.RefreshToken != "" {
		refreshToken = googleToken.RefreshToken
	} else {
		refreshToken = currentToken.RefreshToken
	}

	mergedToken := &GoogleToken{
		Id:           currentToken.Id,
		UserEmail:    currentToken.UserEmail,
		AccessToken:  googleToken.AccessToken,
		TokenType:    currentToken.TokenType,
		RefreshToken: refreshToken,
		Expiry:       googleToken.Expiry,
		CreatedAt:    currentToken.CreatedAt,
		UpdatedAt:    googleToken.UpdatedAt,
	}

	if updateErr := s.repository.Update(mergedToken); updateErr != nil {
		return updateErr
	}

	return nil
}
