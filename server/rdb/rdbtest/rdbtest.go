package rdbtest

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/mrngsht/realworld-goa-react/rdb"
	"github.com/mrngsht/realworld-goa-react/rdb/internal"
	"github.com/mrngsht/realworld-goa-react/rdb/rdbtest/sqlctest"
)

func CreateRDB(t *testing.T, ctx context.Context) (*sql.DB, *sqlctest.Queries) {
	t.Helper()

	db, err := rdb.OpenLocalRDB()
	if err != nil {
		panic(err)
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	internal.Tx = func(ctx context.Context, _ *sql.DB, txFunc func(ctx context.Context, tx *sql.Tx) error) error {
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
		tx.Rollback()
		db.Close()
	})

	return db, sqlctest.New(db).WithTx(tx)
}
