package service

import (
	"context"
	"encoding/json"

	"github.com/cockroachdb/errors"
	"github.com/mrngsht/realworld-goa-react/design"
	"github.com/mrngsht/realworld-goa-react/domain/article"
	goa "github.com/mrngsht/realworld-goa-react/gen/article"
	"github.com/mrngsht/realworld-goa-react/myctx"
	"github.com/mrngsht/realworld-goa-react/myerr"
	"github.com/mrngsht/realworld-goa-react/myrdb"
	"github.com/mrngsht/realworld-goa-react/myrdb/sqlcgen"
	"github.com/mrngsht/realworld-goa-react/mytime"

	"github.com/google/uuid"
)

type Article struct {
	db myrdb.DB
}

func NewArticle(rdb myrdb.DB) *Article {
	return &Article{db: rdb}
}

var _ goa.Service = &Article{}

func (s *Article) Get(ctx context.Context, payload *goa.GetPayload) (res *goa.GetResult, err error) {
	defer func() {
		if apErr, ok := myerr.AsAppErr(err); ok {
			switch apErr {
			case article.ErrArticleNotFound:
				err = &goa.ArticleGetArticleBadRequest{Code: design.ErrCode_Article_ArticleNotFound}
			}
		}
	}()

	userIDOptional := myctx.MayGetAuthenticatedUserID(ctx)

	detail, err := s.getArticleDetail(ctx, uuid.MustParse(payload.ArticleID), userIDOptional, false)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &goa.GetResult{Article: detail}, nil
}

func (s *Article) List(ctx context.Context, payload *goa.ListPayload) (res *goa.ListResult, err error) {
	userIDOptional := myctx.MayGetAuthenticatedUserID(ctx)

	db := s.db

	articleIDs, err := sqlcgen.Q.ListArticleIDsBySearch(ctx, db, sqlcgen.ListArticleIDsBySearchParams{
		Limit:             payload.Limit,
		Offset:            payload.Offset,
		Tag:               payload.Tag,
		AutherUsername:    payload.Author,
		FavoritedUsername: payload.Favorited,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	contents, err := sqlcgen.Q.ListArticleContentsByArticleIDs(ctx, db, articleIDs)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	authorIDs := make([]uuid.UUID, 0, len(contents))
	{
		seen := make(map[uuid.UUID]bool)
		for _, c := range contents {
			if !seen[c.AuthorUserID] {
				authorIDs = append(authorIDs, c.AuthorUserID)
				seen[c.AuthorUserID] = true
			}
		}
	}

	authorProfileMap := make(map[uuid.UUID]sqlcgen.ListUserProfilesByUserIDsRow)
	{
		profiles, err := sqlcgen.Q.ListUserProfilesByUserIDs(ctx, db, authorIDs)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for _, p := range profiles {
			authorProfileMap[p.UserID] = p
		}
	}

	stats, err := sqlcgen.Q.ListArticleStatsByArticleIDs(ctx, db, articleIDs)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	articleTags, err := sqlcgen.Q.ListArticleTagsByArticleIDs(ctx, db, articleIDs)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if userIDOptional != nil {
		userID := *userIDOptional

		favoritedArticles, err := sqlcgen.Q.ListFavoritedArticlesByUserIDAndArticleIDs(ctx, db,
			sqlcgen.ListFavoritedArticlesByUserIDAndArticleIDsParams{
				UserID:     userID,
				ArticleIds: articleIDs,
			})
		if err != nil {
			return nil, errors.WithStack(err)
		}

		followedUserIDs, err := sqlcgen.Q.ListFollowedUserIDsByUserIDAndFollowedUserIDs(ctx, db, sqlcgen.ListFollowedUserIDsByUserIDAndFollowedUserIDsParams{
			UserID:          userID,
			FollowedUserIds: authorIDs,
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}

	}

	return nil, nil
}

func (s *Article) Create(ctx context.Context, payload *goa.CreatePayload) (res *goa.CreateResult, err error) {
	userID, err := myctx.ShouldGetAuthenticatedUserID(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

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
			FavoritesCount: int64(0),
		}); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return &goa.CreateResult{
		Article: &goa.ArticleDetail{
			ArticleID:      articleID.String(),
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
				Image:     profile.ImageUrl,
				Following: false,
			},
		},
	}, nil
}

func (s *Article) Update(ctx context.Context, payload *goa.UpdatePayload) (res *goa.UpdateResult, err error) {
	defer func() {
		if apErr, ok := myerr.AsAppErr(err); ok {
			switch apErr {
			case article.ErrArticleNotFound:
				err = &goa.ArticleUpdateArticleBadRequest{Code: design.ErrCode_Article_ArticleNotFound}
			case article.ErrRequestUserIsNotAuthor:
				err = &goa.ArticleUpdateArticleBadRequest{Code: design.ErrCode_Article_ForbiddenOperation}
			}
		}
	}()

	userID, err := myctx.ShouldGetAuthenticatedUserID(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	db := s.db

	articleID := uuid.MustParse(payload.ArticleID)
	detail, err := s.getArticleDetail(ctx, articleID, &userID, true)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	isUpdated := false
	if payload.Title != nil {
		detail.Title = *payload.Title
		isUpdated = true
	}
	if payload.Description != nil {
		detail.Description = *payload.Description
		isUpdated = true
	}
	if payload.Body != nil {
		detail.Body = *payload.Body
		isUpdated = true
	}

	res = &goa.UpdateResult{
		Article: detail,
	}

	if !isUpdated {
		return res, nil
	}

	now := mytime.Now(ctx)
	if err := myrdb.Tx(ctx, db, func(ctx context.Context, txdb myrdb.TxDB) error {
		db := txdb

		if err := sqlcgen.Q.UpdateArticleContent(ctx, db, sqlcgen.UpdateArticleContentParams{
			UpdatedAt:   now,
			ArticleID:   articleID,
			Title:       detail.Title,
			Description: detail.Description,
			Body:        detail.Body,
		}); err != nil {
			return errors.WithStack(err)
		}

		if err := sqlcgen.Q.InsertArticleContentMutation(ctx, db, sqlcgen.InsertArticleContentMutationParams{
			CreatedAt:    now,
			ArticleID:    articleID,
			Title:        detail.Title,
			Description:  detail.Description,
			Body:         detail.Body,
			AuthorUserID: userID,
		}); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}

func (s *Article) Delete(ctx context.Context, payload *goa.DeletePayload) (err error) {
	defer func() {
		if apErr, ok := myerr.AsAppErr(err); ok {
			switch apErr {
			case article.ErrArticleNotFound:
				err = &goa.ArticleDeleteArticleBadRequest{Code: design.ErrCode_Article_ArticleNotFound}
			case article.ErrRequestUserIsNotAuthor:
				err = &goa.ArticleDeleteArticleBadRequest{Code: design.ErrCode_Article_ForbiddenOperation}
			}
		}
	}()

	userID, err := myctx.ShouldGetAuthenticatedUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	db := s.db

	articleID := uuid.MustParse(payload.ArticleID)

	if _, err = s.getArticleDetail(ctx, articleID, &userID, true); err != nil {
		return errors.WithStack(err)
	}

	now := mytime.Now(ctx)
	if err := myrdb.Tx(ctx, db, func(ctx context.Context, txdb myrdb.TxDB) error {
		db := txdb

		if err := sqlcgen.Q.DeleteArticleContent(ctx, db, articleID); err != nil {
			return errors.WithStack(err)
		}

		if err := sqlcgen.Q.InsertArticleDeleted(ctx, db, sqlcgen.InsertArticleDeletedParams{
			CreatedAt: now,
			ArticleID: articleID,
		}); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *Article) Favorite(ctx context.Context, payload *goa.FavoritePayload) (res *goa.FavoriteResult, err error) {
	defer func() {
		if apErr, ok := myerr.AsAppErr(err); ok {
			switch apErr {
			case article.ErrArticleNotFound:
				err = &goa.ArticleFavoriteArticleBadRequest{Code: design.ErrCode_Article_ArticleNotFound}
			}
		}
	}()

	userID, err := myctx.ShouldGetAuthenticatedUserID(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	db := s.db
	articleID := uuid.MustParse(payload.ArticleID)

	detail, err := s.getArticleDetail(ctx, articleID, &userID, false)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if !detail.Favorited {
		now := mytime.Now(ctx)
		if err := myrdb.Tx(ctx, db, func(ctx context.Context, txdb myrdb.TxDB) error {
			db := txdb

			if err := sqlcgen.Q.InsertArticleFavorite(ctx, db, sqlcgen.InsertArticleFavoriteParams{
				CreatedAt: now,
				ArticleID: articleID,
				UserID:    userID,
			}); err != nil {
				return errors.WithStack(err)
			}

			if err := sqlcgen.Q.InsertArticleFavoriteMutation(ctx, db, sqlcgen.InsertArticleFavoriteMutationParams{
				CreatedAt: now,
				ArticleID: articleID,
				UserID:    userID,
				Type:      sqlcgen.ArticleFavoriteMutationTypeFavorite,
			}); err != nil {
				return errors.WithStack(err)
			}

			stats, err := sqlcgen.Q.GetArticleStatsByArticleIDForUpdate(ctx, db, articleID)
			if err != nil {
				return errors.WithStack(err)
			}
			favoriteCount := stats.FavoritesCount + 1

			if err := sqlcgen.Q.UpdateArticleStatsFavoritesCount(ctx, db, sqlcgen.UpdateArticleStatsFavoritesCountParams{
				FavoritesCount: favoriteCount,
				ArticleID:      articleID,
			}); err != nil {
				return errors.WithStack(err)
			}

			detail.Favorited = true
			detail.FavoritesCount = uint(favoriteCount)

			return nil
		}); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return &goa.FavoriteResult{Article: detail}, nil
}

func (s *Article) Unfavorite(ctx context.Context, payload *goa.UnfavoritePayload) (res *goa.UnfavoriteResult, err error) {
	defer func() {
		if apErr, ok := myerr.AsAppErr(err); ok {
			switch apErr {
			case article.ErrArticleNotFound:
				err = &goa.ArticleUnfavoriteArticleBadRequest{Code: design.ErrCode_Article_ArticleNotFound}
			}
		}
	}()

	userID, err := myctx.ShouldGetAuthenticatedUserID(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	db := s.db
	articleID := uuid.MustParse(payload.ArticleID)

	detail, err := s.getArticleDetail(ctx, articleID, &userID, false)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if detail.Favorited {
		now := mytime.Now(ctx)
		if err := myrdb.Tx(ctx, db, func(ctx context.Context, txdb myrdb.TxDB) error {
			db := txdb

			if err := sqlcgen.Q.DeleteArticleFavorite(ctx, db, sqlcgen.DeleteArticleFavoriteParams{
				ArticleID: articleID,
				UserID:    userID,
			}); err != nil {
				return errors.WithStack(err)
			}

			if err := sqlcgen.Q.InsertArticleFavoriteMutation(ctx, db, sqlcgen.InsertArticleFavoriteMutationParams{
				CreatedAt: now,
				ArticleID: articleID,
				UserID:    userID,
				Type:      sqlcgen.ArticleFavoriteMutationTypeUnfavorite,
			}); err != nil {
				return errors.WithStack(err)
			}

			stats, err := sqlcgen.Q.GetArticleStatsByArticleIDForUpdate(ctx, db, articleID)
			if err != nil {
				return errors.WithStack(err)
			}
			favoriteCount := stats.FavoritesCount - 1

			if err := sqlcgen.Q.UpdateArticleStatsFavoritesCount(ctx, db, sqlcgen.UpdateArticleStatsFavoritesCountParams{
				FavoritesCount: favoriteCount,
				ArticleID:      articleID,
			}); err != nil {
				return errors.WithStack(err)
			}

			detail.Favorited = false
			detail.FavoritesCount = uint(favoriteCount)

			return nil
		}); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return &goa.UnfavoriteResult{Article: detail}, nil
}

func (s *Article) getArticleDetail(
	ctx context.Context,
	articleID uuid.UUID,
	requestUserIDOptional *uuid.UUID,
	requestUserShouleBeAuthor bool,
) (res *goa.ArticleDetail, err error) {
	db := s.db

	a, err := sqlcgen.Q.GetArticleContentByArticleID(ctx, db, articleID)
	if err != nil {
		if myrdb.IsErrNoRows(err) {
			return nil, article.ErrArticleNotFound
		}
		return nil, errors.WithStack(err)
	}

	if requestUserShouleBeAuthor {
		if requestUserIDOptional == nil {
			panic("requestUserIDOptional must be not nil when requestUserShouleBeAuthor is true")
		}
		if *requestUserIDOptional != a.AuthorUserID {
			return nil, article.ErrRequestUserIsNotAuthor
		}
	}

	author, err := sqlcgen.Q.GetUserProfileByUserID(ctx, db, a.AuthorUserID)
	if err != nil {
		// handle ErrNoRows as internal server error
		return nil, errors.WithStack(err)
	}

	stats, err := sqlcgen.Q.GetArticleStatsByArticleID(ctx, db, a.ArticleID)
	if err != nil {
		// handle ErrNoRows as internal server error
		return nil, errors.WithStack(err)
	}

	tags, err := sqlcgen.Q.ListArticleTagByArticleID(ctx, db, a.ArticleID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	favorited := false
	authorFollowing := false
	if requestUserIDOptional != nil {
		favorited, err = sqlcgen.Q.IsArticleFavoritedByArticleIDAndUserID(ctx, db, sqlcgen.IsArticleFavoritedByArticleIDAndUserIDParams{
			ArticleID: a.ArticleID,
			UserID:    *requestUserIDOptional,
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}

		authorFollowing, err = sqlcgen.Q.IsUserFollowing(ctx, db, sqlcgen.IsUserFollowingParams{
			UserID:         *requestUserIDOptional,
			FollowedUserID: a.AuthorUserID,
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return &goa.ArticleDetail{
		ArticleID:      a.ArticleID.String(),
		Title:          a.Title,
		Description:    a.Description,
		Body:           a.Body,
		TagList:        tags,
		CreatedAt:      a.CreatedAt.String(),
		UpdatedAt:      a.UpdatedAt.String(),
		Favorited:      favorited,
		FavoritesCount: uint(stats.FavoritesCount),
		Author: &goa.Profile{
			Username:  author.Username,
			Bio:       author.Bio,
			Image:     author.ImageUrl,
			Following: authorFollowing,
		},
	}, nil
}
