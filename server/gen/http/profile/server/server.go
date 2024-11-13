// Code generated by goa v3.19.1, DO NOT EDIT.
//
// profile HTTP server
//
// Command:
// $ goa gen github.com/mrngsht/realworld-goa-react/design

package server

import (
	"context"
	"net/http"

	profile "github.com/mrngsht/realworld-goa-react/gen/profile"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the profile service endpoint HTTP handlers.
type Server struct {
	Mounts       []*MountPoint
	FollowUser   http.Handler
	UnfollowUser http.Handler
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

// New instantiates HTTP handlers for all the profile service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *profile.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"FollowUser", "POST", "/api/profile/follow_user"},
			{"UnfollowUser", "POST", "/api/profile/unfollow_user"},
		},
		FollowUser:   NewFollowUserHandler(e.FollowUser, mux, decoder, encoder, errhandler, formatter),
		UnfollowUser: NewUnfollowUserHandler(e.UnfollowUser, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "profile" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.FollowUser = m(s.FollowUser)
	s.UnfollowUser = m(s.UnfollowUser)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return profile.MethodNames[:] }

// Mount configures the mux to serve the profile endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountFollowUserHandler(mux, h.FollowUser)
	MountUnfollowUserHandler(mux, h.UnfollowUser)
}

// Mount configures the mux to serve the profile endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountFollowUserHandler configures the mux to serve the "profile" service
// "followUser" endpoint.
func MountFollowUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/api/profile/follow_user", f)
}

// NewFollowUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "profile" service "followUser" endpoint.
func NewFollowUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeFollowUserRequest(mux, decoder)
		encodeResponse = EncodeFollowUserResponse(encoder)
		encodeError    = EncodeFollowUserError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "followUser")
		ctx = context.WithValue(ctx, goa.ServiceKey, "profile")
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

// MountUnfollowUserHandler configures the mux to serve the "profile" service
// "unfollowUser" endpoint.
func MountUnfollowUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/api/profile/unfollow_user", f)
}

// NewUnfollowUserHandler creates a HTTP handler which loads the HTTP request
// and calls the "profile" service "unfollowUser" endpoint.
func NewUnfollowUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUnfollowUserRequest(mux, decoder)
		encodeResponse = EncodeUnfollowUserResponse(encoder)
		encodeError    = EncodeUnfollowUserError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "unfollowUser")
		ctx = context.WithValue(ctx, goa.ServiceKey, "profile")
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
