package user

import (
	"github.com/mrngsht/realworld-goa-react/myerr"
)

var (
	ErrUsernameAlreadyUsed  = myerr.NewAppErr("username already used")
	ErrEmailAlreadyUsed     = myerr.NewAppErr("email already used")
	ErrEmailNotFound        = myerr.NewAppErr("email not found")
	ErrPasswordIsIncorrect  = myerr.NewAppErr("password is incorrect")
	ErrTokenHasExpired      = myerr.NewAppErr("token has expired")
	ErrUserNotFound         = myerr.NewAppErr("user not found")
	ErrUserAlreadyFollowing = myerr.NewAppErr("user already following")
	ErrUserNotFollowing     = myerr.NewAppErr("user not following")
	ErrCannotFollowYourself = myerr.NewAppErr("cannot follow yourself")
)
