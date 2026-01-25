package service_test

import (
	"context"
	"testing"

	goa "github.com/mrngsht/realworld-goa-react/gen/article"
	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest"
	"github.com/mrngsht/realworld-goa-react/service"
	"github.com/mrngsht/realworld-goa-react/service/servicetest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTag_GetTags(t *testing.T) {
	ctx := servicetest.NewContext()
	db := rdbtest.OpenDB(t, ctx)

	svc := service.NewTag(db)
	articleSvc := service.NewArticle(db)

	createArticle := func(t *testing.T, ctx context.Context, author servicetest.CreateUserResult, tagList []string) {
		ctx = servicetest.SetAuthenticatedUser(t, ctx, db, author.Username)
		_, err := articleSvc.Create(ctx, &goa.CreatePayload{
			Title:       "title",
			Description: "description",
			Body:        "body",
			TagList:     tagList,
		})
		require.NoError(t, err)
	}

	t.Run("succeed with tags", func(t *testing.T) {
		author := servicetest.CreateUser(t, ctx, db)

		createArticle(t, ctx, author, []string{"golang", "testing"})
		createArticle(t, ctx, author, []string{"react", "golang"})

		res, err := svc.GetTags(ctx) // Act
		require.NoError(t, err)

		// Should return unique tags in alphabetical order
		assert.ElementsMatch(t, []string{"golang", "react", "testing"}, res.Tags)
	})

	t.Run("empty when no articles", func(t *testing.T) {
		// Use fresh context and DB for isolation
		emptyCtx := servicetest.NewContext()
		emptyDb := rdbtest.OpenDB(t, emptyCtx)
		emptySvc := service.NewTag(emptyDb)

		res, err := emptySvc.GetTags(emptyCtx) // Act
		require.NoError(t, err)
		assert.Empty(t, res.Tags)
	})
}
