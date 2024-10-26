package server

import (
	"net/http"
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
