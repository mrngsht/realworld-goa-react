// Code generated by goa v3.19.1, DO NOT EDIT.
//
// user HTTP server
//
// Command:
// $ goa gen github.com/mrngsht/realworld-goa-react/design

package server

import (
	"context"
	"net/http"

	user "github.com/mrngsht/realworld-goa-react/gen/user"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the user service endpoint HTTP handlers.
type Server struct {
	Mounts         []*MountPoint
	Login          http.Handler
	Register       http.Handler
	GetCurrentUser http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the user service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *user.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Login", "POST", "/api/users/login"},
			{"Register", "POST", "/api/users"},
			{"GetCurrentUser", "GET", "/api/user"},
		},
		Login:          NewLoginHandler(e.Login, mux, decoder, encoder, errhandler, formatter),
		Register:       NewRegisterHandler(e.Register, mux, decoder, encoder, errhandler, formatter),
		GetCurrentUser: NewGetCurrentUserHandler(e.GetCurrentUser, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "user" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Login = m(s.Login)
	s.Register = m(s.Register)
	s.GetCurrentUser = m(s.GetCurrentUser)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return user.MethodNames[:] }

// Mount configures the mux to serve the user endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountLoginHandler(mux, h.Login)
	MountRegisterHandler(mux, h.Register)
	MountGetCurrentUserHandler(mux, h.GetCurrentUser)
}

// Mount configures the mux to serve the user endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountLoginHandler configures the mux to serve the "user" service "login"
// endpoint.
func MountLoginHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/api/users/login", f)
}

// NewLoginHandler creates a HTTP handler which loads the HTTP request and
// calls the "user" service "login" endpoint.
func NewLoginHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeLoginRequest(mux, decoder)
		encodeResponse = EncodeLoginResponse(encoder)
		encodeError    = EncodeLoginError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "login")
		ctx = context.WithValue(ctx, goa.ServiceKey, "user")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountRegisterHandler configures the mux to serve the "user" service
// "register" endpoint.
func MountRegisterHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/api/users", f)
}

// NewRegisterHandler creates a HTTP handler which loads the HTTP request and
// calls the "user" service "register" endpoint.
func NewRegisterHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRegisterRequest(mux, decoder)
		encodeResponse = EncodeRegisterResponse(encoder)
		encodeError    = EncodeRegisterError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "register")
		ctx = context.WithValue(ctx, goa.ServiceKey, "user")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetCurrentUserHandler configures the mux to serve the "user" service
// "getCurrentUser" endpoint.
func MountGetCurrentUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/user", f)
}

// NewGetCurrentUserHandler creates a HTTP handler which loads the HTTP request
// and calls the "user" service "getCurrentUser" endpoint.
func NewGetCurrentUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeGetCurrentUserResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "getCurrentUser")
		ctx = context.WithValue(ctx, goa.ServiceKey, "user")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
