package valueobject

import (
	"github.com/nnniyaz/blog/pkg/core"
	"github.com/nnniyaz/blog/pkg/crypto"
	"github.com/nnniyaz/blog/pkg/hash"
	"unicode/utf8"
)

const (
	minPasswordLength = 6
	maxPasswordLength = 32
)

var (
	ErrEmptyPassword    = core.NewI18NError(core.EINVALID, core.TXT_EMPTY_PASSWORD)
	ErrPasswordTooShort = core.NewI18NError(core.EINVALID, core.TXT_PASSWORD_TOO_SHORT)
	ErrPasswordTooLong  = core.NewI18NError(core.EINVALID, core.TXT_PASSWORD_TOO_LONG)
)

type Password struct {
	hash string
	salt string
}

func NewPassword(password string) (Password, error) {
	if password == "" {
		return Password{}, ErrEmptyPassword
	}

	if utf8.RuneCountInString(password) < minPasswordLength {
		return Password{}, ErrPasswordTooShort
	}
	if utf8.RuneCountInString(password) > maxPasswordLength {
		return Password{}, ErrPasswordTooLong
	}
	passwordSalt, err := crypto.Salt()
	if err != nil {
		return Password{}, err
	}

	passwordHash, err := hash.PasswordHash(password, passwordSalt)
	if err != nil {
		return Password{}, err
	}

	return Password{hash: passwordHash, salt: passwordSalt}, nil
}

func (p Password) GetHash() string {
	return p.hash
}

func (p Password) GetSalt() string {
	return p.salt
}

func (p Password) Compare(password string) bool {
	return hash.ComparePassword(password, p.hash, p.salt)
}

func UnmarshalPasswordFromDatabase(hash, salt string) Password {
	return Password{hash: hash, salt: salt}
}
