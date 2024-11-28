package article

import (
	"github.com/mrngsht/realworld-goa-react/myerr"
)

var (
	ErrArticleNotFound = myerr.NewAppErr("article not found")
)
