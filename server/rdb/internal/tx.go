package internal

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
)

var Tx = DefaultTx

func DefaultTx(ctx context.Context, db *sql.DB, txFunc func(ctx context.Context, tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Errorf("panic error %v", panicErr)
		}
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = errors.Join(err, rollbackErr)
			}
		}
	}()

	if err := txFunc(ctx, tx); err != nil {
		return errors.WithStack(err)
	}

	if err := tx.Commit(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
