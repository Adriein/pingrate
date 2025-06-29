package gmail

import (
	"encoding/base64"
	"github.com/adriein/pingrate/internal/shared/utils"
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
	Id       string `json:"id"`
	ThreadId string `json:"threadId"`
	Body     string `json:"body"`
}

func NewMail(message *gmail.Message) (*Gmail, error) {
	if message.Payload.MimeType == "text/plain" {
		byteMessageBody, decodeBase64Err := base64.URLEncoding.DecodeString(message.Payload.Body.Data)

		if decodeBase64Err != nil {
			return nil, eris.New(decodeBase64Err.Error())
		}

		return &Gmail{
			Id:       message.Id,
			ThreadId: message.ThreadId,
			Body:     string(byteMessageBody),
		}, nil
	}

	decodedBody, decodeMultipartBodyErr := decodeMultipartBody(message)

	if decodeMultipartBodyErr != nil {
		return nil, decodeMultipartBodyErr
	}

	return &Gmail{
		Id:       message.Id,
		ThreadId: message.ThreadId,
		Body:     *decodedBody,
	}, nil
}

func decodeMultipartBody(message *gmail.Message) (*string, error) {
	var result string

	queue := utils.NewQueue[*gmail.MessagePart](message.Payload.Parts...)

	for !queue.IsEmpty() {
		part, err := queue.Dequeue()

		if err != nil {
			return nil, err
		}

		if part.MimeType == "text/plain" {
			byteMessageBody, decodeBase64Err := base64.URLEncoding.DecodeString(part.Body.Data)

			if decodeBase64Err != nil {
				return nil, eris.New(decodeBase64Err.Error())
			}

			result += "\n" + string(byteMessageBody)

			continue
		}

		queue.Enqueue(part.Parts...)
	}

	return &result, nil
}
