package types

import (
	"encoding/base64"
	"github.com/rotisserie/eris"
	"google.golang.org/api/gmail/v1"
)

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
