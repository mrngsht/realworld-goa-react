package mylog

import (
	"context"
	"log/slog"
	"os"

	"github.com/mrngsht/realworld-goa-react/myctx"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func Debug(ctx context.Context, msg string, args ...any) {
	slog.DebugContext(ctx, msg, extendArgs(ctx, args)...)
}

func Info(ctx context.Context, msg string, args ...any) {
	slog.InfoContext(ctx, msg, extendArgs(ctx, args)...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	slog.WarnContext(ctx, msg, extendArgs(ctx, args)...)
}

func Error(ctx context.Context, msg string, args ...any) {
	slog.ErrorContext(ctx, msg, extendArgs(ctx, args)...)
}

func extendArgs(ctx context.Context, args []any) []any {
	return append(args, "requestID", myctx.GetRequestID(ctx))
}
