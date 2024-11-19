package server

import (
	"net/http"

	article "github.com/mrngsht/realworld-goa-react/gen/http/article/server"
	profile "github.com/mrngsht/realworld-goa-react/gen/http/profile/server"
	user "github.com/mrngsht/realworld-goa-react/gen/http/user/server"

	goahttp "goa.design/goa/v3/http"
)

func setupHttpServers(endpoints *endpoints) goahttp.ResolverMuxer {
	mux, dec, enc := goahttp.NewMuxer(), goahttp.RequestDecoder, goahttp.ResponseEncoder

	setupHttpServer(user.New(endpoints.User, mux, dec, enc, nil, nil), mux)
	setupHttpServer(profile.New(endpoints.Profile, mux, dec, enc, nil, nil), mux)
	setupHttpServer(article.New(endpoints.Article, mux, dec, enc, nil, nil), mux)

	return mux
}

type httpServer interface {
	Use(func(http.Handler) http.Handler)
	Mount(mux goahttp.Muxer)
}

func setupHttpServer(s httpServer, mux goahttp.Muxer) {
	s.Use(newPanicRecoverMiddleware())
	s.Use(newUserAuthorizationMiddleware())
	s.Use(newRequestLogMiddleware())
	s.Use(newRequestIDMiddleware())
	s.Mount(mux)
}
