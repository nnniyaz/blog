package hash

import "github.com/alexedwards/argon2id"

func PasswordHash(password, salt string) (string, error) {
	return argon2id.CreateHash(password+salt, argon2id.DefaultParams)
}

func ComparePassword(password, hash, salt string) bool {
	ok, err := argon2id.ComparePasswordAndHash(password+salt, hash)
	if err != nil {
		return false
	}
	return ok
}
