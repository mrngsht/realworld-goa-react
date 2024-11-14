package service_test

import (
	"testing"

	"github.com/mrngsht/realworld-goa-react/design"
	goa "github.com/mrngsht/realworld-goa-react/gen/profile"
	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest"
	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest/sqlctest"
	"github.com/mrngsht/realworld-goa-react/service"
	"github.com/mrngsht/realworld-goa-react/service/servicetest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProfile_FollowUser(t *testing.T) {
	ctx := servicetest.NewContext()
	rdb, _, qt := rdbtest.CreateRDB(t, ctx)

	svc := service.NewProfile(rdb)

	t.Run("succeed", func(t *testing.T) {
		u1 := servicetest.CreateUser(t, ctx, rdb)
		u2 := servicetest.CreateUser(t, ctx, rdb)

		ctx := servicetest.SetRequestUser(t, ctx, qt, u1.Username)
		res, err := svc.FollowUser(ctx, &goa.FollowUserPayload{ // Act
			Username: u2.Username,
		})
		require.NoError(t, err)

		assert.Equal(t, u2.Username, res.Profile.Username)
		assert.Equal(t, u2.Bio, res.Profile.Bio)
		assert.Equal(t, u2.ImageUrl, res.Profile.Image)
		assert.Equal(t, true, res.Profile.Following)

		fs, err := qt.ListUserFollowByUserID(ctx, u1.UserID)
		require.NoError(t, err)

		require.Len(t, fs, 1)
		assert.Equal(t, u2.UserID, fs[0].FollowedUserID)

		ms, err := qt.ListUserFollowMutationByUserID(ctx, u1.UserID)
		require.NoError(t, err)

		require.Len(t, ms, 1)
		assert.Equal(t, u2.UserID, ms[0].FollowedUserID)
		assert.Equal(t, sqlctest.UserFollowMutationTypeFollow, ms[0].Type)
	})

	t.Run("mutual follow", func(t *testing.T) {
		u1 := servicetest.CreateUser(t, ctx, rdb)
		u2 := servicetest.CreateUser(t, ctx, rdb)

		// u1 -> u2
		ctx := servicetest.SetRequestUser(t, ctx, qt, u1.Username)
		_, err := svc.FollowUser(ctx, &goa.FollowUserPayload{
			Username: u2.Username,
		})
		require.NoError(t, err)

		// u2 -> u1
		ctx = servicetest.SetRequestUser(t, ctx, qt, u2.Username)
		res, err := svc.FollowUser(ctx, &goa.FollowUserPayload{ // Act
			Username: u1.Username,
		})
		require.NoError(t, err)

		assert.Equal(t, u1.Username, res.Profile.Username)
		assert.Equal(t, true, res.Profile.Following)

		fs, err := qt.ListUserFollowByUserID(ctx, u2.UserID)
		require.NoError(t, err)

		require.Len(t, fs, 1)
		assert.Equal(t, u1.UserID, fs[0].FollowedUserID)
	})

	t.Run("user not found", func(t *testing.T) {
		u1 := servicetest.CreateUser(t, ctx, rdb)

		ctx := servicetest.SetRequestUser(t, ctx, qt, u1.Username)
		_, err := svc.FollowUser(ctx, &goa.FollowUserPayload{ // Act
			Username: "WRONG_USERNAME",
		})
		require.Error(t, err)
		assert.Equal(t, design.ErrorProfile_UserNotFound, servicetest.GoaServiceErrorName(err))
	})

	t.Run("user already following", func(t *testing.T) {
		u1 := servicetest.CreateUser(t, ctx, rdb)
		u2 := servicetest.CreateUser(t, ctx, rdb)

		// 1st
		ctx := servicetest.SetRequestUser(t, ctx, qt, u1.Username)
		_, err := svc.FollowUser(ctx, &goa.FollowUserPayload{
			Username: u2.Username,
		})
		require.NoError(t, err)

		// 2nd
		_, err = svc.FollowUser(ctx, &goa.FollowUserPayload{ // Act
			Username: u2.Username,
		})
		require.Error(t, err)
		assert.Equal(t, design.ErrorProfile_UserAlreadyFollowing, servicetest.GoaServiceErrorName(err))
	})
}
