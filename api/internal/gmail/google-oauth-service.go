package gmail

import (
	"github.com/adriein/pingrate/internal/shared/external"
	"github.com/adriein/pingrate/internal/shared/repository"
)

type GoogleOauthCallbackService struct {
	googleApi *external.GoogleApi
}

func NewGoogleOauthCallbackService(
	repository repository.UserRepository,
	googleApi *external.GoogleApi,
) *GoogleOauthCallbackService {
	return &GoogleOauthCallbackService{
		googleApi: googleApi,
	}
}

func (s *GoogleOauthCallbackService) Execute(userEmail string) string {
	return s.googleApi.GetAuthCodeUrlForUser(userEmail)
}
