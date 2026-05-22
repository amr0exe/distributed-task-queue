package app

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"tryit.me/internal/db"
)

type App struct {
	db *db.Postgres
}

type Config struct {
	DBString string
}

func New(cfg Config) (*App, error) {
	ctx := context.Background()
	//log.Print(cfg.DBString)

	postgres, err := db.NewPostgres(ctx, db.Config{
		DSN:             cfg.DBString,
		MaxConns:        20,
		MinConns:        5,
		MaxConnLifetime: time.Hour,
		MaxConnIdleTime: 30 * time.Minute,
	})
	if err != nil {
		return nil, err
	}

	return &App{
		db: postgres,
	}, nil
}

func (a *App) DB() *pgxpool.Pool {
	return a.db.Pool
}
