package email

import (
	"github.com/nnniyaz/blog/pkg/core"
	"net/mail"
)

var (
	ErrorEmailInvalid = core.NewI18NError(core.EINVALID, core.TXT_INVALID_EMAIL)
)

type Email string

func NewEmail(email string) (Email, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return "", ErrorEmailInvalid
	}
	return Email(email), nil
}

func (e Email) String() string {
	return string(e)
}

func (e Email) Validate() error {
	if _, err := mail.ParseAddress(e.String()); err != nil {
		return ErrorEmailInvalid
	}
	return nil
}
