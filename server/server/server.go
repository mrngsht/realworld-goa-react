package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/mrngsht/realworld-goa-react/mylog"
	"github.com/mrngsht/realworld-goa-react/myrdb"
)

func Run() error {
	ctx := context.Background()

	db, err := myrdb.OpenLocalDB(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	defer db.Close()

	endpoints := setupEndpoints(db)
	mux := setupHttpServers(endpoints)

	httpsvr := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	mylog.Info(ctx, fmt.Sprintf("running server on %s ...", httpsvr.Addr))

	if err := httpsvr.ListenAndServe(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
