package auth

import (
	"crypto/rand"
	"encoding/base64"
)

const tokenStrength = 32

func NewCSRFToken() []byte {
	key := make([]byte, tokenStrength)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	return key
}

func NewSessionToken() (string, error) {
	key := make([]byte, tokenStrength)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(key), nil
}
