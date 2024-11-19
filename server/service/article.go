package service

import (
	"context"
	"encoding/json"

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

	profile, err := sqlcgen.Q.GetUserProfileByUserID(ctx, db, userID)
	if err != nil {
		// handle ErrNoRows as internal server error
		return nil, errors.WithStack(err)
	}

	now := mytime.Now(ctx)
	articleID := uuid.New()
	if err := myrdb.Tx(ctx, db, func(ctx context.Context, txdb myrdb.TxDB) error {
		db := txdb

		if err := sqlcgen.Q.InsertArticle(ctx, db, sqlcgen.InsertArticleParams{
			CreatedAt: now,
			ID:        articleID,
		}); err != nil {
			return errors.WithStack(err)
		}

		if err := sqlcgen.Q.InsertArticleContent(ctx, db, sqlcgen.InsertArticleContentParams{
			CreatedAt:    now,
			ArticleID:    articleID,
			Title:        payload.Title,
			Description:  payload.Description,
			Body:         payload.Body,
			AuthorUserID: userID,
		}); err != nil {
			return errors.WithStack(err)
		}
		if err := sqlcgen.Q.InsertArticleContentMutation(ctx, db, sqlcgen.InsertArticleContentMutationParams{
			CreatedAt:    now,
			ArticleID:    articleID,
			Title:        payload.Title,
			Description:  payload.Description,
			Body:         payload.Body,
			AuthorUserID: userID,
		}); err != nil {
			return errors.WithStack(err)
		}

		if len(payload.TagList) > 0 {
			tagParams := make([]sqlcgen.InsertArticleTagParams, 0, len(payload.TagList))
			for i, tag := range payload.TagList {
				tagParams = append(tagParams, sqlcgen.InsertArticleTagParams{
					CreatedAt: now,
					ArticleID: articleID,
					SeqNo:     int32(i + 1),
					Tag:       tag,
				})
			}

			if _, err := sqlcgen.Q.InsertArticleTag(ctx, db, tagParams); err != nil {
				return errors.WithStack(err)
			}

			tagsJson, err := json.Marshal(payload.TagList)
			if err != nil {
				return errors.WithStack(err)
			}

			if err := sqlcgen.Q.InsertArticleTagMutation(ctx, db, sqlcgen.InsertArticleTagMutationParams{
				CreatedAt: now,
				ArticleID: articleID,
				Tags:      tagsJson,
			}); err != nil {
				return errors.WithStack(err)
			}
		}

		if err := sqlcgen.Q.InsertArticleStats(ctx, db, sqlcgen.InsertArticleStatsParams{
			CreatedAt:      now,
			ArticleID:      articleID,
			FavoritesCount: null.IntFrom(0).Ptr(),
		}); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return &goa.CreateResult{
		Article: &goa.ArticleDetail{
			ID:             articleID.String(),
			Title:          payload.Title,
			Description:    payload.Description,
			Body:           payload.Body,
			TagList:        payload.TagList,
			CreatedAt:      now.String(),
			UpdatedAt:      now.String(),
			Favorited:      false,
			FavoritesCount: 0,
			Author: &goa.Profile{
				Username:  profile.Username,
				Bio:       profile.Bio,
				Image:     profile.Bio,
				Following: false,
			},
		},
	}, nil
}
