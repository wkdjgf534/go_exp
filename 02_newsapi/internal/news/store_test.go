package news_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	pgtc "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/uptrace/bun"

	"newsapi/internal/postgres"
)

func createTestContainer(ctx context.Context) (ctr *pgtc.PostgresContainer, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return ctr, fmt.Errorf("working dir: %w", err)
	}

	sqlScripts := wd + "/testdata/sql/store.sql"

	ctr, err = pgtc.Run(
		ctx,
		"postgres:16-alpine",
		pgtc.WithInitScripts(sqlScripts),
		pgtc.WithDatabase("postgres"),
		pgtc.WithUsername("postgres"),
		pgtc.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second),
		),
	)

	if err != nil {
		return ctr, fmt.Errorf("run container: %w", err)
	}

	return ctr, nil
}

type DBCleanupFunc func(ctx context.Context) error

func createTestDB(ctx context.Context) (*bun.DB, DBCleanupFunc, error) {
	ctr, err := createTestContainer(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("create test container: %w", err)
	}

	p, err := ctr.MappedPort(ctx, nat.Port("5432/tcp"))
	if err != nil {
		return nil, nil, fmt.Errorf("mapped port: %w", err)
	}

	db, err := postgres.NewDB(&postgres.Config{
		Host:     "localhost",
		Debug:    true,
		DBName:   "postgres",
		User:     "postgres",
		Password: "postgres",
		Port:     p.Port(),
		SSLMode:  "disabled",
	})
	if err != nil {
		return nil, nil, fmt.Errorf("new db: %w", err)
	}

	cf := func(ctx context.Context) error {
		if err := db.Close(); err != nil {
			return fmt.Errorf("db close: %w", err)
		}
		if err := ctr.Terminate(ctx); err != nil {
			return fmt.Errorf("container terminate: %ws", err)
		}

		return nil
	}

	return db, cf, nil
}
