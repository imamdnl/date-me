package common

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IBaseRepository interface {
	SetContext(ctx context.Context)
	GetContext() context.Context

	// GetMyConnection will return MySql connection
	GetMyConnection() *sql.DB

	// GetPgConnection will return Postgres connection
	GetPgConnection() *pgxpool.Pool
}

// BaseRepository database
type BaseRepository struct {
	PostgresSql *pgxpool.Pool
	Redis       *redis.Client

	ctx context.Context
}

func NewBaseRepository(dbp *pgxpool.Pool, redis *redis.Client) *BaseRepository {
	return &BaseRepository{
		PostgresSql: dbp,
		Redis:       redis,
	}
}

func (repo *BaseRepository) SetContext(ctx context.Context) {
	repo.ctx = ctx
}

func (repo *BaseRepository) GetContext() context.Context {
	return repo.ctx
}

func (repo *BaseRepository) GetPgConnection() *pgxpool.Pool {
	return repo.PostgresSql
}
