package services

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"lenslocked/pkg/auth"
	"lenslocked/pkg/store"
	"strings"
	"time"
)

type UserService struct {
	ctx context.Context
	db  *store.Store
}

func NewUserService(ctx context.Context, db *store.Store) *UserService {
	return &UserService{
		ctx: ctx,
		db:  db,
	}
}

func (s *UserService) Create(email, password string) (int, error) {
	var id int
	pwHash, err := auth.HashToString(password)
	if err != nil {
		return 0, fmt.Errorf("could not HashToBytes password: %w", err)

	}
	err = s.db.Psql.QueryRow(s.ctx, UserAddSQL,
		strings.ToLower(email),
		pwHash,
		time.Now().UTC().Format(time.RFC3339)).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not create user: %w", err)
	}
	return id, nil
}

func (s *UserService) Authenticate(email, password string) (int, error) {
	var pwHash string
	var id int
	err := s.db.Psql.QueryRow(s.ctx, UserGetPasswordHashSQL, strings.ToLower(email)).Scan(&id, &pwHash)
	if err != nil {
		return 0, fmt.Errorf("could not find user: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(pwHash), []byte(password))
	if err != nil {
		bts, _ := auth.HashToBytes(password)
		return 0, fmt.Errorf("user pwd %s\n user-HashToBytes: %s\n required HashToBytes: %s\n: %w", password, pwHash, string(bts), err)
	}

	return id, nil
}
