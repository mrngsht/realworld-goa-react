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
		return nil, err
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

	articles, err := s.getArticleList(ctx, articleIDs, userIDOptional)
	if err != nil {
		return nil, err
	}

	return &goa.ListResult{Articles: articles}, nil
}

func (s *Article) Feed(ctx context.Context, payload *goa.FeedPayload) (res *goa.FeedResult, err error) {
	userID, err := myctx.ShouldGetAuthenticatedUserID(ctx)
	if err != nil {
		return nil, err
	}

	db := s.db

	articleIDs, err := sqlcgen.Q.ListArticleIDsByFeed(ctx, db, sqlcgen.ListArticleIDsByFeedParams{
		Limit:  payload.Limit,
		Offset: payload.Offset,
		UserID: userID,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	articles, err := s.getArticleList(ctx, articleIDs, &userID)
	if err != nil {
		return nil, err
	}

	return &goa.FeedResult{Articles: articles}, nil
}

func (s *Article) Create(ctx context.Context, payload *goa.CreatePayload) (res *goa.CreateResult, err error) {
	userID, err := myctx.ShouldGetAuthenticatedUserID(ctx)
	if err != nil {
		return nil, err
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
		return nil, err
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
		return nil, err
	}

	db := s.db

	articleID := uuid.MustParse(payload.ArticleID)
	detail, err := s.getArticleDetail(ctx, articleID, &userID, true)
	if err != nil {
		return nil, err
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
		return nil, err
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
		return err
	}

	db := s.db

	articleID := uuid.MustParse(payload.ArticleID)

	if _, err = s.getArticleDetail(ctx, articleID, &userID, true); err != nil {
		return err
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
		return err
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
		return nil, err
	}

	db := s.db
	articleID := uuid.MustParse(payload.ArticleID)

	detail, err := s.getArticleDetail(ctx, articleID, &userID, false)
	if err != nil {
		return nil, err
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
			return nil, err
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
		return nil, err
	}

	db := s.db
	articleID := uuid.MustParse(payload.ArticleID)

	detail, err := s.getArticleDetail(ctx, articleID, &userID, false)
	if err != nil {
		return nil, err
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
			return nil, err
		}
	}

	return &goa.UnfavoriteResult{Article: detail}, nil
}

func (s *Article) getArticleDetail(
	ctx context.Context,
	articleID uuid.UUID,
	userIDOptional *uuid.UUID,
	userShouleBeAuthor bool,
) (res *goa.ArticleDetail, err error) {
	props, err := s.getArticleProperties(ctx, []uuid.UUID{articleID}, userIDOptional)
	if err != nil {
		return nil, err
	}

	a, ok := props.articleContentMap[articleID]
	if !ok {
		return nil, article.ErrArticleNotFound
	}

	if userShouleBeAuthor {
		if userIDOptional == nil {
			panic("requestUserIDOptional must be not nil when requestUserShouleBeAuthor is true")
		}
		if *userIDOptional != a.AuthorUserID {
			return nil, article.ErrRequestUserIsNotAuthor
		}
	}

	author, ok := props.authorProfileMap[a.AuthorUserID]
	if !ok {
		panic("author shouldn't be found") // if user profile can be deleted, you should handle this correctly
	}

	return &goa.ArticleDetail{
		ArticleID:      a.ArticleID.String(),
		Title:          a.Title,
		Description:    a.Description,
		Body:           a.Body,
		TagList:        props.articleTagsMap[articleID], // use zero value as it is
		CreatedAt:      a.CreatedAt.String(),
		UpdatedAt:      a.UpdatedAt.String(),
		Favorited:      props.userFavoriteArticleMap[articleID],               // use zero value as it is
		FavoritesCount: uint(props.articleStatsMap[articleID].FavoritesCount), // use zero value as it is
		Author: &goa.Profile{
			Username:  author.Username,
			Bio:       author.Bio,
			Image:     author.ImageUrl,
			Following: props.userFollowingAuthorMap[author.UserID], // use zero value as it is
		},
	}, nil
}

func (s *Article) getArticleList(
	ctx context.Context,
	articleIDs []uuid.UUID,
	userIDOptional *uuid.UUID,
) ([]*goa.ArticleSummary, error) {
	props, err := s.getArticleProperties(ctx, articleIDs, userIDOptional)
	if err != nil {
		return nil, err
	}

	articles := make([]*goa.ArticleSummary, 0, len(articleIDs))
	for _, id := range articleIDs {
		content, ok := props.articleContentMap[id]
		if !ok {
			continue
		}
		author, ok := props.authorProfileMap[content.AuthorUserID]
		if !ok {
			panic("author shouldn't be found") // if user profile can be deleted, you should handle this correctly
		}

		articles = append(articles, &goa.ArticleSummary{
			ArticleID:      id.String(),
			Title:          content.Title,
			Description:    content.Description,
			TagList:        props.articleTagsMap[id], // use zero value as it is
			CreatedAt:      content.CreatedAt.String(),
			UpdatedAt:      content.UpdatedAt.String(),
			Favorited:      props.userFavoriteArticleMap[id],               // use zero value as it is
			FavoritesCount: uint(props.articleStatsMap[id].FavoritesCount), // use zero value as it is
			Author: &goa.Profile{
				Username:  author.Username,
				Bio:       author.Bio,
				Image:     author.ImageUrl,
				Following: props.userFollowingAuthorMap[author.UserID], // use zero value as it is
			},
		})
	}

	return articles, nil
}

type articleProperties struct {
	articleContentMap      map[uuid.UUID]sqlcgen.ArticleContent
	authorProfileMap       map[uuid.UUID]sqlcgen.ListUserProfilesByUserIDsRow
	articleStatsMap        map[uuid.UUID]sqlcgen.ListArticleStatsByArticleIDsRow
	articleTagsMap         map[uuid.UUID][]string
	userFavoriteArticleMap map[uuid.UUID]bool
	userFollowingAuthorMap map[uuid.UUID]bool
}

func (s *Article) getArticleProperties(
	ctx context.Context,
	articleIDs []uuid.UUID,
	userIDOptional *uuid.UUID,
) (articleProperties, error) {
	db := s.db

	contents, err := sqlcgen.Q.ListArticleContentsByArticleIDs(ctx, db, articleIDs)
	if err != nil {
		return articleProperties{}, errors.WithStack(err)
	}

	articleContentMap := make(map[uuid.UUID]sqlcgen.ArticleContent)
	for _, c := range contents {
		articleContentMap[c.ArticleID] = c
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
			return articleProperties{}, errors.WithStack(err)
		}
		for _, p := range profiles {
			authorProfileMap[p.UserID] = p
		}
	}

	articleStatsMap := make(map[uuid.UUID]sqlcgen.ListArticleStatsByArticleIDsRow)
	{
		stats, err := sqlcgen.Q.ListArticleStatsByArticleIDs(ctx, db, articleIDs)
		if err != nil {
			return articleProperties{}, errors.WithStack(err)
		}
		for _, s := range stats {
			articleStatsMap[s.ArticleID] = s
		}
	}

	articleTagsMap := make(map[uuid.UUID][]string)
	{
		articleTags, err := sqlcgen.Q.ListArticleTagsByArticleIDs(ctx, db, articleIDs)
		if err != nil {
			return articleProperties{}, errors.WithStack(err)
		}
		for _, at := range articleTags {
			articleTagsMap[at.ArticleID] = append(articleTagsMap[at.ArticleID], at.Tag)
		}
	}

	userFavoriteArticleMap := make(map[uuid.UUID]bool)
	userFollowingAuthorMap := make(map[uuid.UUID]bool)
	if userIDOptional != nil {
		userID := *userIDOptional

		{
			favoritedArticleIDs, err := sqlcgen.Q.ListFavoritedArticlesByUserIDAndArticleIDs(ctx, db,
				sqlcgen.ListFavoritedArticlesByUserIDAndArticleIDsParams{
					UserID:     userID,
					ArticleIds: articleIDs,
				})
			if err != nil {
				return articleProperties{}, errors.WithStack(err)
			}
			for _, id := range favoritedArticleIDs {
				userFavoriteArticleMap[id] = true
			}
		}

		{
			followingAuthorIDs, err := sqlcgen.Q.ListFollowedUserIDsByUserIDAndFollowedUserIDs(ctx, db, sqlcgen.ListFollowedUserIDsByUserIDAndFollowedUserIDsParams{
				UserID:          userID,
				FollowedUserIds: authorIDs,
			})
			if err != nil {
				return articleProperties{}, errors.WithStack(err)
			}
			for _, id := range followingAuthorIDs {
				userFollowingAuthorMap[id] = true
			}
		}
	}

	return articleProperties{
		articleContentMap:      articleContentMap,
		authorProfileMap:       authorProfileMap,
		articleStatsMap:        articleStatsMap,
		articleTagsMap:         articleTagsMap,
		userFavoriteArticleMap: userFavoriteArticleMap,
		userFollowingAuthorMap: userFollowingAuthorMap,
	}, nil
}

func (s *Article) AddComments(ctx context.Context, payload *goa.AddCommentsPayload) (res *goa.AddCommentsResult, err error) {
	defer func() {
		if apErr, ok := myerr.AsAppErr(err); ok {
			switch apErr {
			case article.ErrArticleNotFound:
				err = &goa.ArticleAddCommentsBadRequest{Code: design.ErrCode_Article_ArticleNotFound}
			}
		}
	}()

	userID, err := myctx.ShouldGetAuthenticatedUserID(ctx)
	if err != nil {
		return nil, err
	}

	db := s.db
	articleID := uuid.MustParse(payload.ArticleID)

	_, err = s.getArticleDetail(ctx, articleID, &userID, false)
	if err != nil {
		return nil, err
	}

	profile, err := sqlcgen.Q.GetUserProfileByUserID(ctx, db, userID)
	if err != nil {
		// handle ErrNoRows as internal server error
		return nil, errors.WithStack(err)
	}

	now := mytime.Now(ctx)
	commentID := uuid.New()
	if err := myrdb.Tx(ctx, db, func(ctx context.Context, txdb myrdb.TxDB) error {
		db := txdb

		if err := sqlcgen.Q.InsertArticleComment(ctx, db, sqlcgen.InsertArticleCommentParams{
			CreatedAt: now,
			ID:        commentID,
		}); err != nil {
			return errors.WithStack(err)
		}

		if err := sqlcgen.Q.InsertArticleCommentContent(ctx, db, sqlcgen.InsertArticleCommentContentParams{
			CreatedAt:        now,
			ArticleCommentID: commentID,
			ArticleID:        articleID,
			Body:             payload.Body,
			UserID:           userID,
		}); err != nil {
			return errors.WithStack(err)
		}

		if err := sqlcgen.Q.InsertArticleCommentContentMutation(ctx, db, sqlcgen.InsertArticleCommentContentMutationParams{
			CreatedAt:        now,
			ArticleCommentID: commentID,
			ArticleID:        articleID,
			Body:             payload.Body,
			UserID:           userID,
		}); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &goa.AddCommentsResult{
		Comment: &goa.Comment{
			ID:        commentID.String(),
			CreatedAt: now.String(),
			UpdatedAt: now.String(),
			Body:      payload.Body,
			Author: &goa.Profile{
				Username:  profile.Username,
				Bio:       profile.Bio,
				Image:     profile.ImageUrl,
				Following: false, // author of the comment is the current user, so following is false (or true if they follow themselves? usually false in this context or self-follow check unimplemented here, simply false is safe for "you can't follow yourself")
			},
		},
	}, nil
}

func (s *Article) GetComments(ctx context.Context, payload *goa.GetCommentsPayload) (res *goa.GetCommentsResult, err error) {
	defer func() {
		if apErr, ok := myerr.AsAppErr(err); ok {
			switch apErr {
			case article.ErrArticleNotFound:
				err = &goa.ArticleGetCommentsBadRequest{Code: design.ErrCode_Article_ArticleNotFound}
			}
		}
	}()

	userIDOptional := myctx.MayGetAuthenticatedUserID(ctx)

	db := s.db
	articleID := uuid.MustParse(payload.ArticleID)

	// Check if article exists
	_, err = s.getArticleDetail(ctx, articleID, userIDOptional, false)
	if err != nil {
		return nil, err
	}

	commentContents, err := sqlcgen.Q.ListArticleCommentContentsByArticleID(ctx, db, articleID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(commentContents) == 0 {
		return &goa.GetCommentsResult{Comments: []*goa.Comment{}}, nil
	}

	userIDs := make([]uuid.UUID, 0, len(commentContents))
	seenUserIDs := make(map[uuid.UUID]bool)
	for _, c := range commentContents {
		if !seenUserIDs[c.UserID] {
			userIDs = append(userIDs, c.UserID)
			seenUserIDs[c.UserID] = true
		}
	}

	userProfiles := make(map[uuid.UUID]sqlcgen.ListUserProfilesByUserIDsRow)
	if len(userIDs) > 0 {
		profiles, err := sqlcgen.Q.ListUserProfilesByUserIDs(ctx, db, userIDs)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for _, p := range profiles {
			userProfiles[p.UserID] = p
		}
	}

	userFollowMap := make(map[uuid.UUID]bool)
	if userIDOptional != nil {
		followingAuthorIDs, err := sqlcgen.Q.ListFollowedUserIDsByUserIDAndFollowedUserIDs(ctx, db, sqlcgen.ListFollowedUserIDsByUserIDAndFollowedUserIDsParams{
			UserID:          *userIDOptional,
			FollowedUserIds: userIDs,
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for _, id := range followingAuthorIDs {
			userFollowMap[id] = true
		}
	}

	resComments := make([]*goa.Comment, 0, len(commentContents))
	for _, c := range commentContents {
		author, ok := userProfiles[c.UserID]
		if !ok {
			// Should not happen if data consistency is maintained
			continue
		}

		resComments = append(resComments, &goa.Comment{
			ID:        c.ArticleCommentID.String(),
			CreatedAt: c.CreatedAt.String(),
			UpdatedAt: c.CreatedAt.String(),
			Body:      c.Body,
			Author: &goa.Profile{
				Username:  author.Username,
				Bio:       author.Bio,
				Image:     author.ImageUrl,
				Following: userFollowMap[c.UserID],
			},
		})
	}

	return &goa.GetCommentsResult{Comments: resComments}, nil
}
