package service_test

import (
	"testing"
	"time"

	"github.com/mrngsht/realworld-goa-react/design"
	"github.com/mrngsht/realworld-goa-react/gen/user"
	"github.com/mrngsht/realworld-goa-react/myrdb"
	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest"
	"github.com/mrngsht/realworld-goa-react/mytime/mytimetest"
	"github.com/mrngsht/realworld-goa-react/service"
	"github.com/mrngsht/realworld-goa-react/service/servicetest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUser_Login(t *testing.T) {
	ctx := servicetest.NewContext()
	rdb, _, _ := rdbtest.CreateRDB(t, ctx)

	svc := service.NewUser(rdb)

	t.Run("succeed", func(t *testing.T) {
		registerPayload := &user.RegisterPayload{
			Username: "succeed",
			Email:    "succeed@example.com",
			Password: "password",
		}
		_, err := svc.Register(ctx, registerPayload)
		require.NoError(t, err)

		res, err := svc.Login(ctx, &user.LoginPayload{ // Act
			Email:    registerPayload.Email,
			Password: registerPayload.Password,
		})
		require.NoError(t, err)

		assert.Equal(t, registerPayload.Email, res.User.Email)
		assert.Equal(t, registerPayload.Username, res.User.Username)
		assert.Equal(t, "", res.User.Bio)
		assert.Equal(t, "", res.User.Image)
		assert.NotEmpty(t, res.User.Token)
	})

	t.Run("email not found", func(t *testing.T) {
		registerPayload := &user.RegisterPayload{
			Username: "emailnotfound",
			Email:    "emailnotfound@example.com",
			Password: "password",
		}
		_, err := svc.Register(ctx, registerPayload)
		require.NoError(t, err)

		res, err := svc.Login(ctx, &user.LoginPayload{ // Act
			Email:    "WRONG_EMAIL_ADDRESS@example.com",
			Password: registerPayload.Password,
		})
		require.Error(t, err)
		assert.Equal(t, design.ErrorUser_EmailNotFound, servicetest.GoaServiceErrorName(err))

		assert.Empty(t, res)
	})

	t.Run("password is incorrect", func(t *testing.T) {
		registerPayload := &user.RegisterPayload{
			Username: "incorrectpass",
			Email:    "incorrectpass@example.com",
			Password: "password",
		}
		_, err := svc.Register(ctx, registerPayload)
		require.NoError(t, err)

		res, err := svc.Login(ctx, &user.LoginPayload{ // Act
			Email:    registerPayload.Email,
			Password: "INCORRECT_PASSWORD",
		})
		require.Error(t, err)
		assert.Equal(t, design.ErrorUser_PasswordIsIncorrect, servicetest.GoaServiceErrorName(err))

		assert.Empty(t, res)
	})
}

func TestUser_Register(t *testing.T) {
	ctx := servicetest.NewContext()
	rdb, _, qt := rdbtest.CreateRDB(t, ctx)

	svc := service.NewUser(rdb)

	t.Run("succeed", func(t *testing.T) {
		executedAt := mytimetest.AdjustTimeForTest(time.Now())
		ctx := mytimetest.WithFixedNow(t, ctx, executedAt)
		payload := &user.RegisterPayload{
			Username: "succeed",
			Email:    "succeed@example.com",
			Password: "password",
		}
		res, err := svc.Register(ctx, payload) // Act
		require.NoError(t, err)

		assert.Equal(t, payload.Username, res.User.Username)
		assert.Equal(t, payload.Email, res.User.Email)
		assert.Equal(t, "TODO", res.User.Token) // FIXME
		assert.Equal(t, "", res.User.Bio)
		assert.Equal(t, "", res.User.Image)

		p, err := qt.GetUserProfileByUsername(ctx, payload.Username)
		require.NoError(t, err)
		assert.Equal(t, "", p.Bio)
		assert.Equal(t, "", p.ImageUrl)
		assert.Equal(t, executedAt, p.CreatedAt)

		pms, err := qt.ListUserProfileMutationByUserID(ctx, p.UserID)
		require.NoError(t, err)
		require.Len(t, pms, 1)
		pm := pms[0]
		assert.Equal(t, payload.Username, pm.Username)
		assert.Equal(t, "", pm.Bio)
		assert.Equal(t, "", pm.ImageUrl)
		assert.Equal(t, executedAt, pm.CreatedAt)

		e, err := qt.GetUserEmailByID(ctx, p.UserID)
		require.NoError(t, err)
		assert.Equal(t, payload.Email, e.Email)
		assert.Equal(t, executedAt, e.CreatedAt)

		ems, err := qt.ListUserEmailMutationByUserID(ctx, p.UserID)
		require.NoError(t, err)
		require.Len(t, ems, 1)
		em := ems[0]
		assert.Equal(t, payload.Email, em.Email)
		assert.Equal(t, executedAt, em.CreatedAt)

		u, err := qt.GetUserByID(ctx, p.UserID)
		require.NoError(t, err)
		assert.Equal(t, executedAt, u.CreatedAt)
	})

	t.Run("username already used", func(t *testing.T) {
		{ // 1st: expect to succeed
			payload := &user.RegisterPayload{
				Username: "dup_username",
				Email:    "dup_username@example.com",
				Password: "password",
			}
			_, err := svc.Register(ctx, payload)
			require.NoError(t, err)
		}

		{ // 2nd: expect to fail
			payload := &user.RegisterPayload{
				Username: "dup_username",
				Email:    "dup_username_different_email@example.com",
				Password: "password",
			}
			_, err := svc.Register(ctx, payload) // Act
			require.Error(t, err)
			assert.Equal(t, design.ErrorUser_UsernameAlreadyUsed, servicetest.GoaServiceErrorName(err))

			_, err = qt.GetUserEmailByEmail(ctx, payload.Email)
			require.Error(t, err)
			assert.True(t, myrdb.IsErrNoRows(err))
		}
	})

	t.Run("email already used", func(t *testing.T) {
		{ // 1st: expect to succeed
			payload := &user.RegisterPayload{
				Username: "dup_email",
				Email:    "dup_email@example.com",
				Password: "password",
			}
			_, err := svc.Register(ctx, payload)
			require.NoError(t, err)
		}

		{ // 2nd: expect to fail
			payload := &user.RegisterPayload{
				Username: "dup_email_different_username",
				Email:    "dup_email@example.com",
				Password: "password",
			}
			_, err := svc.Register(ctx, payload) // Act
			require.Error(t, err)
			assert.Equal(t, design.ErrorUser_EmailAlreadyUsed, servicetest.GoaServiceErrorName(err))

			_, err = qt.GetUserProfileByUsername(ctx, payload.Username)
			require.Error(t, err)
			assert.True(t, myrdb.IsErrNoRows(err))
		}
	})
}
