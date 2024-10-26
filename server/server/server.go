package server

import (
	"net/http"

	goahttp "goa.design/goa/v3/http"

	"github.com/cockroachdb/errors"
	"github.com/mrngsht/realworld-goa-react/gen/http/user/server"
	"github.com/mrngsht/realworld-goa-react/gen/user"
	"github.com/mrngsht/realworld-goa-react/myrdb"
	"github.com/mrngsht/realworld-goa-react/service"
)

func Run() error {
	db, err := myrdb.OpenLocalRDB()
	if err != nil {
		return errors.WithStack(err)
	}
	defer db.Close()

	s := service.NewUser(db)
	endpoints := user.NewEndpoints(s)
	mux := goahttp.NewMuxer()
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	svr := server.New(endpoints, mux, dec, enc, nil, nil)

	svr.Use(newUserAuthorizationMiddleware())

	server.Mount(mux, svr)
	httpsvr := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	if err := httpsvr.ListenAndServe(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
