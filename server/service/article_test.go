package service_test

import (
	"testing"
	"time"

	goa "github.com/mrngsht/realworld-goa-react/gen/article"
	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest"
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
		executedAt := mytimetest.TruncateTimeForDB(time.Now())

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

	})

	t.Run("succeed without tag", func(t *testing.T) {

	})
}
