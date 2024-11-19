package service_test

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	goa "github.com/mrngsht/realworld-goa-react/gen/article"
	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest"
	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest/sqlctest"
	"github.com/mrngsht/realworld-goa-react/mytime"
	"github.com/mrngsht/realworld-goa-react/mytime/mytimetest"
	"github.com/mrngsht/realworld-goa-react/service"
	"github.com/mrngsht/realworld-goa-react/service/servicetest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArticle_Create(t *testing.T) {
	ctx := servicetest.NewContext()
	db := rdbtest.CreateDB(t, ctx)

	svc := service.NewArticle(db)

	t.Run("succeed with tag", func(t *testing.T) {
		u := servicetest.CreateUser(t, ctx, db)

		executedAt := mytime.Now(ctx)
		ctx := mytimetest.WithFixedNow(t, ctx, executedAt)

		ctx = servicetest.SetRequestUser(t, ctx, db, u.Username)
		payload := &goa.CreatePayload{
			Title:       "title",
			Description: "description",
			Body:        "body",
			TagList:     []string{"tag1", "tag2"},
		}
		res, err := svc.Create(ctx, payload)
		require.NoError(t, err)

		assert.NotEmpty(t, res.Article.ID)
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
		articleID := uuid.MustParse(res.Article.ID)

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

		ctx = servicetest.SetRequestUser(t, ctx, db, u.Username)
		payload := &goa.CreatePayload{
			Title:       "title",
			Description: "description",
			Body:        "body",
			TagList:     []string{},
		}
		res, err := svc.Create(ctx, payload)
		require.NoError(t, err)
		assert.Equal(t, payload.TagList, res.Article.TagList)

		articleID := uuid.MustParse(res.Article.ID)

		ats, err := sqlctest.Q.ListArticleTagByArticleID(ctx, db, articleID)
		require.NoError(t, err)
		assert.Len(t, ats, 0)

		atms, err := sqlctest.Q.ListArticleTagMutationByArticleID(ctx, db, articleID)
		require.NoError(t, err)
		assert.Len(t, atms, 0)
	})
}
