package server

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/mrngsht/realworld-goa-react/design"
	"github.com/mrngsht/realworld-goa-react/myctx"
	"github.com/mrngsht/realworld-goa-react/mylog"
	goa "goa.design/goa/v3/pkg"
)

func newErrorHandlerMiddleware() func(goa.Endpoint) goa.Endpoint {
	return func(e goa.Endpoint) goa.Endpoint {
		return goa.Endpoint(func(ctx context.Context, req interface{}) (interface{}, error) {
			res, err := e(ctx, req)
			if err == nil {
				return res, err
			}
			if errors.Is(err, myctx.ErrAuthenticationRequired) {
				return res, goa.NewServiceError(err, design.ErrorCommon_AuthenticationRequired, false, false, false)
			}
			if serr := (*goa.ServiceError)(nil); errors.As(err, &serr) {
				//already handled
				return res, err
			}

			mylog.Error(ctx, "[UNHANDLED ERROR]", "err", fmt.Sprintf("%+v", err))
			return res, goa.NewServiceError(errors.New("internal server error"), "internal server error", false, false, true)
		})
	}
}
