package service_test

import (
	"testing"
	"time"

	"github.com/mrngsht/realworld-goa-react/ctxtime/ctxtimetest"
	"github.com/mrngsht/realworld-goa-react/gen/user"
	"github.com/mrngsht/realworld-goa-react/rdb/rdbtest"
	"github.com/mrngsht/realworld-goa-react/service"
	"github.com/mrngsht/realworld-goa-react/service/servicetest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUser_Register(t *testing.T) {
	ctx := servicetest.NewContext()
	rdb, q := rdbtest.CreateRDB(t, ctx)

	svc := service.NewUser(rdb)

	t.Run("succeed", func(t *testing.T) {
		executedAt := ctxtimetest.AdjustTimeForTest(time.Now())
		ctx := ctxtimetest.WithFixedNow(t, ctx, executedAt)
		payload := &user.RegisterPayload{
			Username: "taro",
			Email:    "taro@example.com",
			Password: "taro_pass",
		}
		res, err := svc.Register(ctx, payload)
		require.NoError(t, err)

		assert.Equal(t, payload.Username, res.User.Username)
		assert.Equal(t, payload.Email, res.User.Email)
		assert.Equal(t, "TODO", res.User.Token) // FIXME
		assert.Equal(t, "", res.User.Bio)
		assert.Equal(t, "", res.User.Image)

		p, err := q.GetUserProfileByUsername(ctx, payload.Username)
		require.NoError(t, err)
		assert.Equal(t, payload.Email, p.Email)
		assert.Equal(t, "", p.Bio)
		assert.Equal(t, "", p.ImageUrl)
		assert.Equal(t, executedAt, p.CreatedAt)

		u, err := q.GetUserByID(ctx, p.UserID)
		require.NoError(t, err)
		assert.Equal(t, executedAt, u.CreatedAt)
	})
}
