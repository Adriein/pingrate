package types

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"github.com/rotisserie/eris"
	"strings"
)

var (
	UserNotFoundError = eris.New("user not found")
)

var (
	UserAlreadyExistError = eris.New("user already exists")
)

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (u *User) generateSalt(size int) ([]byte, error) {
	salt := make([]byte, size)

	_, err := rand.Read(salt)

	if err != nil {
		return nil, eris.New(err.Error())
	}

	return salt, nil
}

func (u *User) hashPassword(password string, salt []byte) string {
	combined := append(salt, []byte(password)...)
	hash := sha256.Sum256(combined)
	return base64.StdEncoding.EncodeToString(hash[:])
}

func (u *User) SecurePassword() error {
	var result []string

	salt, err := u.generateSalt(16)

	if err != nil {
		return eris.New(err.Error())
	}

	hash := u.hashPassword(u.Password, salt)

	saltEncoded := base64.StdEncoding.EncodeToString(salt)

	result = append(result, hash, saltEncoded)

	hashedAndSalted := strings.Join(result, "$")

	u.Password = hashedAndSalted

	return nil
}
