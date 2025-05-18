package gmail

import (
	"github.com/adriein/pingrate/internal/shared/external"
)

type GoogleOauthCallbackService struct {
	googleApi *external.GoogleApi
}

func NewGoogleOauthCallbackService(
	googleApi *external.GoogleApi,
) *GoogleOauthCallbackService {
	return &GoogleOauthCallbackService{
		googleApi: googleApi,
	}
}

func (s *GoogleOauthCallbackService) Execute(userEmail string) string {
	return s.googleApi.GetAuthCodeUrlForUser(userEmail)
}
