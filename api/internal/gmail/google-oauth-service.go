package gmail

import (
	"github.com/adriein/pingrate/internal/shared/external"
	"github.com/adriein/pingrate/internal/shared/repository"
)

type GoogleOauthService struct {
	googleApi *external.GoogleApi
}

func NewGoogleOauthService(
	repository repository.UserRepository,
	googleApi *external.GoogleApi,
) *GoogleOauthService {
	return &GoogleOauthService{
		googleApi: googleApi,
	}
}

func (s *GoogleOauthService) Execute(userEmail string) string {
	return s.googleApi.GetAuthCodeUrlForUser(userEmail)
}
