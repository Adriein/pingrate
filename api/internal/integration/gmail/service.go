package gmail

import (
	"errors"
)

type Service struct {
	repository GoogleTokenRepository
	googleApi  *GoogleApi
}

func NewService(repository GoogleTokenRepository, googleApi *GoogleApi) *Service {
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

func (s *Service) GetGmailInbox(email string) ([]*Gmail, error) {
	token, tokenFindOneErr := s.repository.FindByEmail(email)

	if tokenFindOneErr != nil {
		return nil, tokenFindOneErr
	}

	gmailClient, gmailClientErr := s.googleApi.GmailClient(token)

	if gmailClientErr != nil {
		return nil, gmailClientErr
	}

	response, getMessagesErr := gmailClient.Users.Messages.
		List("me").
		Q("after:2012/01/01 before:2012/07/01").
		Do()

	if getMessagesErr != nil {
		return nil, getMessagesErr
	}

	var rawResult []*Gmail

	for _, message := range response.Messages {
		fullMessage, getMessageErr := gmailClient.Users.Messages.Get("me", message.Id).Do()

		if getMessageErr != nil {
			return nil, getMessageErr
		}

		mail, mailErr := NewMail(fullMessage)

		if mailErr != nil {
			return nil, mailErr
		}

		rawResult = append(rawResult, mail)
	}

	return rawResult, nil
	//return s.mergeThreads(rawResult), nil
}

func (s *Service) mergeThreads(emails []*Gmail) []*Gmail {
	threadMap := make(map[string]*Gmail)

	for _, email := range emails {
		if merged, exists := threadMap[email.ThreadId]; exists {
			merged.Body += "\n" + email.Body

			continue
		}

		threadMap[email.ThreadId] = &Gmail{
			Id:       email.Id,
			ThreadId: email.ThreadId,
			Body:     email.Body,
		}

	}

	var result []*Gmail

	for _, mergedEmail := range threadMap {
		result = append(result, mergedEmail)
	}

	return result
}
