package myctx

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

type (
	ctxKeyAuthenticatedUserID struct{}
	ctxKeyRequestID           struct{}
)

var (
	ErrAuthenticationRequired = errors.New("authentication required")
)

func SetAuthenticatedUserID(ctx context.Context, userID uuid.UUID) context.Context {
	return context.WithValue(ctx, ctxKeyAuthenticatedUserID{}, userID)
}

func ShouldGetAuthenticatedUserID(ctx context.Context) (uuid.UUID, error) {
	uid := getAuthenticatedUserID(ctx)
	if uid == nil {
		return uuid.Nil, ErrAuthenticationRequired
	}
	return *uid, nil
}

func MayGetAuthenticatedUserID(ctx context.Context) *uuid.UUID {
	return getAuthenticatedUserID(ctx)
}

func getAuthenticatedUserID(ctx context.Context) *uuid.UUID {
	v := ctx.Value(ctxKeyAuthenticatedUserID{})
	if v == nil {
		return nil
	}
	userID, ok := v.(uuid.UUID)
	if !ok || userID.String() == "" {
		panic("ctx value of request userID is empty or invalid")
	}
	return &userID
}

func SetRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ctxKeyRequestID{}, id)
}

func GetRequestID(ctx context.Context) string {
	v := ctx.Value(ctxKeyRequestID{})
	if v == nil {
		return ""
	}
	requestID, ok := v.(string)
	if !ok {
		return "" // fail safe
	}
	return requestID
}
