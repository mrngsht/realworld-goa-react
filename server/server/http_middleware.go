package server

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"runtime"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	domainUser "github.com/mrngsht/realworld-goa-react/domain/user"
	"github.com/mrngsht/realworld-goa-react/myctx"
	"github.com/mrngsht/realworld-goa-react/mylog"
	"github.com/mrngsht/realworld-goa-react/mytime"

	user "github.com/mrngsht/realworld-goa-react/gen/http/user/server"
)

var (
	noAuthorizationRequired = map[string]bool{
		user.LoginUserPath():    true,
		user.RegisterUserPath(): true,
	}
)

func newUserAuthorizationMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := authorize(r)
			if err != nil {
				if noAuthorizationRequired[r.URL.Path] {
					// swallow error
				} else {
					if errors.Is(err, domainUser.ErrTokenHasExpired) {
						http.Error(w, "token has expired", http.StatusUnauthorized)
						return
					}
					http.Error(w, "missing or invalid authorization token", http.StatusUnauthorized)
					return
				}
			}

			if token != nil {
				ctx := myctx.SetRequestUserID(r.Context(), token.UserID)
				r = r.WithContext(ctx)
			}

			h.ServeHTTP(w, r)
		})
	}
}

func authorize(r *http.Request) (*domainUser.Token, error) {
	authHeader := r.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Token ")
	if tokenString == "" {
		return nil, errors.New("missing or invalid authorization header")
	}

	token, err := domainUser.ParseAndVerifyToken(tokenString, mytime.Now(r.Context()))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &token, nil
}

func newPanicRecoverMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				rec := recover()
				if rec == nil {
					return
				}

				err, ok := rec.(error)
				if !ok {
					err = fmt.Errorf("%v", rec)
				}

				var stack []byte
				{
					stack = make([]byte, 4<<10) // 4kb
					length := runtime.Stack(stack, true)
					stack = stack[:length]
				}

				mylog.Error(r.Context(), "[PANIC RECOVERED]",
					"message", err.Error(),
					"stack trace", string(stack),
				)

				http.Error(w, "internal server error", http.StatusInternalServerError)
			}()

			h.ServeHTTP(w, r)
		})
	}
}

func newRequestIDMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := uuid.New().String()
			ctx := myctx.SetRequestID(r.Context(), reqID)
			r = r.WithContext(ctx)

			// NOTE: may be better to receive request id from client
			// if id := r.Header.Get("X-Request-Id"); id != "" {
			// 	ctx := myctx.SetRequestID(r.Context(), id)
			// 	r = r.WithContext(ctx)
			// }

			h.ServeHTTP(w, r)
		})
	}
}

func newRequestLogMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			started := mytime.Now(ctx)
			var reqBody string
			if r.Header.Get("Content-Type") == "application/json" {
				var buf bytes.Buffer
				tee := io.TeeReader(r.Body, &buf)
				b, err := io.ReadAll(tee)
				r.Body = io.NopCloser(&buf)
				if err != nil {
					mylog.Error(ctx, "failed to read request body") // swallow error
				} else {
					reqBody = string(b)
				}
			}

			mylog.Info(ctx, "[REQUEST]",
				"method", r.Method,
				"url", r.URL.String(),
				"from", from(r),
				"body", reqBody,
				"started", started,
			)

			rc := CaptureResponse(w)
			h.ServeHTTP(rc, r)

			ended := mytime.Now(ctx)

			mylog.Info(ctx, "[RESPONSE]",
				"method", r.Method,
				"url", r.URL.String(),
				"from", from(r),
				"status", rc.StatusCode,
				"body", rc.Content.String(),
				"started", started,
				"ended", ended,
				"latency_ms", ended.Sub(started).Milliseconds(),
			)

		})
	}
}

type ResponseCapture struct {
	OriginalWriter http.ResponseWriter
	StatusCode     int
	Content        bytes.Buffer
}

var _ http.ResponseWriter = (*ResponseCapture)(nil)

func CaptureResponse(w http.ResponseWriter) *ResponseCapture {
	return &ResponseCapture{OriginalWriter: w}
}

func (rc *ResponseCapture) Header() http.Header {
	return rc.OriginalWriter.Header()
}

func (rc *ResponseCapture) Write(b []byte) (int, error) {
	_, _ = rc.Content.Write(b) // swallow error
	return rc.OriginalWriter.Write(b)
}

func (rc *ResponseCapture) WriteHeader(statusCode int) {
	rc.StatusCode = statusCode
	rc.OriginalWriter.WriteHeader(statusCode)
}

func from(req *http.Request) string {
	if f := req.Header.Get("X-Forwarded-For"); f != "" {
		return f
	}
	f := req.RemoteAddr
	ip, _, err := net.SplitHostPort(f)
	if err != nil {
		return f
	}
	return ip
}
