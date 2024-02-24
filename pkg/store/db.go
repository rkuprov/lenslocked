package store

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"lenslocked/cfg"
	"lenslocked/migrations"
)

type Store struct {
	Psql *pgxpool.Pool
}

func NewStore(pCfg cfg.Postgres) (*Store, error) {
	psql, err := pgxpool.New(context.Background(), fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s pool_max_conns=%s",
		pCfg.User,
		pCfg.Password,
		pCfg.Host,
		pCfg.Port,
		pCfg.DBName,
		pCfg.SSLMode,
		pCfg.PoolMaxConns))
	if err != nil {
		return nil, err
	}
	return &Store{
		Psql: psql,
	}, nil
}

func (s *Store) Setup(ctx context.Context) error {
	_, err := s.Psql.Exec(ctx, migrations.UserTable)
	if err != nil {
		return err
	}
	_, err = s.Psql.Exec(ctx, migrations.SessionTable)
	if err != nil {
		return err
	}
	return nil
}
