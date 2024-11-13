// Code generated by goa v3.19.1, DO NOT EDIT.
//
// profile HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/mrngsht/realworld-goa-react/design

package server

import (
	"context"
	"errors"
	"io"
	"net/http"

	profile "github.com/mrngsht/realworld-goa-react/gen/profile"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeFollowUserResponse returns an encoder for responses returned by the
// profile followUser endpoint.
func EncodeFollowUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*profile.FollowUserResult)
		enc := encoder(ctx, w)
		body := NewFollowUserResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeFollowUserRequest returns a decoder for requests sent to the profile
// followUser endpoint.
func DecodeFollowUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body FollowUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateFollowUserRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewFollowUserPayload(&body)

		return payload, nil
	}
}

// EncodeFollowUserError returns an encoder for errors returned by the
// followUser profile endpoint.
func EncodeFollowUserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "UserAlreadyFollowing":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewFollowUserUserAlreadyFollowingResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUnfollowUserResponse returns an encoder for responses returned by the
// profile unfollowUser endpoint.
func EncodeUnfollowUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*profile.UnfollowUserResult)
		enc := encoder(ctx, w)
		body := NewUnfollowUserResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUnfollowUserRequest returns a decoder for requests sent to the profile
// unfollowUser endpoint.
func DecodeUnfollowUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body UnfollowUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUnfollowUserRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewUnfollowUserPayload(&body)

		return payload, nil
	}
}

// EncodeUnfollowUserError returns an encoder for errors returned by the
// unfollowUser profile endpoint.
func EncodeUnfollowUserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "UserNotFollowing":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUnfollowUserUserNotFollowingResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalProfileProfileToProfileResponseBody builds a value of type
// *ProfileResponseBody from a value of type *profile.Profile.
func marshalProfileProfileToProfileResponseBody(v *profile.Profile) *ProfileResponseBody {
	res := &ProfileResponseBody{
		Username:  v.Username,
		Bio:       v.Bio,
		Image:     v.Image,
		Following: v.Following,
	}

	return res
}
