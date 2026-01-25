package service

import (
	"context"

	"github.com/cockroachdb/errors"
	goa "github.com/mrngsht/realworld-goa-react/gen/tag"
	"github.com/mrngsht/realworld-goa-react/myrdb"
	"github.com/mrngsht/realworld-goa-react/myrdb/sqlcgen"
)

type Tag struct {
	db myrdb.DB
}

func NewTag(db myrdb.DB) *Tag {
	return &Tag{db: db}
}

func (s *Tag) GetTags(ctx context.Context) (res *goa.GetTagsResult, err error) {
	db := s.db

	tags, err := sqlcgen.Q.ListAllTags(ctx, db)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &goa.GetTagsResult{Tags: tags}, nil
}
