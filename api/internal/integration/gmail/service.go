package gmail

import (
	"errors"
	"github.com/rotisserie/eris"
	"google.golang.org/api/gmail/v1"
	"sync"
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

	var emails []*Gmail
	pageToken := ""

	for {
		/*response, getMessagesErr := gmailClient.Users.Messages.
		List("me").
		MaxResults(250).
		Q("after:2025/01/01 before:2025/07/01").
		Do()*/

		call := gmailClient.Users.Messages.
			List("me").
			MaxResults(20).
			Q("after:2025/01/01 before:2025/02/01")

		if pageToken != "" {
			call.PageToken(pageToken)
		}

		response, getMessagesErr := call.Do()

		if getMessagesErr != nil {
			return nil, eris.New(getMessagesErr.Error())
		}

		ch := make(chan *ResultChannelResponse)
		var wg sync.WaitGroup

		for _, message := range response.Messages {
			wg.Add(1)

			go s.fetchFullEmail(gmailClient, message.Id, ch, &wg)
		}

		go func() {
			wg.Wait()
			close(ch)
		}()

		for result := range ch {
			if result.Err != nil {
				return nil, eris.New(result.Err.Error())
			}

			if result.Gmail.isEmpty() {
				continue
			}

			emails = append(emails, result.Gmail)
		}

		if response.NextPageToken == "" {
			break
		}

		pageToken = response.NextPageToken
	}

	return s.mergeThreads(emails), nil
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

func (s *Service) fetchFullEmail(
	client *gmail.Service,
	messageId string,
	ch chan<- *ResultChannelResponse,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	fullMessage, clientErr := client.Users.Messages.Get("me", messageId).Do()

	if clientErr != nil {
		ch <- &ResultChannelResponse{Gmail: nil, Err: eris.New(clientErr.Error())}
		return
	}

	mail, err := NewMail(fullMessage)

	if err != nil {
		ch <- &ResultChannelResponse{Gmail: nil, Err: err}
		return
	}

	ch <- &ResultChannelResponse{Gmail: mail, Err: nil}
}
