package database

import (
	"context"
	"database/sql"
	"embed"
	"encoding/json"
	"fmt"
	"time"

	"github.com/navikt/nada-pg-test/pkg/database/gensql"
	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type Repo struct {
	querier Querier
	db      *sql.DB
}

type Querier interface {
	gensql.Querier
	WithTx(tx *sql.Tx) *gensql.Queries
}

func New(dbConnDSN string) (*Repo, error) {
	db, err := sql.Open("postgres", dbConnDSN)
	if err != nil {
		return nil, fmt.Errorf("open sql connection: %w", err)
	}

	goose.SetBaseFS(embedMigrations)

	err = goose.Up(db, "migrations")
	if err != nil {
		return nil, err
	}

	return &Repo{
		querier: gensql.New(db),
		db:      db,
	}, nil
}

func (r *Repo) InsertData(ctx context.Context, data map[string]any) (time.Duration, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	start := time.Now()
	_, err = r.querier.InsertData(ctx, dataBytes)
	if err != nil {
		return 0, err
	}
	elapsed := time.Since(start)

	return elapsed, nil
}
