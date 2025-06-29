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

func NewMail(message *gmail.Message) (*Gmail, error) {
	if message.Payload.MimeType == "text/plain" {
		byteMessageBody, decodeBase64Err := base64.StdEncoding.DecodeString(message.Payload.Body.Data)

		if decodeBase64Err != nil {
			return nil, eris.New(decodeBase64Err.Error())
		}

		return &Gmail{
			Body: string(byteMessageBody),
		}, nil
	}

	decodedBody, decodeMultipartBodyErr := decodeMultipartBody(message)

	if decodeMultipartBodyErr != nil {
		return nil, decodeMultipartBodyErr
	}

	return &Gmail{Body: *decodedBody}, nil
}

func decodeMultipartBody(message *gmail.Message) (*string, error) {
	var queue []*gmail.MessagePart
	var result string

	queue = append(queue, message.Payload.Parts...)

	for len(queue) > 0 {
		part := queue[0]

		if part.MimeType == "text/plain" {
			byteMessageBody, decodeBase64Err := base64.StdEncoding.DecodeString(part.Body.Data)

			if decodeBase64Err != nil {
				return nil, eris.New(decodeBase64Err.Error())
			}

			result += "\n" + string(byteMessageBody)

			continue
		}

		queue = append(queue, part.Parts...)
	}

	return &result, nil
}
