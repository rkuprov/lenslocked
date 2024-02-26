package store

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"lenslocked/cfg"
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
