package myrdb

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/mrngsht/realworld-goa-react/config"
	"github.com/mrngsht/realworld-goa-react/myrdb/sqlcgen"

	_ "github.com/lib/pq"
)

func OpenLocalRDB() (*rdb, error) {
	sqldb, err := sql.Open("postgres", config.C.RDBConnectionString)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &rdb{sqldb}, nil
}

type rdb struct {
	*sql.DB
}

var _ RDB = (*rdb)(nil)

func (r rdb) BeginTx(ctx context.Context, opts *sql.TxOptions) (TxDB, error) {
	return r.DB.BeginTx(ctx, opts)
}

type RDB interface {
	sqlcgen.DBTX
	BeginTx(context.Context, *sql.TxOptions) (TxDB, error)
}

type TxDB interface {
	sqlcgen.DBTX
	Rollback() error
	Commit() error
}

func Tx(ctx context.Context, db RDB, txFunc func(context.Context, TxDB) error) (err error) {
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
