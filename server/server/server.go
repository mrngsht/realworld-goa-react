package server

import (
	"context"
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/mrngsht/realworld-goa-react/myrdb"
)

func Run() error {
	ctx := context.Background()

	rdb, err := myrdb.OpenLocalRDB(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	defer rdb.Close(ctx)

	endpoints := setupEndpoints(rdb)
	mux := setupHttpServers(endpoints)

	httpsvr := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	if err := httpsvr.ListenAndServe(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
