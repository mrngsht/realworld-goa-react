package myrdb

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/mrngsht/realworld-goa-react/config"
	"github.com/mrngsht/realworld-goa-react/myrdb/internal"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func OpenLocalRDB() (*sql.DB, error) {
	return sql.Open("postgres", config.C.RDBConnectionString)
}

func IsErrNoRows(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

const pqUniqueViolationCode = pq.ErrorCode("23505")

func IsErrUniqueViolation(err error) bool {
	// ref: https://github.com/go-gorm/gorm/issues/4135#issuecomment-790584782
	var pqErr *pq.Error
	return errors.As(err, &pqErr) && pqErr.Code == pqUniqueViolationCode
}

func Tx(ctx context.Context, db *sql.DB, txFunc func(ctx context.Context, tx *sql.Tx) error) (err error) {
	return internal.Tx(ctx, db, txFunc)
}
