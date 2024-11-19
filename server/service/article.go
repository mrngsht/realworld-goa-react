package service

import (
	"context"

	"github.com/cockroachdb/errors"

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

func (s *Article) Create(ctx context.Context, payload *goa.CreatePayload) (res *goa.CreateResult, err error) {
	if err := myrdb.Tx(ctx, s.rdb, func(ctx context.Context, tx myrdb.TxDB) error {

		return nil
	}); err != nil {
		return nil, errors.WithStack(err)
	}
	return nil, nil
}
