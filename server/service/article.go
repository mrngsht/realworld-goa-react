package service

import (
	"context"

	goa "github.com/mrngsht/realworld-goa-react/gen/article"
	"github.com/mrngsht/realworld-goa-react/myrdb"
)

type Article struct {
	rdb myrdb.RDB
}

func NewArticle(rdb myrdb.RDB) *Article {
	return &Article{rdb: rdb}
}

var _ goa.Service = &Article{}

func (s *Article) Create(context.Context, *goa.CreatePayload) (res *goa.CreateResult, err error) {
	return nil, nil
}
