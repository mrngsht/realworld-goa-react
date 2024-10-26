package server

import (
	"net/http"
	"strings"

	"github.com/cockroachdb/errors"
	domainUser "github.com/mrngsht/realworld-goa-react/domain/user"
	"github.com/mrngsht/realworld-goa-react/mytime"

	user "github.com/mrngsht/realworld-goa-react/gen/http/user/server"
)

type (
	ctxKeyUserID struct{}
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
			path := "/" + r.URL.Path

			if noAuthorizationRequired[path] {
				h.ServeHTTP(w, r)
				return
			}

			tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Token ")
			if tokenString == "" {
				http.Error(w, "missing or invalid authorization header", http.StatusUnauthorized)
				return
			}

			token, err := domainUser.ParseAndVerifyToken(tokenString, mytime.Now(r.Context()))
			if err != nil {
				if errors.Is(err, domainUser.ErrTokenHasExpired) {
					http.Error(w, "token has expired", http.StatusUnauthorized)
					return
				}
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			// noauthのときもユーザーがとれるときは取る
			// どういう場合にerrorとすべきか/しないべきか

			h.ServeHTTP(w, r)
		})
	}
}
