package crypto

import (
	"crypto/rand"
	"math/big"
)

const (
	saltSize       = 16
	sessionKeySize = 120
	codeSize       = 120

	saltChars       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	sessionKeyChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	codeChars       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func Salt() (string, error) {
	salt := make([]rune, saltSize)
	for i := range salt {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(saltChars))))
		if err != nil {
			return "", err
		}
		salt[i] = rune(saltChars[n.Int64()])
	}
	return string(salt), nil
}
