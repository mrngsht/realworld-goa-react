package service

import (
	"context"

	goa "github.com/mrngsht/realworld-goa-react/gen/article"
	"github.com/mrngsht/realworld-goa-react/myrdb"
)

type Article struct {
	rdb myrdb.Conn
}

func NewArticle(rdb myrdb.Conn) *Article {
	return &Article{rdb: rdb}
}

var _ goa.Service = &Article{}

func (s *Article) Create(ctx context.Context, payload *goa.CreatePayload) (res *goa.CreateResult, err error) {
	// userID := myctx.MustGetRequestUserID(ctx)
	//
	// var articleID = uuid.Nil
	// if err := myrdb.Tx(ctx, s.rdb, func(ctx context.Context, tx myrdb.TxConn) error {
	// 	q := sqlcgen.New(tx)
	//
	// 	now := mytime.Now(ctx)
	// 	newArticleID := uuid.New()
	//
	// 	if err := q.InsertArticle(ctx, sqlcgen.InsertArticleParams{
	// 		CreatedAt: now,
	// 		ID:        newArticleID,
	// 	}); err != nil {
	// 		return errors.WithStack(err)
	// 	}
	//
	// 	if err := q.InsertArticleContent(ctx, sqlcgen.InsertArticleContentParams{
	// 		CreatedAt:    now,
	// 		ArticleID:    newArticleID,
	// 		Title:        payload.Title,
	// 		Description:  payload.Title,
	// 		Body:         payload.Body,
	// 		AuthorUserID: userID,
	// 	}); err != nil {
	// 		return errors.WithStack(err)
	// 	}
	// 	if err := q.InsertArticleContentMutation(ctx, sqlcgen.InsertArticleContentMutationParams{
	// 		CreatedAt:    now,
	// 		ArticleID:    newArticleID,
	// 		Title:        payload.Title,
	// 		Description:  payload.Title,
	// 		Body:         payload.Body,
	// 		AuthorUserID: userID,
	// 	}); err != nil {
	// 		return errors.WithStack(err)
	// 	}
	//
	// 	if err := q.InsertArticleStats(ctx, sqlcgen.InsertArticleStatsParams{
	// 		CreatedAt:      now,
	// 		ArticleID:      newArticleID,
	// 		FavoritesCount: sql.NullInt64{Int64: 0, Valid: true},
	// 	}); err != nil {
	// 		return errors.WithStack(err)
	// 	}
	//
	// 	articleID = newArticleID
	//
	// 	return nil
	// }); err != nil {
	// 	return nil, errors.WithStack(err)
	// }
	return nil, nil
}
