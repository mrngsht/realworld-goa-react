package service_test

import (
	"testing"
	"time"

	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest"
	"github.com/mrngsht/realworld-goa-react/mytime/mytimetest"
	"github.com/mrngsht/realworld-goa-react/service"
	"github.com/mrngsht/realworld-goa-react/service/servicetest"
)

func TestArticle_Create(t *testing.T) {
	ctx := servicetest.NewContext()
	db := rdbtest.CreateDB(t, ctx)

	svc := service.NewArticle(db)

	t.Run("succeed", func(t *testing.T) {
		executedAt := mytimetest.TruncateTimeForDB(time.Now())

	})
}
