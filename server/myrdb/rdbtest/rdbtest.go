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
	"github.com/mrngsht/realworld-goa-react/myrdb/internal"
	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest/sqlctest"
)

func CreateRDB(t *testing.T, ctx context.Context) (*sql.DB, *sqlctest.Queries) {
	t.Helper()

	db, err := myrdb.OpenLocalRDB()
	if err != nil {
		panic(err)
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	internal.Tx = func(ctx context.Context, _ *sql.DB, txFunc func(ctx context.Context, tx *sql.Tx) error) (err error) {
		savePointName := uuid.New().String()

		_, err = tx.ExecContext(ctx, fmt.Sprintf(`SAVEPOINT "%s"`, savePointName))
		if err != nil {
			return errors.WithStack(err)
		}

		defer func() {
			if panicErr := recover(); panicErr != nil {
				err = errors.Errorf("panic error %v", panicErr)
			}
			if err != nil {
				if _, rollbackErr := tx.ExecContext(ctx, fmt.Sprintf(`ROLLBACK TO SAVEPOINT "%s"`, savePointName)); err != nil {
					err = errors.Join(err, rollbackErr)
				}
			}
		}()

		if err := txFunc(ctx, tx); err != nil {
			return errors.WithStack(err)
		}

		_, err = tx.ExecContext(ctx, fmt.Sprintf(`RELEASE SAVEPOINT "%s"`, savePointName))
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	}

	t.Cleanup(func() {
		if err := tx.Rollback(); err != nil {
			log.Default().Printf("test tx rollback error: %v", err)
		}
		if err := db.Close(); err != nil {
			log.Default().Printf("test db close error: %v", err)
		}
	})

	return db, sqlctest.New(db).WithTx(tx)
}
