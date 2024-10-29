package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime"
	"strings"

	"github.com/cockroachdb/errors"
	domainUser "github.com/mrngsht/realworld-goa-react/domain/user"
	"github.com/mrngsht/realworld-goa-react/myctx"
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

func panicRecoverMiddleware() func(http.Handler) http.Handler {
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

				slog.ErrorContext(r.Context(), "PANIC RECOVERED",
					"message", err.Error(),
					"stack trace", string(stack),
				)

				http.Error(w, "internal server error", http.StatusInternalServerError)
			}()

			h.ServeHTTP(w, r)
		})
	}
}
