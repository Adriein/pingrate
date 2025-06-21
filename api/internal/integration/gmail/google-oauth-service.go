package gmail

import (
	"github.com/adriein/pingrate/internal/shared/external"
)

type GoogleOauthService struct {
	googleApi *external.GoogleApi
}

func NewGoogleOauthService(
	googleApi *external.GoogleApi,
) *GoogleOauthService {
	return &GoogleOauthService{
		googleApi: googleApi,
	}
}

func (s *GoogleOauthService) Execute(userEmail string) string {
	return s.googleApi.GetAuthCodeUrlForUser(userEmail)
}
