package gmail

import (
	"fmt"
	"github.com/adriein/pingrate/internal/shared/external"
	"github.com/adriein/pingrate/internal/shared/repository"
	"github.com/adriein/pingrate/internal/shared/types"
)

type GetGmailService struct {
	googleApi  *external.GoogleApi
	repository repository.GoogleIntegrationRepository
}

func NewGetGmailService(
	googleApi *external.GoogleApi,
	repository repository.GoogleIntegrationRepository,
) *GetGmailService {
	return &GetGmailService{
		googleApi:  googleApi,
		repository: repository,
	}
}

func (s *GetGmailService) Execute(userEmail string) error {
	googleToken, findOneErr := s.repository.FindOne(
		types.NewCriteria().Equal("gi_user_email", userEmail),
	)

	if findOneErr != nil {
		return findOneErr
	}

	gmailClient, gmailClientErr := s.googleApi.GmailClient(googleToken)

	if gmailClientErr != nil {
		return gmailClientErr
	}

	do, err := gmailClient.Users.Messages.List("me").Q("after:2014/01/01 before:2014/02/01").Do()
	if err != nil {
		return err
	}

	fmt.Println(do.Messages)
	return nil
}
