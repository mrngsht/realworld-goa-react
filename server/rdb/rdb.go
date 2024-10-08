package rdb

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/mrngsht/realworld-goa-react/rdb/internal"

	_ "github.com/lib/pq"
)

const (
	LocalConnectionString = "host=localhost user=postgres password=postgres dbname=realworld sslmode=disable timezone=UTC"
)

func OpenLocalRDB() (*sql.DB, error) {
	return sql.Open("postgres", LocalConnectionString)
}

func IsErrNoRows(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

func Tx(ctx context.Context, db *sql.DB, txFunc func(ctx context.Context, tx *sql.Tx) error) (err error) {
	return internal.Tx(ctx, db, txFunc)
}
