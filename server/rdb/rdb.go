package rdb

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	LocalConnectionString = "host=localhost user=postgres password=postgres dbname=realworld sslmode=disable"
)

func Tx(ctx context.Context, db *sql.DB, txFunc func(ctx context.Context, tx *sql.Tx) error) (err error) {
	tx, beginErr := db.BeginTx(ctx, nil)
	if beginErr != nil {
		return beginErr
	}

	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = fmt.Errorf("panic error %v", panicErr)
		}
		if err != nil {
			tx.Rollback()
		}
	}()

	if err := txFunc(ctx, tx); err != nil {
		return err
	}

	return tx.Commit()
}
