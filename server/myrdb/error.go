package myrdb

import (
	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func IsErrNoRows(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}

func IsErrUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505" // https://www.postgresql.jp/document/8.0/html/errcodes-appendix.html
}
