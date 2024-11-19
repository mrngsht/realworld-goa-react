package server

import (
	"github.com/mrngsht/realworld-goa-react/gen/profile"
	"github.com/mrngsht/realworld-goa-react/gen/user"
	"github.com/mrngsht/realworld-goa-react/myrdb"
	"github.com/mrngsht/realworld-goa-react/service"
	goa "goa.design/goa/v3/pkg"
)

type endpoints struct {
	User    *user.Endpoints
	Profile *profile.Endpoints
}

func setupEndpoints(rdb myrdb.DB) *endpoints {
	return &endpoints{
		User:    setupEndpoint(user.NewEndpoints(service.NewUser(rdb))),
		Profile: setupEndpoint(profile.NewEndpoints(service.NewProfile(rdb))),
	}
}

type endpoint interface {
	Use(func(goa.Endpoint) goa.Endpoint)
}

func setupEndpoint[T endpoint](u T) T {
	u.Use(newErrorHandlerMiddleware())
	return u
}
