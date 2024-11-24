package server

import (
	"context"
	"net/http"

	"github.com/mrngsht/realworld-goa-react/design"
	article "github.com/mrngsht/realworld-goa-react/gen/http/article/server"
	profile "github.com/mrngsht/realworld-goa-react/gen/http/profile/server"
	user "github.com/mrngsht/realworld-goa-react/gen/http/user/server"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

func setupHttpServers(endpoints *endpoints) goahttp.ResolverMuxer {
	mux, dec, enc, er := goahttp.NewMuxer(), goahttp.RequestDecoder, goahttp.ResponseEncoder, customErrorResponse

	setupHttpServer(user.New(endpoints.User, mux, dec, enc, nil, er), mux)
	setupHttpServer(profile.New(endpoints.Profile, mux, dec, enc, nil, er), mux)
	setupHttpServer(article.New(endpoints.Article, mux, dec, enc, nil, er), mux)

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

func customErrorResponse(ctx context.Context, err error) goahttp.Statuser {
	if serr, ok := err.(*goa.ServiceError); ok {
		switch serr.Name {
		case design.ErrorCommon_AuthenticationRequired:
			return errAuthenticationRequired
		default:
			// Use Goa default
			return goahttp.NewErrorResponse(ctx, err)
		}
	}
	// Use Goa default for all other error types
	return goahttp.NewErrorResponse(ctx, err)
}

var errAuthenticationRequired = errUnauthorized{Message: "authentication required"}

type errUnauthorized struct {
	Message string `json:"message"`
}

func (errUnauthorized) StatusCode() int {
	return http.StatusUnauthorized
}
