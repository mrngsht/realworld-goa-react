package mytime

import (
	"context"
	"time"

	"github.com/mrngsht/realworld-goa-react/mytime/internal"
)

func Now(ctx context.Context) time.Time {
	return internal.Now(ctx)
}
