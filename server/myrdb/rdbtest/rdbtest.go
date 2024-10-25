package rdbtest

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/mrngsht/realworld-goa-react/myrdb"
	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest/sqlctest"
	"github.com/mrngsht/realworld-goa-react/myrdb/sqlcgen"
)

func CreateRDB(t *testing.T, ctx context.Context) (*testdb, *sqlcgen.Queries, *sqlctest.Queries) {
	t.Helper()

	db, err := myrdb.OpenLocalRDB()
	if err != nil {
		panic(err)
	}

	tx, err := db.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		if err := tx.Rollback(); err != nil {
			log.Default().Printf("test tx rollback error: %v", err)
		}
		if err := db.Close(); err != nil {
			log.Default().Printf("test db close error: %v", err)
		}
	})

	return &testdb{Tx: tx, savePointName: uuid.New().String()}, sqlcgen.New(tx), sqlctest.New(tx)
}

type testdb struct {
	*sql.Tx

	savePointName string
}

var _ myrdb.RDB = (*testdb)(nil)
var _ myrdb.TxDB = (*testdb)(nil)

func (t *testdb) BeginTx(context.Context, *sql.TxOptions) (myrdb.TxDB, error) {
	_, err := t.Exec(fmt.Sprintf(`SAVEPOINT "%s"`, t.savePointName))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return t, nil
}

func (t *testdb) Commit() error {
	_, err := t.Exec(fmt.Sprintf(`RELEASE SAVEPOINT "%s"`, t.savePointName))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (t *testdb) Rollback() error {
	_, err := t.Exec(fmt.Sprintf(`ROLLBACK TO SAVEPOINT "%s"`, t.savePointName))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
