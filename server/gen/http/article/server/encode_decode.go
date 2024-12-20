// Code generated by goa v3.19.1, DO NOT EDIT.
//
// article HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/mrngsht/realworld-goa-react/design

package server

import (
	"context"
	"errors"
	"io"
	"net/http"

	article "github.com/mrngsht/realworld-goa-react/gen/article"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetResponse returns an encoder for responses returned by the article
// get endpoint.
func EncodeGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*article.GetResult)
		enc := encoder(ctx, w)
		body := NewGetResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetRequest returns a decoder for requests sent to the article get
// endpoint.
func DecodeGetRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			articleID string
			err       error

			params = mux.Vars(r)
		)
		articleID = params["articleId"]
		err = goa.MergeErrors(err, goa.ValidateFormat("articleId", articleID, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
		payload := NewGetPayload(articleID)

		return payload, nil
	}
}

// EncodeGetError returns an encoder for errors returned by the get article
// endpoint.
func EncodeGetError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "ArticleGetArticleBadRequest":
			var res *article.ArticleGetArticleBadRequest
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetArticleGetArticleBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateResponse returns an encoder for responses returned by the
// article create endpoint.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*article.CreateResult)
		enc := encoder(ctx, w)
		body := NewCreateResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeCreateRequest returns a decoder for requests sent to the article
// create endpoint.
func DecodeCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body CreateRequestBody
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
		err = ValidateCreateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreatePayload(&body)

		return payload, nil
	}
}

// EncodeFavoriteResponse returns an encoder for responses returned by the
// article favorite endpoint.
func EncodeFavoriteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*article.FavoriteResult)
		enc := encoder(ctx, w)
		body := NewFavoriteResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeFavoriteRequest returns a decoder for requests sent to the article
// favorite endpoint.
func DecodeFavoriteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			articleID string
			err       error

			params = mux.Vars(r)
		)
		articleID = params["articleId"]
		err = goa.MergeErrors(err, goa.ValidateFormat("articleId", articleID, goa.FormatUUID))
		if err != nil {
			return nil, err
		}
		payload := NewFavoritePayload(articleID)

		return payload, nil
	}
}

// marshalArticleArticleDetailToArticleDetailResponseBody builds a value of
// type *ArticleDetailResponseBody from a value of type *article.ArticleDetail.
func marshalArticleArticleDetailToArticleDetailResponseBody(v *article.ArticleDetail) *ArticleDetailResponseBody {
	res := &ArticleDetailResponseBody{
		ArticleID:      v.ArticleID,
		Title:          v.Title,
		Description:    v.Description,
		Body:           v.Body,
		CreatedAt:      v.CreatedAt,
		UpdatedAt:      v.UpdatedAt,
		Favorited:      v.Favorited,
		FavoritesCount: v.FavoritesCount,
	}
	if v.TagList != nil {
		res.TagList = make([]string, len(v.TagList))
		for i, val := range v.TagList {
			res.TagList[i] = val
		}
	} else {
		res.TagList = []string{}
	}
	if v.Author != nil {
		res.Author = marshalArticleProfileToProfileResponseBody(v.Author)
	}

	return res
}

// marshalArticleProfileToProfileResponseBody builds a value of type
// *ProfileResponseBody from a value of type *article.Profile.
func marshalArticleProfileToProfileResponseBody(v *article.Profile) *ProfileResponseBody {
	res := &ProfileResponseBody{
		Username:  v.Username,
		Bio:       v.Bio,
		Image:     v.Image,
		Following: v.Following,
	}

	return res
}
