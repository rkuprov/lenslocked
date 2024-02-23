package services

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
	pwHash, err := hashPassword(password)
	if err != nil {
		return 0, fmt.Errorf("could not hash password: %w", err)

	}
	err = s.db.Psql.QueryRow(s.ctx, UserAddSQL,
		strings.ToLower(email),
		string(pwHash),
		time.Now().UTC().Format(time.RFC3339)).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not create user: %w", err)
	}
	return id, nil
}

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
