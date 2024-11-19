// Code generated by goa v3.19.1, DO NOT EDIT.
//
// article HTTP client types
//
// Command:
// $ goa gen github.com/mrngsht/realworld-goa-react/design

package client

import (
	article "github.com/mrngsht/realworld-goa-react/gen/article"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "article" service "create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	Title       string   `form:"title" json:"title" xml:"title"`
	Description string   `form:"description" json:"description" xml:"description"`
	Body        string   `form:"body" json:"body" xml:"body"`
	TagList     []string `form:"tagList" json:"tagList" xml:"tagList"`
}

// CreateResponseBody is the type of the "article" service "create" endpoint
// HTTP response body.
type CreateResponseBody struct {
	Article *ArticleDetailResponseBody `form:"article,omitempty" json:"article,omitempty" xml:"article,omitempty"`
}

// ArticleDetailResponseBody is used to define fields on response body types.
type ArticleDetailResponseBody struct {
	ID             *string              `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Title          *string              `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	Description    *string              `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Body           *string              `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
	TagList        []string             `form:"tagList,omitempty" json:"tagList,omitempty" xml:"tagList,omitempty"`
	CreatedAt      *string              `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	UpdatedAt      *string              `form:"updatedAt,omitempty" json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	Favorited      *bool                `form:"favorited,omitempty" json:"favorited,omitempty" xml:"favorited,omitempty"`
	FavoritesCount *uint                `form:"favoritesCount,omitempty" json:"favoritesCount,omitempty" xml:"favoritesCount,omitempty"`
	Author         *ProfileResponseBody `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
}

// ProfileResponseBody is used to define fields on response body types.
type ProfileResponseBody struct {
	Username  *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
	Bio       *string `form:"bio,omitempty" json:"bio,omitempty" xml:"bio,omitempty"`
	Image     *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	Following *bool   `form:"following,omitempty" json:"following,omitempty" xml:"following,omitempty"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "article" service.
func NewCreateRequestBody(p *article.CreatePayload) *CreateRequestBody {
	body := &CreateRequestBody{
		Title:       p.Title,
		Description: p.Description,
		Body:        p.Body,
	}
	if p.TagList != nil {
		body.TagList = make([]string, len(p.TagList))
		for i, val := range p.TagList {
			body.TagList[i] = val
		}
	} else {
		body.TagList = []string{}
	}
	return body
}

// NewCreateResultOK builds a "article" service "create" endpoint result from a
// HTTP "OK" response.
func NewCreateResultOK(body *CreateResponseBody) *article.CreateResult {
	v := &article.CreateResult{}
	v.Article = unmarshalArticleDetailResponseBodyToArticleArticleDetail(body.Article)

	return v
}

// ValidateCreateResponseBody runs the validations defined on CreateResponseBody
func ValidateCreateResponseBody(body *CreateResponseBody) (err error) {
	if body.Article == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("article", "body"))
	}
	if body.Article != nil {
		if err2 := ValidateArticleDetailResponseBody(body.Article); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateArticleDetailResponseBody runs the validations defined on
// ArticleDetailResponseBody
func ValidateArticleDetailResponseBody(body *ArticleDetailResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Body == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("body", "body"))
	}
	if body.TagList == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("tagList", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("createdAt", "body"))
	}
	if body.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updatedAt", "body"))
	}
	if body.Favorited == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("favorited", "body"))
	}
	if body.FavoritesCount == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("favoritesCount", "body"))
	}
	if body.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	if body.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.createdAt", *body.CreatedAt, goa.FormatDateTime))
	}
	if body.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.updatedAt", *body.UpdatedAt, goa.FormatDateTime))
	}
	if body.Author != nil {
		if err2 := ValidateProfileResponseBody(body.Author); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateProfileResponseBody runs the validations defined on
// ProfileResponseBody
func ValidateProfileResponseBody(body *ProfileResponseBody) (err error) {
	if body.Username == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("username", "body"))
	}
	if body.Bio == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("bio", "body"))
	}
	if body.Image == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("image", "body"))
	}
	if body.Following == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("following", "body"))
	}
	return
}