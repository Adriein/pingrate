package gmail

import (
	"context"
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/google/uuid"
	"github.com/rotisserie/eris"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"os"
	"time"
)

type GoogleApi struct{}

func NewGoogleApi() *GoogleApi {
	return &GoogleApi{}
}

func (g *GoogleApi) GetOauth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv(constants.GoogleClientId),
		ClientSecret: os.Getenv(constants.GoogleClientSecret),
		RedirectURL:  "http://localhost:4000/api/v1/integrations/gmail/oauth-callback",
		Endpoint:     google.Endpoint,
		Scopes:       []string{gmail.GmailReadonlyScope},
	}
}

func (g *GoogleApi) GetAuthCodeUrlForUser(userEmail string) string {
	config := g.GetOauth2Config()

	return config.AuthCodeURL(userEmail, oauth2.AccessTypeOffline, oauth2.ApprovalForce, oauth2.S256ChallengeOption(userEmail))
}

func (g *GoogleApi) ExchangeToken(state string, code string) (*GoogleToken, error) {
	ctx := context.Background()
	config := g.GetOauth2Config()

	token, exchangeErr := config.Exchange(ctx, code, oauth2.VerifierOption(state))

	if exchangeErr != nil {
		return nil, eris.New(exchangeErr.Error())
	}

	googleToken := &GoogleToken{
		Id:           uuid.New().String(),
		UserEmail:    state,
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry.Format(time.DateTime),
		CreatedAt:    time.Now().UTC().Format(time.DateTime),
		UpdatedAt:    time.Now().UTC().Format(time.DateTime),
	}

	return googleToken, nil
}

func (g *GoogleApi) GmailClient(userToken *GoogleToken) (*gmail.Service, error) {
	ctx := context.Background()

	config := g.GetOauth2Config()

	expiry, parseErr := time.Parse(time.DateTime, userToken.Expiry)

	if parseErr != nil {
		return nil, eris.Wrap(parseErr, "failed to parse token expiry time")
	}

	token := &oauth2.Token{
		AccessToken:  userToken.AccessToken,
		TokenType:    userToken.TokenType,
		RefreshToken: userToken.RefreshToken,
		Expiry:       expiry,
	}

	client, newServiceErr := gmail.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))

	if newServiceErr != nil {
		return nil, eris.New(newServiceErr.Error())
	}

	return client, nil
}
