package service_test

import (
	"testing"
	"time"

	"github.com/guregu/null"
	"github.com/mrngsht/realworld-goa-react/design"
	goa "github.com/mrngsht/realworld-goa-react/gen/user"
	"github.com/mrngsht/realworld-goa-react/myrdb"
	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest"
	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest/sqlctest"
	"github.com/mrngsht/realworld-goa-react/mytime/mytimetest"
	"github.com/mrngsht/realworld-goa-react/service"
	"github.com/mrngsht/realworld-goa-react/service/servicetest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUser_Login(t *testing.T) {
	ctx := servicetest.NewContext()
	db := rdbtest.CreateDB(t, ctx)

	svc := service.NewUser(db)

	t.Run("succeed", func(t *testing.T) {
		registerPayload := &goa.RegisterPayload{
			Username: "succeed",
			Email:    "succeed@example.com",
			Password: "password",
		}
		_, err := svc.Register(ctx, registerPayload)
		require.NoError(t, err)

		res, err := svc.Login(ctx, &goa.LoginPayload{ // Act
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
		registerPayload := &goa.RegisterPayload{
			Username: "emailnotfound",
			Email:    "emailnotfound@example.com",
			Password: "password",
		}
		_, err := svc.Register(ctx, registerPayload)
		require.NoError(t, err)

		res, err := svc.Login(ctx, &goa.LoginPayload{ // Act
			Email:    "WRONG_EMAIL_ADDRESS@example.com",
			Password: registerPayload.Password,
		})
		require.Error(t, err)
		assert.Equal(t, design.ErrorUser_EmailNotFound, servicetest.GoaServiceErrorName(err))

		assert.Empty(t, res)
	})

	t.Run("password is incorrect", func(t *testing.T) {
		registerPayload := &goa.RegisterPayload{
			Username: "incorrectpass",
			Email:    "incorrectpass@example.com",
			Password: "password",
		}
		_, err := svc.Register(ctx, registerPayload)
		require.NoError(t, err)

		res, err := svc.Login(ctx, &goa.LoginPayload{ // Act
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
	db := rdbtest.CreateDB(t, ctx)

	svc := service.NewUser(db)

	t.Run("succeed", func(t *testing.T) {
		executedAt := mytimetest.TruncateTimeForDB(time.Now())
		ctx := mytimetest.WithFixedNow(t, ctx, executedAt)
		payload := &goa.RegisterPayload{
			Username: "succeed",
			Email:    "succeed@example.com",
			Password: "password",
		}
		res, err := svc.Register(ctx, payload) // Act
		require.NoError(t, err)

		assert.Equal(t, payload.Username, res.User.Username)
		assert.Equal(t, payload.Email, res.User.Email)
		assert.NotEmpty(t, res.User.Token)
		assert.Equal(t, "", res.User.Bio)
		assert.Equal(t, "", res.User.Image)

		p, err := sqlctest.Q.GetUserProfileByUsername(ctx, db, payload.Username)
		require.NoError(t, err)
		assert.Equal(t, "", p.Bio)
		assert.Equal(t, "", p.ImageUrl)
		assert.Equal(t, executedAt, p.CreatedAt)

		pms, err := sqlctest.Q.ListUserProfileMutationByUserID(ctx, db, p.UserID)
		require.NoError(t, err)
		require.Len(t, pms, 1)
		pm := pms[0]
		assert.Equal(t, payload.Username, pm.Username)
		assert.Equal(t, "", pm.Bio)
		assert.Equal(t, "", pm.ImageUrl)
		assert.Equal(t, executedAt, pm.CreatedAt)

		e, err := sqlctest.Q.GetUserEmailByID(ctx, db, p.UserID)
		require.NoError(t, err)
		assert.Equal(t, payload.Email, e.Email)
		assert.Equal(t, executedAt, e.CreatedAt)

		ems, err := sqlctest.Q.ListUserEmailMutationByUserID(ctx, db, p.UserID)
		require.NoError(t, err)
		require.Len(t, ems, 1)
		em := ems[0]
		assert.Equal(t, payload.Email, em.Email)
		assert.Equal(t, executedAt, em.CreatedAt)

		u, err := sqlctest.Q.GetUserByID(ctx, db, p.UserID)
		require.NoError(t, err)
		assert.Equal(t, executedAt, u.CreatedAt)
	})

	t.Run("username already used", func(t *testing.T) {
		{ // 1st: expect to succeed
			payload := &goa.RegisterPayload{
				Username: "dup_username",
				Email:    "dup_username@example.com",
				Password: "password",
			}
			_, err := svc.Register(ctx, payload)
			require.NoError(t, err)
		}

		{ // 2nd: expect to fail
			payload := &goa.RegisterPayload{
				Username: "dup_username",
				Email:    "dup_username_different_email@example.com",
				Password: "password",
			}
			_, err := svc.Register(ctx, payload) // Act
			require.Error(t, err)
			assert.Equal(t, design.ErrorUser_UsernameAlreadyUsed, servicetest.GoaServiceErrorName(err))

			_, err = sqlctest.Q.GetUserEmailByEmail(ctx, db, payload.Email)
			require.Error(t, err)
			assert.True(t, myrdb.IsErrNoRows(err))
		}
	})

	t.Run("email already used", func(t *testing.T) {
		{ // 1st: expect to succeed
			payload := &goa.RegisterPayload{
				Username: "dup_email",
				Email:    "dup_email@example.com",
				Password: "password",
			}
			_, err := svc.Register(ctx, payload)
			require.NoError(t, err)
		}

		{ // 2nd: expect to fail
			payload := &goa.RegisterPayload{
				Username: "dup_email_different_username",
				Email:    "dup_email@example.com",
				Password: "password",
			}
			_, err := svc.Register(ctx, payload) // Act
			require.Error(t, err)
			assert.Equal(t, design.ErrorUser_EmailAlreadyUsed, servicetest.GoaServiceErrorName(err))

			_, err = sqlctest.Q.GetUserProfileByUsername(ctx, db, payload.Username)
			require.Error(t, err)
			assert.True(t, myrdb.IsErrNoRows(err))
		}
	})
}

func TestUser_GetCurrentUser(t *testing.T) {
	ctx := servicetest.NewContext()
	db := rdbtest.CreateDB(t, ctx)

	svc := service.NewUser(db)

	t.Run("succeed", func(t *testing.T) {
		registerPayload := &goa.RegisterPayload{
			Username: "succeed",
			Email:    "succeed@example.com",
			Password: "password",
		}
		_, err := svc.Register(ctx, registerPayload)
		require.NoError(t, err)

		ctx = servicetest.SetRequestUser(t, ctx, db, registerPayload.Username)

		res, err := svc.GetCurrent(ctx) // Act
		require.NoError(t, err)

		assert.Equal(t, registerPayload.Email, res.User.Email)
		assert.Equal(t, registerPayload.Username, res.User.Username)
		assert.Equal(t, "", res.User.Bio)
		assert.Equal(t, "", res.User.Image)
		assert.NotEmpty(t, res.User.Token)
	})
}

func TestUser_Update(t *testing.T) {
	ctx := servicetest.NewContext()
	db := rdbtest.CreateDB(t, ctx)

	svc := service.NewUser(db)

	t.Run("update all", func(t *testing.T) {
		registerPayload := &goa.RegisterPayload{
			Username: "all",
			Email:    "all@example.com",
			Password: "all",
		}
		_, err := svc.Register(ctx, registerPayload)
		require.NoError(t, err)

		_, err = svc.Login(ctx, &goa.LoginPayload{
			Email:    registerPayload.Email,
			Password: registerPayload.Password,
		})
		require.NoError(t, err)

		ctx = servicetest.SetRequestUser(t, ctx, db, registerPayload.Username)

		updatePayload := &goa.UpdatePayload{
			Username: null.StringFrom("update_all").Ptr(),
			Email:    null.StringFrom("update_all@example.com").Ptr(),
			Password: null.StringFrom("update_all").Ptr(),
			Image:    null.StringFrom("http://example.com/file/update_all.png").Ptr(),
			Bio:      null.StringFrom("update_all").Ptr(),
		}
		res, err := svc.Update(ctx, updatePayload) // Act
		require.NoError(t, err)

		assert.Equal(t, *updatePayload.Username, res.User.Username)
		assert.Equal(t, *updatePayload.Email, res.User.Email)
		assert.Equal(t, *updatePayload.Image, res.User.Image)
		assert.Equal(t, *updatePayload.Bio, res.User.Bio)

		curr, err := svc.GetCurrent(ctx)
		require.NoError(t, err)
		assert.Equal(t, *curr.User, *res.User)

		_, err = svc.Login(ctx, &goa.LoginPayload{
			Email:    *updatePayload.Email,
			Password: *updatePayload.Password,
		})
		require.NoError(t, err)

		_, err = svc.Login(ctx, &goa.LoginPayload{
			Email:    registerPayload.Email,
			Password: registerPayload.Password,
		})
		assert.Error(t, err)
	})

	t.Run("update email only", func(t *testing.T) {
		registerPayload := &goa.RegisterPayload{
			Username: "emailonly",
			Email:    "emailonly@example.com",
			Password: "emailonly",
		}
		_, err := svc.Register(ctx, registerPayload)
		require.NoError(t, err)

		_, err = svc.Login(ctx, &goa.LoginPayload{
			Email:    registerPayload.Email,
			Password: registerPayload.Password,
		})
		require.NoError(t, err)

		ctx = servicetest.SetRequestUser(t, ctx, db, registerPayload.Username)

		updatePayload := &goa.UpdatePayload{
			Username: nil,
			Email:    null.StringFrom("update_emailonly@example.com").Ptr(),
			Password: nil,
			Image:    nil,
			Bio:      nil,
		}
		res, err := svc.Update(ctx, updatePayload) // Act
		require.NoError(t, err)

		assert.Equal(t, registerPayload.Username, res.User.Username)
		assert.Equal(t, *updatePayload.Email, res.User.Email)
		assert.Equal(t, "", res.User.Image)
		assert.Equal(t, "", res.User.Bio)

		curr, err := svc.GetCurrent(ctx)
		require.NoError(t, err)
		assert.Equal(t, *curr.User, *res.User)

		_, err = svc.Login(ctx, &goa.LoginPayload{
			Email:    *updatePayload.Email,
			Password: registerPayload.Password,
		})
		require.NoError(t, err)
	})

	t.Run("update password only", func(t *testing.T) {
		registerPayload := &goa.RegisterPayload{
			Username: "passwordonly",
			Email:    "passwordonly@example.com",
			Password: "passwordonly",
		}
		_, err := svc.Register(ctx, registerPayload)
		require.NoError(t, err)

		_, err = svc.Login(ctx, &goa.LoginPayload{
			Email:    registerPayload.Email,
			Password: registerPayload.Password,
		})
		require.NoError(t, err)

		ctx = servicetest.SetRequestUser(t, ctx, db, registerPayload.Username)

		updatePayload := &goa.UpdatePayload{
			Username: nil,
			Email:    nil,
			Password: null.StringFrom("update_passwordonly").Ptr(),
			Image:    nil,
			Bio:      nil,
		}
		res, err := svc.Update(ctx, updatePayload) // Act
		require.NoError(t, err)

		assert.Equal(t, registerPayload.Username, res.User.Username)
		assert.Equal(t, registerPayload.Email, res.User.Email)
		assert.Equal(t, "", res.User.Image)
		assert.Equal(t, "", res.User.Bio)

		curr, err := svc.GetCurrent(ctx)
		require.NoError(t, err)
		assert.Equal(t, *curr.User, *res.User)

		_, err = svc.Login(ctx, &goa.LoginPayload{
			Email:    registerPayload.Email,
			Password: *updatePayload.Password,
		})
		require.NoError(t, err)

		_, err = svc.Login(ctx, &goa.LoginPayload{
			Email:    registerPayload.Email,
			Password: registerPayload.Password,
		})
		assert.Error(t, err)
	})

	t.Run("update bio only", func(t *testing.T) { // on behalf of user_profile_
		registerPayload := &goa.RegisterPayload{
			Username: "bioonly",
			Email:    "bioonly@example.com",
			Password: "bioonly",
		}
		_, err := svc.Register(ctx, registerPayload)
		require.NoError(t, err)

		_, err = svc.Login(ctx, &goa.LoginPayload{
			Email:    registerPayload.Email,
			Password: registerPayload.Password,
		})
		require.NoError(t, err)

		ctx = servicetest.SetRequestUser(t, ctx, db, registerPayload.Username)

		updatePayload := &goa.UpdatePayload{
			Username: nil,
			Email:    nil,
			Password: nil,
			Image:    nil,
			Bio:      null.StringFrom("update_bioonly").Ptr(),
		}
		res, err := svc.Update(ctx, updatePayload) // Act
		require.NoError(t, err)

		assert.Equal(t, registerPayload.Username, res.User.Username)
		assert.Equal(t, registerPayload.Email, res.User.Email)
		assert.Equal(t, "", res.User.Image)
		assert.Equal(t, *updatePayload.Bio, res.User.Bio)

		curr, err := svc.GetCurrent(ctx)
		require.NoError(t, err)
		assert.Equal(t, *curr.User, *res.User)

		_, err = svc.Login(ctx, &goa.LoginPayload{
			Email:    registerPayload.Email,
			Password: registerPayload.Password,
		})
		require.NoError(t, err)
	})

	t.Run("update nothing", func(t *testing.T) {
		registerPayload := &goa.RegisterPayload{
			Username: "nothing",
			Email:    "nothing@example.com",
			Password: "nothing",
		}
		_, err := svc.Register(ctx, registerPayload)
		require.NoError(t, err)

		_, err = svc.Login(ctx, &goa.LoginPayload{
			Email:    registerPayload.Email,
			Password: registerPayload.Password,
		})
		require.NoError(t, err)

		ctx = servicetest.SetRequestUser(t, ctx, db, registerPayload.Username)

		res, err := svc.Update(ctx, &goa.UpdatePayload{}) // Act
		require.NoError(t, err)

		assert.Equal(t, registerPayload.Username, res.User.Username)
		assert.Equal(t, registerPayload.Email, res.User.Email)
		assert.Equal(t, "", res.User.Image)
		assert.Equal(t, "", res.User.Bio)

		curr, err := svc.GetCurrent(ctx)
		require.NoError(t, err)
		assert.Equal(t, *curr.User, *res.User)

		_, err = svc.Login(ctx, &goa.LoginPayload{
			Email:    registerPayload.Email,
			Password: registerPayload.Password,
		})
		require.NoError(t, err)
	})
}
