package article

import (
	"github.com/mrngsht/realworld-goa-react/myerr"
)

var (
	ErrArticleNotFound        = myerr.NewAppErr("article not found")
	ErrArticleCommentNotFound = myerr.NewAppErr("article comment not found")
	ErrRequestUserIsNotAuthor = myerr.NewAppErr("request user is not author")
)
