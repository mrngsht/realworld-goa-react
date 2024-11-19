package mytimetest

import (
	"context"
	"testing"
	"time"

	"github.com/mrngsht/realworld-goa-react/mytime/internal"
)

func init() {
	if testing.Testing() {
		internal.Now = nowForTest
	}
}

func nowForTest(ctx context.Context) time.Time {
	tm, ok := ctx.Value(ctxkeynow{}).(time.Time)
	if !ok {
		return internal.DefaultNow(ctx)
	}
	return tm
}

type ctxkeynow struct{}

func WithFixedNow(t *testing.T, ctx context.Context, tm time.Time) context.Context {
	t.Helper()
	return context.WithValue(ctx, ctxkeynow{}, tm)
}

func AdjustTimeForTest(tm time.Time) time.Time {
	return tm.Truncate(time.Microsecond).Local()
}
