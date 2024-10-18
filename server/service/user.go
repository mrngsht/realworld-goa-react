package service

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"

	"github.com/mrngsht/realworld-goa-react/domain/user"
	goa "github.com/mrngsht/realworld-goa-react/gen/user"
	"github.com/mrngsht/realworld-goa-react/myerr"
	"github.com/mrngsht/realworld-goa-react/myrdb"
	"github.com/mrngsht/realworld-goa-react/myrdb/sqlcgen"
	"github.com/mrngsht/realworld-goa-react/mytime"
)

type User struct {
	rdb *sql.DB
}

func NewUser(rdb *sql.DB) User {
	return User{rdb: rdb}
}

var _ goa.Service = User{}

func (u User) Login(ctx context.Context, payload *goa.LoginPayload) (res *goa.LoginResult, err error) {
	defer func() {
		// FIXME*
		if apErr, ok := myerr.AsAppErr(err); ok {
			switch apErr {

			}
		}
	}()

	q := sqlcgen.New(u.rdb)

	userID, err := q.GetUserIDByEmail(ctx, payload.Email)
	if err != nil {
		if myrdb.IsErrNoRows(err) {
			return nil, user.ErrEmailNotFound
		}
		return nil, errors.WithStack(err)
	}

	storedPasswordHash, err := q.GetPasswordHashByUserID(ctx, userID)
	if err != nil {
		// handle ErrNoRows as internal server error
		return nil, errors.WithStack(err)
	}

	matched, err := user.MatchPassword([]byte(storedPasswordHash), []byte(payload.Password))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if !matched {
		return nil, user.ErrPasswordIsIncorrect
	}

	token, err := user.IssueToken(userID, mytime.Now(ctx))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	profile, err := q.GetUserProfileByUserID(ctx, userID)
	if err != nil {
		// handle ErrNoRows as internal server error
		return nil, errors.WithStack(err)
	}

	return &goa.LoginResult{
		User: &goa.UserType{
			Email:    payload.Email,
			Token:    token,
			Username: profile.Username,
			Bio:      profile.Bio,
			Image:    profile.ImageUrl,
		},
	}, nil
}

func (u User) Register(ctx context.Context, payload *goa.RegisterPayload) (res *goa.RegisterResult, err error) {
	defer func() {
		if apErr, ok := myerr.AsAppErr(err); ok {
			switch apErr {
			case user.ErrUsernameAlreadyUsed:
				err = goa.MakeUsernameAlreadyUsed(err)
			case user.ErrEmailAlreadyUsed:
				err = goa.MakeEmailAlreadyUsed(err)
			}
		}
	}()

	q := sqlcgen.New(u.rdb)

	passwordHash, err := user.GenPasswordHash([]byte(payload.Password))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := myrdb.Tx(ctx, u.rdb, func(ctx context.Context, tx *sql.Tx) error {
		q = q.WithTx(tx)

		now := mytime.Now(ctx)
		userID := uuid.New()

		if err := q.InsertUser(ctx, sqlcgen.InsertUserParams{
			CreatedAt: now,
			ID:        userID,
		}); err != nil {
			return errors.WithStack(err)
		}

		if err := q.InsertUserProfile(ctx, sqlcgen.InsertUserProfileParams{
			UserID:    userID,
			Username:  payload.Username,
			Bio:       "",
			ImageUrl:  "",
			CreatedAt: now,
		}); err != nil {
			if myrdb.IsErrUniqueViolation(err) {
				return user.ErrUsernameAlreadyUsed
			}
			return errors.WithStack(err)
		}
		if err := q.InsertUserProfileMutation(ctx, sqlcgen.InsertUserProfileMutationParams{
			UserID:    userID,
			Username:  payload.Username,
			Bio:       "",
			ImageUrl:  "",
			CreatedAt: now,
		}); err != nil {
			return errors.WithStack(err)
		}

		if err := q.InsertUserEmail(ctx, sqlcgen.InsertUserEmailParams{
			UserID:    userID,
			Email:     payload.Email,
			CreatedAt: now,
		}); err != nil {
			if myrdb.IsErrUniqueViolation(err) {
				return user.ErrEmailAlreadyUsed
			}
			return errors.WithStack(err)
		}
		if err := q.InsertUserEmailMutation(ctx, sqlcgen.InsertUserEmailMutationParams{
			UserID:    userID,
			Email:     payload.Email,
			CreatedAt: now,
		}); err != nil {
			return errors.WithStack(err)
		}

		if err := q.InsertUserAuthPassword(ctx, sqlcgen.InsertUserAuthPasswordParams{
			UserID:       userID,
			PasswordHash: string(passwordHash),
			CreatedAt:    now,
		}); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return &goa.RegisterResult{
		User: &goa.UserType{
			Email:    payload.Email,
			Username: payload.Username,
			Token:    "TODO",
			Bio:      "",
			Image:    "",
		},
	}, nil
}
