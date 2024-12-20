// Code generated by goa v3.19.1, DO NOT EDIT.
//
// article service
//
// Command:
// $ goa gen github.com/mrngsht/realworld-goa-react/design

package article

import (
	"context"
)

// article
type Service interface {
	// Get implements get.
	Get(context.Context, *GetPayload) (res *GetResult, err error)
	// Create implements create.
	Create(context.Context, *CreatePayload) (res *CreateResult, err error)
	// Favorite implements favorite.
	Favorite(context.Context, *FavoritePayload) (res *FavoriteResult, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "readlworld"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "article"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"get", "create", "favorite"}

type ArticleDetail struct {
	ArticleID      string
	Title          string
	Description    string
	Body           string
	TagList        []string
	CreatedAt      string
	UpdatedAt      string
	Favorited      bool
	FavoritesCount uint
	Author         *Profile
}

type ArticleGetArticleBadRequest struct {
	Code string
}

// CreatePayload is the payload type of the article service create method.
type CreatePayload struct {
	Title       string
	Description string
	Body        string
	TagList     []string
}

// CreateResult is the result type of the article service create method.
type CreateResult struct {
	Article *ArticleDetail
}

// FavoritePayload is the payload type of the article service favorite method.
type FavoritePayload struct {
	ArticleID string
}

// FavoriteResult is the result type of the article service favorite method.
type FavoriteResult struct {
	Article *ArticleDetail
}

// GetPayload is the payload type of the article service get method.
type GetPayload struct {
	ArticleID string
}

// GetResult is the result type of the article service get method.
type GetResult struct {
	Article *ArticleDetail
}

type Profile struct {
	Username  string
	Bio       string
	Image     string
	Following bool
}

// Error returns an error description.
func (e *ArticleGetArticleBadRequest) Error() string {
	return ""
}

// ErrorName returns "ArticleGetArticleBadRequest".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *ArticleGetArticleBadRequest) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "ArticleGetArticleBadRequest".
func (e *ArticleGetArticleBadRequest) GoaErrorName() string {
	return "ArticleGetArticleBadRequest"
}
