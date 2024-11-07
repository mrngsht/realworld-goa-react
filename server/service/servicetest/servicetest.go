package servicetest

import (
	"context"
	"errors"
	"testing"

	"github.com/mrngsht/realworld-goa-react/myctx"
	"github.com/mrngsht/realworld-goa-react/myrdb/rdbtest/sqlctest"
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

func SetRequestUser(t *testing.T, ctx context.Context, qt *sqlctest.Queries, username string) context.Context {
	t.Helper()
	p, err := qt.GetUserProfileByUsername(ctx, username)
	require.NoError(t, err)
	return myctx.SetRequestUserID(ctx, p.UserID)
}
