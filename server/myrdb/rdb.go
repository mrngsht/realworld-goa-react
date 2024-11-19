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

func OpenLocalDB(ctx context.Context) (db, error) {
	cfg, err := pgxpool.ParseConfig(config.C.RDBConnectionString)
	if err != nil {
		return db{}, errors.WithStack(err)
	}
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return db{}, errors.WithStack(err)
	}
	return db{pool}, nil
}

type db struct {
	*pgxpool.Pool // use pool as rdb because we don't need the queries to run on the same connection except the transaction.
}

var _ DB = (*db)(nil)

func (r db) BeginTx(ctx context.Context, opts pgx.TxOptions) (TxDB, error) {
	return r.Pool.BeginTx(ctx, opts)
}

type DB interface {
	sqlcgen.DBTX
	BeginTx(context.Context, pgx.TxOptions) (TxDB, error)
}

type TxDB interface {
	sqlcgen.DBTX
	Rollback(context.Context) error
	Commit(context.Context) error
}

func Tx(ctx context.Context, db DB, txFunc func(context.Context, TxDB) error) (err error) {
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
