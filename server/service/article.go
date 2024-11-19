package service

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/guregu/null"
	goa "github.com/mrngsht/realworld-goa-react/gen/article"
	"github.com/mrngsht/realworld-goa-react/myctx"
	"github.com/mrngsht/realworld-goa-react/myrdb"
	"github.com/mrngsht/realworld-goa-react/myrdb/sqlcgen"
	"github.com/mrngsht/realworld-goa-react/mytime"
)

type Article struct {
	db myrdb.DB
}

func NewArticle(rdb myrdb.DB) *Article {
	return &Article{db: rdb}
}

var _ goa.Service = &Article{}

func (s *Article) Create(ctx context.Context, payload *goa.CreatePayload) (res *goa.CreateResult, err error) {
	userID := myctx.MustGetRequestUserID(ctx)
	db := s.db

	var articleID = uuid.Nil
	if err := myrdb.Tx(ctx, db, func(ctx context.Context, txdb myrdb.TxDB) error {
		db := txdb

		now := mytime.Now(ctx)
		newArticleID := uuid.New()

		if err := sqlcgen.Q.InsertArticle(ctx, db, sqlcgen.InsertArticleParams{
			CreatedAt: now,
			ID:        newArticleID,
		}); err != nil {
			return errors.WithStack(err)
		}

		if err := sqlcgen.Q.InsertArticleContent(ctx, db, sqlcgen.InsertArticleContentParams{
			CreatedAt:    now,
			ArticleID:    newArticleID,
			Title:        payload.Title,
			Description:  payload.Title,
			Body:         payload.Body,
			AuthorUserID: userID,
		}); err != nil {
			return errors.WithStack(err)
		}
		if err := sqlcgen.Q.InsertArticleContentMutation(ctx, db, sqlcgen.InsertArticleContentMutationParams{
			CreatedAt:    now,
			ArticleID:    newArticleID,
			Title:        payload.Title,
			Description:  payload.Title,
			Body:         payload.Body,
			AuthorUserID: userID,
		}); err != nil {
			return errors.WithStack(err)
		}

		if err := sqlcgen.Q.InsertArticleStats(ctx, db, sqlcgen.InsertArticleStatsParams{
			CreatedAt:      now,
			ArticleID:      newArticleID,
			FavoritesCount: null.IntFrom(0).Ptr(),
		}); err != nil {
			return errors.WithStack(err)
		}

		articleID = newArticleID

		return nil
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return &goa.CreateResult{
		Article: &goa.ArticleDetail{
			ID: articleID.String(),
		},
	}, nil
}
