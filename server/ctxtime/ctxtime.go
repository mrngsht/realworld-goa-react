package ctxtime

import (
	"context"
	"time"

	"github.com/mrngsht/realworld-goa-react/ctxtime/internal"
)

func Now(ctx context.Context) time.Time {
	return internal.Now(ctx)
}
