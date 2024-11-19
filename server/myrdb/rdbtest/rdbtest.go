package rdbtest

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/mrngsht/realworld-goa-react/myrdb"
	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest/sqlctest"
	"github.com/mrngsht/realworld-goa-react/myrdb/sqlcgen"
)

func CreateRDB(t *testing.T, ctx context.Context) (*testdb, *sqlcgen.Queries, *sqlctest.Queries) {
	t.Helper()

	db, err := myrdb.OpenLocalRDB(ctx)
	if err != nil {
		panic(err)
	}

	tx, err := db.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Default().Printf("test tx rollback error: %v", err)
		}
		db.Pool.Close()
	})

	return &testdb{Tx: tx, savePointName: uuid.New().String()}, sqlcgen.New(tx), sqlctest.New(tx)
}

type testdb struct {
	pgx.Tx

	savePointName string
}

var _ myrdb.Conn = (*testdb)(nil)
var _ myrdb.TxConn = (*testdb)(nil)

func (t *testdb) BeginTx(ctx context.Context, _ pgx.TxOptions) (myrdb.TxConn, error) {
	_, err := t.Exec(ctx, fmt.Sprintf(`SAVEPOINT "%s"`, t.savePointName))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return t, nil
}

func (t *testdb) Commit(ctx context.Context) error {
	_, err := t.Exec(ctx, fmt.Sprintf(`RELEASE SAVEPOINT "%s"`, t.savePointName))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (t *testdb) Rollback(ctx context.Context) error {
	_, err := t.Exec(ctx, fmt.Sprintf(`ROLLBACK TO SAVEPOINT "%s"`, t.savePointName))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
