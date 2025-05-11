package oauth

import "github.com/adriein/pingrate/internal/shared/external"

type GoogleOauthControllerService struct {
	googleApi *external.GoogleApi
}

func (s *GoogleOauthControllerService) Execute(userId string) string {
	return s.googleApi.GetAuthCodeUrlForUser(userId)
}
