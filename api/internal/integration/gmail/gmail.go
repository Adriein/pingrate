package gmail

import (
	"encoding/base64"
	"github.com/rotisserie/eris"
	"google.golang.org/api/gmail/v1"
)

var (
	GoogleTokenNotFoundError = eris.New("google token not found")
)

type GoogleToken struct {
	Id           string
	UserEmail    string
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiry       string
	CreatedAt    string
	UpdatedAt    string
}
type Gmail struct {
	Body string
}

func GmailFromMessage(message *gmail.Message) (*Gmail, error) {
	if message.Payload.MimeType == "text/plain" {
		byteMessageBody, decodeBase64Err := base64.StdEncoding.DecodeString(message.Payload.Body.Data)

		if decodeBase64Err != nil {
			return nil, eris.New(decodeBase64Err.Error())
		}

		return &Gmail{
			Body: string(byteMessageBody),
		}, nil
	}

	return nil, nil
}
