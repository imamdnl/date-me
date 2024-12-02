package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func PostgreSql(ctx context.Context, logger *zap.Logger) *pgxpool.Pool {
	logger.Info("initializing postgresql connection")
	cfg, err := NewPostgresConfig()
	if err != nil {
		logger.Error("error while initializing database connection", zap.Error(err))
		panic(err)
	}

	// Build the connection string based on the target server type and the master and slave host/port values
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s&pool_max_conns=250&pool_min_conns=1&pool_max_conn_lifetime=10s&pool_max_conn_idle_time=5s",
		cfg.GetUsername(),
		cfg.GetPassword(),
		cfg.GetHost(),
		cfg.GetPort(),
		cfg.GetDatabase(),
		cfg.GetSchema(),
	)

	// Open a database connection using the generated connection string
	dbase, err := pgxpool.New(ctx, connString)
	if err != nil {
		logger.Error("error when try to open database", zap.Error(err))
		panic(err)
	}

	return dbase
}
