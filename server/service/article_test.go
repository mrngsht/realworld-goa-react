package service_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/guregu/null"
	"github.com/mrngsht/realworld-goa-react/design"
	goa "github.com/mrngsht/realworld-goa-react/gen/article"
	"github.com/mrngsht/realworld-goa-react/gen/profile"
	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest"
	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest/sqlctest"
	"github.com/mrngsht/realworld-goa-react/mytime"
	"github.com/mrngsht/realworld-goa-react/mytime/mytimetest"
	"github.com/mrngsht/realworld-goa-react/service"
	"github.com/mrngsht/realworld-goa-react/service/servicetest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArticle_Get(t *testing.T) {
	ctx := servicetest.NewContext()
	db := rdbtest.OpenDB(t, ctx)

	svc := service.NewArticle(db)

	t.Run("succeed", func(t *testing.T) {
		author := servicetest.CreateUser(t, ctx, db)
		viewer := servicetest.CreateUser(t, ctx, db)
		other := servicetest.CreateUser(t, ctx, db)

		createdAt := mytime.Now(ctx)
		ctx := mytimetest.WithFixedNow(t, ctx, createdAt)

		ctx = servicetest.SetAuthenticatedUser(t, ctx, db, author.Username)
		payload := &goa.CreatePayload{
			Title:       "title",
			Description: "description",
			Body:        "body",
			TagList:     []string{"tag1", "tag2"},
		}
		createRes, err := svc.Create(ctx, payload)
		require.NoError(t, err)
		articleID := createRes.Article.ArticleID

		ctx = servicetest.SetAuthenticatedUser(t, ctx, db, viewer.Username)
		res, err := svc.Get(ctx, &goa.GetPayload{ArticleID: articleID}) // Act
		require.NoError(t, err)

		createdAtOnDB := mytimetest.TruncateTimeForDB(createdAt)

		assert.Equal(t, articleID, res.Article.ArticleID)
		assert.Equal(t, payload.Title, res.Article.Title)
		assert.Equal(t, payload.Description, res.Article.Description)
		assert.Equal(t, payload.Body, res.Article.Body)
		assert.Equal(t, payload.TagList, res.Article.TagList)
		assert.Equal(t, createdAtOnDB.String(), res.Article.CreatedAt)
		assert.Equal(t, createdAtOnDB.String(), res.Article.UpdatedAt)
		assert.Equal(t, false, res.Article.Favorited)
		assert.Equal(t, uint(0), res.Article.FavoritesCount)
		assert.Equal(t, author.Username, res.Article.Author.Username)
		assert.Equal(t, author.Bio, res.Article.Author.Bio)
		assert.Equal(t, author.ImageUrl, res.Article.Author.Image)
		assert.Equal(t, false, res.Article.Author.Following)

		t.Run("when the other user favorited", func(t *testing.T) {
			ctx := servicetest.SetAuthenticatedUser(t, ctx, db, other.Username)
			_, err := svc.Favorite(ctx, &goa.FavoritePayload{ArticleID: articleID})
			require.NoError(t, err)

			ctx = servicetest.SetAuthenticatedUser(t, ctx, db, viewer.Username)
			res, err := svc.Get(ctx, &goa.GetPayload{ArticleID: articleID}) // Act
			require.NoError(t, err)

			assert.Equal(t, false, res.Article.Favorited)
			assert.Equal(t, uint(1), res.Article.FavoritesCount)

			t.Run("also when user itself favorited", func(t *testing.T) {
				ctx := servicetest.SetAuthenticatedUser(t, ctx, db, viewer.Username)
				_, err := svc.Favorite(ctx, &goa.FavoritePayload{ArticleID: articleID})
				require.NoError(t, err)

				res, err := svc.Get(ctx, &goa.GetPayload{ArticleID: articleID}) // Act
				require.NoError(t, err)

				assert.Equal(t, true, res.Article.Favorited)
				assert.Equal(t, uint(2), res.Article.FavoritesCount)
			})
		})

		t.Run("when user Following", func(t *testing.T) {
			ctx = servicetest.SetAuthenticatedUser(t, ctx, db, viewer.Username)

			_, err := service.NewProfile(db).FollowUser(ctx, &profile.FollowUserPayload{
				Username: author.Username,
			})
			require.NoError(t, err)

			res, err := svc.Get(ctx, &goa.GetPayload{ArticleID: articleID}) // Act
			require.NoError(t, err)

			assert.Equal(t, true, res.Article.Author.Following)
		})
	})

	t.Run("article not exists", func(t *testing.T) {
		_, err := svc.Get(ctx, &goa.GetPayload{ArticleID: uuid.NewString()}) // Act
		require.Error(t, err)

		var badRequest *goa.ArticleGetArticleBadRequest
		require.ErrorAs(t, err, &badRequest)
		assert.Equal(t, design.ErrCode_Article_ArticleNotFound, badRequest.Code)
	})

	t.Run("article deleted", func(t *testing.T) {
		// TODO:
	})
}

func TestArticle_Create(t *testing.T) {
	ctx := servicetest.NewContext()
	db := rdbtest.OpenDB(t, ctx)

	svc := service.NewArticle(db)

	t.Run("succeed with tag", func(t *testing.T) {
		u := servicetest.CreateUser(t, ctx, db)

		executedAt := mytime.Now(ctx)
		ctx := mytimetest.WithFixedNow(t, ctx, executedAt)

		ctx = servicetest.SetAuthenticatedUser(t, ctx, db, u.Username)
		payload := &goa.CreatePayload{
			Title:       "title",
			Description: "description",
			Body:        "body",
			TagList:     []string{"tag1", "tag2"},
		}
		res, err := svc.Create(ctx, payload)
		require.NoError(t, err)

		assert.NotEmpty(t, res.Article.ArticleID)
		assert.Equal(t, payload.Title, res.Article.Title)
		assert.Equal(t, payload.Description, res.Article.Description)
		assert.Equal(t, payload.Body, res.Article.Body)
		assert.Equal(t, payload.TagList, res.Article.TagList)
		assert.Equal(t, executedAt.String(), res.Article.CreatedAt)
		assert.Equal(t, executedAt.String(), res.Article.UpdatedAt)
		assert.Equal(t, false, res.Article.Favorited)
		assert.Equal(t, uint(0), res.Article.FavoritesCount)
		assert.Equal(t, u.Username, res.Article.Author.Username)
		assert.Equal(t, u.Bio, res.Article.Author.Bio)
		assert.Equal(t, u.ImageUrl, res.Article.Author.Image)
		assert.Equal(t, false, res.Article.Author.Following)

		executedAtOnDB := mytimetest.TruncateTimeForDB(executedAt)
		articleID := uuid.MustParse(res.Article.ArticleID)

		a, err := sqlctest.Q.GetArticleByID(ctx, db, articleID)
		require.NoError(t, err)
		assert.Equal(t, executedAtOnDB, a.CreatedAt)

		ac, err := sqlctest.Q.GetArticleContentByArticleID(ctx, db, articleID)
		require.NoError(t, err)
		assert.Equal(t, executedAtOnDB, ac.CreatedAt)
		assert.Equal(t, executedAtOnDB, ac.UpdatedAt)
		assert.Equal(t, payload.Title, ac.Title)
		assert.Equal(t, payload.Description, ac.Description)
		assert.Equal(t, payload.Body, ac.Body)
		assert.Equal(t, u.UserID, ac.AuthorUserID)

		acms, err := sqlctest.Q.ListArticleContentMutationByArticleID(ctx, db, articleID)
		require.NoError(t, err)
		require.Len(t, acms, 1)
		acm := acms[0]
		assert.Equal(t, executedAtOnDB, acm.CreatedAt)
		assert.Equal(t, payload.Title, acm.Title)
		assert.Equal(t, payload.Description, acm.Description)
		assert.Equal(t, payload.Body, acm.Body)
		assert.Equal(t, u.UserID, acm.AuthorUserID)

		ats, err := sqlctest.Q.ListArticleTagByArticleID(ctx, db, articleID)
		require.NoError(t, err)
		var actualTags []string
		for _, at := range ats {
			actualTags = append(actualTags, at.Tag)
		}
		assert.Equal(t, payload.TagList, actualTags)

		atms, err := sqlctest.Q.ListArticleTagMutationByArticleID(ctx, db, articleID)
		require.NoError(t, err)
		require.Len(t, atms, 1)
		atm := atms[0]
		assert.Equal(t, executedAtOnDB, atm.CreatedAt)

		var actualTagList []string
		require.NoError(t, json.Unmarshal(atm.Tags, &actualTagList))
		assert.Equal(t, payload.TagList, actualTagList)

		as, err := sqlctest.Q.GetArticleStatsByArticleID(ctx, db, articleID)
		require.NoError(t, err)
		assert.Equal(t, int64(0), as.FavoritesCount)
	})

	t.Run("succeed without tag", func(t *testing.T) {
		u := servicetest.CreateUser(t, ctx, db)

		executedAt := mytime.Now(ctx)
		ctx := mytimetest.WithFixedNow(t, ctx, executedAt)

		ctx = servicetest.SetAuthenticatedUser(t, ctx, db, u.Username)
		payload := &goa.CreatePayload{
			Title:       "title",
			Description: "description",
			Body:        "body",
			TagList:     []string{},
		}
		res, err := svc.Create(ctx, payload)
		require.NoError(t, err)
		assert.Equal(t, payload.TagList, res.Article.TagList)

		articleID := uuid.MustParse(res.Article.ArticleID)

		ats, err := sqlctest.Q.ListArticleTagByArticleID(ctx, db, articleID)
		require.NoError(t, err)
		assert.Len(t, ats, 0)

		atms, err := sqlctest.Q.ListArticleTagMutationByArticleID(ctx, db, articleID)
		require.NoError(t, err)
		assert.Len(t, atms, 0)
	})
}

func TestArticle_Update(t *testing.T) {
	ctx := servicetest.NewContext()
	db := rdbtest.OpenDB(t, ctx)

	svc := service.NewArticle(db)

	createArticle := func(t *testing.T, ctx context.Context) (string, *goa.CreatePayload) {
		payload := &goa.CreatePayload{
			Title:       "title",
			Description: "description",
			Body:        "body",
			TagList:     []string{"tag1", "tag2"},
		}
		res, err := svc.Create(ctx, payload)
		require.NoError(t, err)
		return res.Article.ArticleID, payload
	}

	t.Run("update all", func(t *testing.T) {
		author := servicetest.CreateUser(t, ctx, db)
		ctx := servicetest.SetAuthenticatedUser(t, ctx, db, author.Username)

		articleID, _ := createArticle(t, ctx)

		updatedAt := mytime.Now(ctx)
		ctx = mytimetest.WithFixedNow(t, ctx, updatedAt)

		updatePayload := &goa.UpdatePayload{
			ArticleID:   articleID,
			Title:       null.StringFrom("updated_title").Ptr(),
			Description: null.StringFrom("updated_description").Ptr(),
			Body:        null.StringFrom("updated_body").Ptr(),
		}
		res, err := svc.Update(ctx, updatePayload) // Act
		require.NoError(t, err)

		assert.Equal(t, *updatePayload.Title, res.Article.Title)
		assert.Equal(t, *updatePayload.Description, res.Article.Description)
		assert.Equal(t, *updatePayload.Body, res.Article.Body)

		articleUUID := uuid.MustParse(articleID)
		updatedAtOnDB := mytimetest.TruncateTimeForDB(updatedAt)

		ac, err := sqlctest.Q.GetArticleContentByArticleID(ctx, db, articleUUID)
		require.NoError(t, err)
		assert.Equal(t, updatedAtOnDB, ac.UpdatedAt)
		assert.Equal(t, *updatePayload.Title, ac.Title)
		assert.Equal(t, *updatePayload.Description, ac.Description)
		assert.Equal(t, *updatePayload.Body, ac.Body)

		acms, err := sqlctest.Q.ListArticleContentMutationByArticleID(ctx, db, articleUUID)
		require.NoError(t, err)
		require.Len(t, acms, 2)
		acm := acms[1]
		assert.Equal(t, updatedAtOnDB, acm.CreatedAt)
		assert.Equal(t, *updatePayload.Title, acm.Title)
		assert.Equal(t, *updatePayload.Description, acm.Description)
		assert.Equal(t, *updatePayload.Body, acm.Body)
		assert.Equal(t, author.UserID, acm.AuthorUserID)
	})

	t.Run("update one thing", func(t *testing.T) {
		author := servicetest.CreateUser(t, ctx, db)
		ctx := servicetest.SetAuthenticatedUser(t, ctx, db, author.Username)

		articleID, createPayload := createArticle(t, ctx)

		updatePayload := &goa.UpdatePayload{
			ArticleID:   articleID,
			Title:       null.StringFrom("updated_title").Ptr(), // change this field only
			Description: nil,
			Body:        nil,
		}
		res, err := svc.Update(ctx, updatePayload) // Act
		require.NoError(t, err)

		assert.Equal(t, *updatePayload.Title, res.Article.Title)
		assert.Equal(t, createPayload.Description, res.Article.Description)
		assert.Equal(t, createPayload.Body, res.Article.Body)

		articleUUID := uuid.MustParse(articleID)

		ac, err := sqlctest.Q.GetArticleContentByArticleID(ctx, db, articleUUID)
		require.NoError(t, err)
		assert.Equal(t, *updatePayload.Title, ac.Title)
		assert.Equal(t, createPayload.Description, ac.Description)
		assert.Equal(t, createPayload.Body, ac.Body)

		acms, err := sqlctest.Q.ListArticleContentMutationByArticleID(ctx, db, articleUUID)
		require.NoError(t, err)
		require.Len(t, acms, 2)
		acm := acms[1]
		assert.Equal(t, *updatePayload.Title, acm.Title)
		assert.Equal(t, createPayload.Description, acm.Description)
		assert.Equal(t, createPayload.Body, acm.Body)
		assert.Equal(t, author.UserID, acm.AuthorUserID)
	})

	t.Run("update nothing", func(t *testing.T) {
		author := servicetest.CreateUser(t, ctx, db)
		ctx := servicetest.SetAuthenticatedUser(t, ctx, db, author.Username)

		createdAt := mytime.Now(ctx)
		ctx = mytimetest.WithFixedNow(t, ctx, createdAt)

		articleID, createPayload := createArticle(t, ctx)

		updatedAt := mytime.Now(ctx)
		ctx = mytimetest.WithFixedNow(t, ctx, updatedAt)

		updatePayload := &goa.UpdatePayload{
			ArticleID:   articleID,
			Title:       nil,
			Description: nil,
			Body:        nil,
		}
		res, err := svc.Update(ctx, updatePayload) // Act
		require.NoError(t, err)

		assert.Equal(t, createPayload.Title, res.Article.Title)
		assert.Equal(t, createPayload.Description, res.Article.Description)
		assert.Equal(t, createPayload.Body, res.Article.Body)

		articleUUID := uuid.MustParse(articleID)
		createdAtOnDB := mytimetest.TruncateTimeForDB(createdAt)

		ac, err := sqlctest.Q.GetArticleContentByArticleID(ctx, db, articleUUID)
		require.NoError(t, err)
		assert.Equal(t, createdAtOnDB, ac.UpdatedAt)
		assert.Equal(t, createPayload.Title, ac.Title)
		assert.Equal(t, createPayload.Description, ac.Description)
		assert.Equal(t, createPayload.Body, ac.Body)

		acms, err := sqlctest.Q.ListArticleContentMutationByArticleID(ctx, db, articleUUID)
		require.NoError(t, err)
		assert.Len(t, acms, 1)
	})

	t.Run("article not exists", func(t *testing.T) {
		user := servicetest.CreateUser(t, ctx, db)

		ctx = servicetest.SetAuthenticatedUser(t, ctx, db, user.Username)
		_, err := svc.Update(ctx, &goa.UpdatePayload{ArticleID: uuid.NewString()}) // Act
		require.Error(t, err)

		var badRequest *goa.ArticleUpdateArticleBadRequest
		require.ErrorAs(t, err, &badRequest)
		assert.Equal(t, design.ErrCode_Article_ArticleNotFound, badRequest.Code)
	})

	t.Run("try to update other's article", func(t *testing.T) {
		author := servicetest.CreateUser(t, ctx, db)
		other := servicetest.CreateUser(t, ctx, db)

		ctx = servicetest.SetAuthenticatedUser(t, ctx, db, author.Username)
		articleID, _ := createArticle(t, ctx)

		ctx = servicetest.SetAuthenticatedUser(t, ctx, db, other.Username)
		_, err := svc.Update(ctx, &goa.UpdatePayload{ArticleID: articleID}) // Act
		require.Error(t, err)

		var badRequest *goa.ArticleUpdateArticleBadRequest
		require.ErrorAs(t, err, &badRequest)
		assert.Equal(t, design.ErrCode_Article_ForbiddenOperation, badRequest.Code)
	})
}

func TestArticle_Favoite(t *testing.T) {
	ctx := servicetest.NewContext()
	db := rdbtest.OpenDB(t, ctx)

	svc := service.NewArticle(db)

	createArticle := func(t *testing.T, ctx context.Context) string {
		payload := &goa.CreatePayload{
			Title:       "title",
			Description: "description",
			Body:        "body",
			TagList:     []string{"tag1", "tag2"},
		}
		res, err := svc.Create(ctx, payload)
		require.NoError(t, err)
		return res.Article.ArticleID
	}

	t.Run("succeed", func(t *testing.T) {
		author := servicetest.CreateUser(t, ctx, db)
		viewer := servicetest.CreateUser(t, ctx, db)

		ctx := servicetest.SetAuthenticatedUser(t, ctx, db, author.Username)
		articleID := createArticle(t, ctx)

		ctx = servicetest.SetAuthenticatedUser(t, ctx, db, viewer.Username)
		res, err := svc.Favorite(ctx, &goa.FavoritePayload{ArticleID: articleID}) // Act
		require.NoError(t, err)

		assert.Equal(t, true, res.Article.Favorited)
		assert.Equal(t, uint(1), res.Article.FavoritesCount)
	})

	t.Run("already favorited", func(t *testing.T) {
		author := servicetest.CreateUser(t, ctx, db)
		viewer := servicetest.CreateUser(t, ctx, db)

		ctx := servicetest.SetAuthenticatedUser(t, ctx, db, author.Username)
		articleID := createArticle(t, ctx)

		{ // 1st
			ctx = servicetest.SetAuthenticatedUser(t, ctx, db, viewer.Username)
			_, err := svc.Favorite(ctx, &goa.FavoritePayload{ArticleID: articleID})
			require.NoError(t, err)
		}
		{ // 2nd
			ctx = servicetest.SetAuthenticatedUser(t, ctx, db, viewer.Username)
			res, err := svc.Favorite(ctx, &goa.FavoritePayload{ArticleID: articleID}) // Act
			require.NoError(t, err)

			assert.Equal(t, true, res.Article.Favorited)
			assert.Equal(t, uint(1), res.Article.FavoritesCount)
		}
	})

	t.Run("article not exists", func(t *testing.T) {
		viewer := servicetest.CreateUser(t, ctx, db)

		ctx = servicetest.SetAuthenticatedUser(t, ctx, db, viewer.Username)
		_, err := svc.Favorite(ctx, &goa.FavoritePayload{ArticleID: uuid.NewString()}) // Act
		require.Error(t, err)

		var badRequest *goa.ArticleFavoriteArticleBadRequest
		require.ErrorAs(t, err, &badRequest)
		assert.Equal(t, design.ErrCode_Article_ArticleNotFound, badRequest.Code)
	})
}

func TestArticle_Unfavoite(t *testing.T) {
	ctx := servicetest.NewContext()
	db := rdbtest.OpenDB(t, ctx)

	svc := service.NewArticle(db)

	createArticle := func(t *testing.T, ctx context.Context) string {
		payload := &goa.CreatePayload{
			Title:       "title",
			Description: "description",
			Body:        "body",
			TagList:     []string{"tag1", "tag2"},
		}
		res, err := svc.Create(ctx, payload)
		require.NoError(t, err)
		return res.Article.ArticleID
	}

	t.Run("succeed", func(t *testing.T) {
		author := servicetest.CreateUser(t, ctx, db)
		viewer := servicetest.CreateUser(t, ctx, db)

		ctx := servicetest.SetAuthenticatedUser(t, ctx, db, author.Username)
		articleID := createArticle(t, ctx)

		ctx = servicetest.SetAuthenticatedUser(t, ctx, db, viewer.Username)

		{
			res, err := svc.Favorite(ctx, &goa.FavoritePayload{ArticleID: articleID})
			require.NoError(t, err)

			require.Equal(t, true, res.Article.Favorited)
			require.Equal(t, uint(1), res.Article.FavoritesCount)
		}

		res, err := svc.Unfavorite(ctx, &goa.UnfavoritePayload{ArticleID: articleID}) // Act
		require.NoError(t, err)

		assert.Equal(t, false, res.Article.Favorited)
		assert.Equal(t, uint(0), res.Article.FavoritesCount)
	})

	t.Run("havn't favorited yet", func(t *testing.T) {
		author := servicetest.CreateUser(t, ctx, db)
		viewer := servicetest.CreateUser(t, ctx, db)

		ctx := servicetest.SetAuthenticatedUser(t, ctx, db, author.Username)
		articleID := createArticle(t, ctx)

		ctx = servicetest.SetAuthenticatedUser(t, ctx, db, viewer.Username)
		res, err := svc.Unfavorite(ctx, &goa.UnfavoritePayload{ArticleID: articleID}) // Act
		require.NoError(t, err)

		assert.Equal(t, false, res.Article.Favorited)
		assert.Equal(t, uint(0), res.Article.FavoritesCount)
	})

	t.Run("article not exists", func(t *testing.T) {
		viewer := servicetest.CreateUser(t, ctx, db)

		ctx = servicetest.SetAuthenticatedUser(t, ctx, db, viewer.Username)
		_, err := svc.Unfavorite(ctx, &goa.UnfavoritePayload{ArticleID: uuid.NewString()}) // Act
		require.Error(t, err)

		var badRequest *goa.ArticleUnfavoriteArticleBadRequest
		require.ErrorAs(t, err, &badRequest)
		assert.Equal(t, design.ErrCode_Article_ArticleNotFound, badRequest.Code)
	})
}
