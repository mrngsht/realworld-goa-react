package internal

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
)

var Tx = DefaultTx

type TxStarter interface {
	BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
}

type TxEnder interface {
	Rollback() error
	Commit() error
}

func DefaultTx(ctx context.Context, txStarter TxStarter, txFunc func(ctx context.Context, txEnder TxEnder) error) (err error) {
	tx, err := txStarter.BeginTx(ctx, nil)
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
		} else {
			if err := tx.Commit(); err != nil {
				err = errors.WithStack(err)
			}
		}
	}()

	if err := txFunc(ctx, tx); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
