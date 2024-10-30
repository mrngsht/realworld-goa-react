package myctx

import (
	"context"

	"github.com/google/uuid"
)

type (
	ctxKeyRequestUserID struct{}
	ctxKeyRequestID     struct{}
)

func SetRequestUserID(ctx context.Context, userID uuid.UUID) context.Context {
	return context.WithValue(ctx, ctxKeyRequestUserID{}, userID)
}

func MustGetRequestUserID(ctx context.Context) uuid.UUID {
	userID := GetRequestUserID(ctx)
	if userID == nil {
		panic("ctx value of request userID must not be nil")
	}
	return *userID
}

func GetRequestUserID(ctx context.Context) *uuid.UUID {
	v := ctx.Value(ctxKeyRequestUserID{})
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
