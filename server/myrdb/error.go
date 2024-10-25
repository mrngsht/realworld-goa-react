package myrdb

import (
	"database/sql"

	"github.com/cockroachdb/errors"

	"github.com/lib/pq"
)

func IsErrNoRows(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

const pqUniqueViolationCode = pq.ErrorCode("23505")

func IsErrUniqueViolation(err error) bool {
	// ref: https://github.com/go-gorm/gorm/issues/4135#issuecomment-790584782
	var pqErr *pq.Error
	return errors.As(err, &pqErr) && pqErr.Code == pqUniqueViolationCode
}
