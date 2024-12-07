// Code generated by goa v3.19.1, DO NOT EDIT.
//
// article HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/mrngsht/realworld-goa-react/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	article "github.com/mrngsht/realworld-goa-react/gen/article"
	goahttp "goa.design/goa/v3/http"
)

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "article" service "get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		articleID string
	)
	{
		p, ok := v.(*article.GetPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("article", "get", "*article.GetPayload", v)
		}
		articleID = p.ArticleID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetArticlePath(articleID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("article", "get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetResponse returns a decoder for responses returned by the article
// get endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetResponse may return the following errors:
//   - "ArticleGetArticleBadRequest" (type *article.ArticleGetArticleBadRequest): http.StatusBadRequest
//   - error: internal error
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("article", "get", err)
			}
			err = ValidateGetResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("article", "get", err)
			}
			res := NewGetResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body GetArticleGetArticleBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("article", "get", err)
			}
			err = ValidateGetArticleGetArticleBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("article", "get", err)
			}
			return nil, NewGetArticleGetArticleBadRequest(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("article", "get", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "article" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateArticlePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("article", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the article
// create server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*article.CreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("article", "create", "*article.CreatePayload", v)
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("article", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the article
// create endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("article", "create", err)
			}
			err = ValidateCreateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("article", "create", err)
			}
			res := NewCreateResultOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("article", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildFavoriteRequest instantiates a HTTP request object with method and path
// set to call the "article" service "favorite" endpoint
func (c *Client) BuildFavoriteRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		articleID string
	)
	{
		p, ok := v.(*article.FavoritePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("article", "favorite", "*article.FavoritePayload", v)
		}
		articleID = p.ArticleID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: FavoriteArticlePath(articleID)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("article", "favorite", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeFavoriteResponse returns a decoder for responses returned by the
// article favorite endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeFavoriteResponse may return the following errors:
//   - "ArticleFavoriteArticleBadRequest" (type *article.ArticleFavoriteArticleBadRequest): http.StatusBadRequest
//   - error: internal error
func DecodeFavoriteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body FavoriteResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("article", "favorite", err)
			}
			err = ValidateFavoriteResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("article", "favorite", err)
			}
			res := NewFavoriteResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body FavoriteArticleFavoriteArticleBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("article", "favorite", err)
			}
			err = ValidateFavoriteArticleFavoriteArticleBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("article", "favorite", err)
			}
			return nil, NewFavoriteArticleFavoriteArticleBadRequest(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("article", "favorite", resp.StatusCode, string(body))
		}
	}
}

// unmarshalArticleDetailResponseBodyToArticleArticleDetail builds a value of
// type *article.ArticleDetail from a value of type *ArticleDetailResponseBody.
func unmarshalArticleDetailResponseBodyToArticleArticleDetail(v *ArticleDetailResponseBody) *article.ArticleDetail {
	res := &article.ArticleDetail{
		ArticleID:      *v.ArticleID,
		Title:          *v.Title,
		Description:    *v.Description,
		Body:           *v.Body,
		CreatedAt:      *v.CreatedAt,
		UpdatedAt:      *v.UpdatedAt,
		Favorited:      *v.Favorited,
		FavoritesCount: *v.FavoritesCount,
	}
	res.TagList = make([]string, len(v.TagList))
	for i, val := range v.TagList {
		res.TagList[i] = val
	}
	res.Author = unmarshalProfileResponseBodyToArticleProfile(v.Author)

	return res
}

// unmarshalProfileResponseBodyToArticleProfile builds a value of type
// *article.Profile from a value of type *ProfileResponseBody.
func unmarshalProfileResponseBodyToArticleProfile(v *ProfileResponseBody) *article.Profile {
	res := &article.Profile{
		Username:  *v.Username,
		Bio:       *v.Bio,
		Image:     *v.Image,
		Following: *v.Following,
	}

	return res
}
