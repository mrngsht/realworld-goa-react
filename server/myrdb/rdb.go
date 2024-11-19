package myrdb

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cockroachdb/errors"
	"github.com/mrngsht/realworld-goa-react/config"
	"github.com/mrngsht/realworld-goa-react/myrdb/sqlcgen"

	_ "github.com/lib/pq"
)

func OpenLocalRDB(ctx context.Context) (conn, error) {
	cfg, err := pgxpool.ParseConfig(config.C.RDBConnectionString)
	if err != nil {
		return conn{}, errors.WithStack(err)
	}
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return conn{}, errors.WithStack(err)
	}
	return conn{pool}, nil
}

type conn struct {
	*pgxpool.Pool // use pool as connection because we don't need the queries to run on the same connection except the transaction.
}

var _ Conn = (*conn)(nil)

func (r conn) BeginTx(ctx context.Context, opts pgx.TxOptions) (TxConn, error) {
	return r.Pool.BeginTx(ctx, opts)
}

type Conn interface {
	sqlcgen.DBTX
	BeginTx(context.Context, pgx.TxOptions) (TxConn, error)
}

type TxConn interface {
	sqlcgen.DBTX
	Rollback(context.Context) error
	Commit(context.Context) error
}

func Tx(ctx context.Context, db Conn, txFunc func(context.Context, TxConn) error) (err error) {
	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return errors.WithStack(err)
	}

	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Errorf("panic error %v", panicErr)
		}

		if err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				err = errors.Join(err, rollbackErr)
			}
		} else {
			if err := tx.Commit(ctx); err != nil {
				err = errors.WithStack(err)
			}
		}
	}()

	if err := txFunc(ctx, tx); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
