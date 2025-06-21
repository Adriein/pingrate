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

	response, getMessagesErr := gmailClient.Users.Messages.List("me").Q("after:2012/01/01 before:2012/02/01").Do()

	if getMessagesErr != nil {
		return getMessagesErr
	}

	for _, message := range response.Messages {
		fullMessage, getMessageErr := gmailClient.Users.Messages.Get("me", message.Id).Do()

		if getMessageErr != nil {
			return getMessageErr
		}

		a, _ := types.GmailFromMessage(fullMessage)

		fmt.Println(a)
	}

	fmt.Println(response.Messages)
	return nil
}
