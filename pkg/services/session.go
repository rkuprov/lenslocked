package services

import (
	"context"
	"lenslocked/pkg/auth"
	"lenslocked/pkg/store"
	"time"
)

type Session struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	TokenHash string `json:"token_hash"`
}

type SessionService struct {
	ctx context.Context
	db  *store.Store
}

func (s *SessionService) Create(userID int) (*Session, error) {
	token, err := auth.NewSessionToken()
	if err != nil {
		return nil, err

	}
	th, err := hashToString(token)
	if err != nil {
		return nil, err

	}
	_, err = s.db.Psql.Exec(s.ctx, SessionAddSQL, userID, th, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *SessionService) ValidateUserSession(userID int, token string) error {

	return nil
}
