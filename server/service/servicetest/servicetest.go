package servicetest

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/guregu/null"
	"github.com/mrngsht/realworld-goa-react/gen/user"
	"github.com/mrngsht/realworld-goa-react/myctx"
	"github.com/mrngsht/realworld-goa-react/myrdb"
	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest/sqlctest"
	"github.com/mrngsht/realworld-goa-react/service"
	"github.com/stretchr/testify/require"
	goa "goa.design/goa/v3/pkg"
)

func NewContext() context.Context {
	return context.Background()
}

func GoaServiceErrorName(err error) string {
	if serr := (*goa.ServiceError)(nil); errors.As(err, &serr) {
		return serr.GoaErrorName()
	}
	return "NOT_GOA_SERVICE_ERROR"
}

func SetAuthenticatedUser(t *testing.T, ctx context.Context, db myrdb.DB, username string) context.Context {
	t.Helper()
	p, err := sqlctest.Q.GetUserProfileByUsername(ctx, db, username)
	require.NoError(t, err)
	return myctx.SetAuthenticatedUserID(ctx, p.UserID)
}

type CreateUserResult struct {
	UserID   uuid.UUID
	Username string
	Bio      string
	ImageUrl string
}

func CreateUser(t *testing.T, ctx context.Context, db myrdb.DB) CreateUserResult {
	t.Helper()

	username := uuid.New().String()

	svc := service.NewUser(db)

	_, err := svc.Register(ctx, &user.RegisterPayload{
		Username: username,
		Email:    username + "@example.com",
		Password: "password",
	})
	require.NoError(t, err)

	p, err := sqlctest.Q.GetUserProfileByUsername(ctx, db, username)
	require.NoError(t, err)

	ctx = myctx.SetAuthenticatedUserID(ctx, p.UserID)
	u, err := svc.Update(ctx, &user.UpdatePayload{
		Image: null.StringFrom("http://example.com/file/profile.png").Ptr(),
		Bio:   null.StringFrom("hello world").Ptr(),
	})
	require.NoError(t, err)

	return CreateUserResult{
		UserID:   p.UserID,
		Username: u.User.Username,
		Bio:      u.User.Bio,
		ImageUrl: u.User.Image,
	}
}
