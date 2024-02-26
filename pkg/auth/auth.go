package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type CookieType = string

const (
	tokenStrength     = 32
	CookieTypeSession = "session"
	CookieTypeCSRF    = "csrf"
)

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

func newCookie(name CookieType, value string) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: true,
		Path:     "/",
	}
}

func SetCookie(w http.ResponseWriter, name CookieType, value string) {
	http.SetCookie(w, newCookie(name, value))
}

func HashToBytes(toHash string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(toHash), bcrypt.DefaultCost)
}

func HashToString(toHash string) (string, error) {
	bts, err := HashToBytes(toHash)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bts), nil
}

func SHAHash(toHash string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(toHash)))
}
