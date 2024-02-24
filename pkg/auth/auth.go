package auth

import (
	"crypto/rand"
)

func NewCSRFToken() []byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	return key
}
