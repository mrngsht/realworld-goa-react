package servicetest

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mrngsht/realworld-goa-react/myctx"
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

func SetRequestUser(ctx context.Context, userID uuid.UUID) context.Context {
	return myctx.SetRequestUserID(ctx, userID)
}
