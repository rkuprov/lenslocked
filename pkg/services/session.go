package services

import (
	"context"
	"fmt"
	"lenslocked/pkg/auth"
	"lenslocked/pkg/datamodel"
	"lenslocked/pkg/store"
	"log"
	"time"
)

type Session struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Token     string `json:"token"`
	TokenHash string `json:"token_hash"`
}

type SessionService struct {
	ctx context.Context
	db  *store.Store
}

func NewSessionService(ctx context.Context, db *store.Store) *SessionService {
	return &SessionService{
		ctx: ctx,
		db:  db,
	}
}

func (s *SessionService) Create(userID int) (*Session, error) {
	token, err := auth.NewSessionToken()
	if err != nil {
		return nil, err

	}
	th := auth.SHAHash(token)
	var id int
	err = s.db.Psql.QueryRow(s.ctx,
		SessionAddSQL,
		userID,
		th,
		time.Now().UTC().Format(time.RFC3339),
	).Scan(&id)
	if err != nil {
		return nil, err
	}
	log.Default().Println("session id", id)

	return &Session{
		ID:        id,
		UserID:    userID,
		Token:     token,
		TokenHash: th,
	}, nil
}

func (s *SessionService) ValidateUserSession(userID int, token string) error {

	return nil
}

func (s *SessionService) GetUserForSession(token string) (*datamodel.User, error) {
	user := &datamodel.User{}
	tokenHash := auth.SHAHash(token)
	err := s.db.Psql.QueryRow(s.ctx, SessionGetUserSQL, tokenHash).Scan(&user.ID, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("could not find session for user %d: %w", user.ID, err)
	}

	return user, nil
}

func (s *SessionService) Delete(token string) error {
	_, err := s.db.Psql.Exec(s.ctx, SessionDeleteSQL, auth.SHAHash(token))
	if err != nil {
		return fmt.Errorf("could not delete session: %w", err)
	}

	return nil
}
