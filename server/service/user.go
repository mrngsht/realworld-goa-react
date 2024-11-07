package service

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"

	"github.com/mrngsht/realworld-goa-react/domain/user"
	goa "github.com/mrngsht/realworld-goa-react/gen/user"
	"github.com/mrngsht/realworld-goa-react/myctx"
	"github.com/mrngsht/realworld-goa-react/myerr"
	"github.com/mrngsht/realworld-goa-react/myrdb"
	"github.com/mrngsht/realworld-goa-react/myrdb/sqlcgen"
	"github.com/mrngsht/realworld-goa-react/mytime"
)

type User struct {
	rdb myrdb.RDB
}

func NewUser(rdb myrdb.RDB) User {
	return User{rdb: rdb}
}

var _ goa.Service = User{}

func (u User) Login(ctx context.Context, payload *goa.LoginPayload) (res *goa.LoginResult, err error) {
	defer func() {
		if apErr, ok := myerr.AsAppErr(err); ok {
			switch apErr {
			case user.ErrEmailNotFound:
				err = goa.MakeEmailNotFound(err)
			case user.ErrPasswordIsIncorrect:
				err = goa.MakePasswordIsIncorrect(err)
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
		User: &goa.User{
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

	var userID = uuid.Nil
	if err := myrdb.Tx(ctx, u.rdb, func(ctx context.Context, tx myrdb.TxDB) error {
		q = sqlcgen.New(tx)

		now := mytime.Now(ctx)
		newUserID := uuid.New()

		if err := q.InsertUser(ctx, sqlcgen.InsertUserParams{
			CreatedAt: now,
			ID:        newUserID,
		}); err != nil {
			return errors.WithStack(err)
		}

		if err := q.InsertUserProfile(ctx, sqlcgen.InsertUserProfileParams{
			UserID:    newUserID,
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
			UserID:    newUserID,
			Username:  payload.Username,
			Bio:       "",
			ImageUrl:  "",
			CreatedAt: now,
		}); err != nil {
			return errors.WithStack(err)
		}

		if err := q.InsertUserEmail(ctx, sqlcgen.InsertUserEmailParams{
			UserID:    newUserID,
			Email:     payload.Email,
			CreatedAt: now,
		}); err != nil {
			if myrdb.IsErrUniqueViolation(err) {
				return user.ErrEmailAlreadyUsed
			}
			return errors.WithStack(err)
		}
		if err := q.InsertUserEmailMutation(ctx, sqlcgen.InsertUserEmailMutationParams{
			UserID:    newUserID,
			Email:     payload.Email,
			CreatedAt: now,
		}); err != nil {
			return errors.WithStack(err)
		}

		if err := q.InsertUserAuthPassword(ctx, sqlcgen.InsertUserAuthPasswordParams{
			UserID:       newUserID,
			PasswordHash: string(passwordHash),
			CreatedAt:    now,
		}); err != nil {
			return errors.WithStack(err)
		}

		userID = newUserID

		return nil
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	token, err := user.IssueToken(userID, mytime.Now(ctx))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &goa.RegisterResult{
		User: &goa.User{
			Email:    payload.Email,
			Username: payload.Username,
			Token:    token,
			Bio:      "",
			Image:    "",
		},
	}, nil
}

func (u User) GetCurrent(ctx context.Context) (*goa.GetCurrentResult, error) {
	q := sqlcgen.New(u.rdb)

	userID := myctx.MustGetRequestUserID(ctx)

	user, err := u.getUserByUserID(ctx, q, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &goa.GetCurrentResult{User: user}, nil
}

func (u User) Update(ctx context.Context, payload *goa.UpdatePayload) (res *goa.UpdateResult, err error) {
	q := sqlcgen.New(u.rdb)

	userID := myctx.MustGetRequestUserID(ctx)
	now := mytime.Now(ctx)

	if err := myrdb.Tx(ctx, u.rdb, func(ctx context.Context, tx myrdb.TxDB) error {
		if payload.Email != nil {
			if err := q.UpdateUserEmail(ctx, sqlcgen.UpdateUserEmailParams{
				UserID:    userID,
				UpdatedAt: now,
				Email:     *payload.Email,
			}); err != nil {
				return errors.WithStack(err)
			}

			if err := q.InsertUserEmailMutation(ctx, sqlcgen.InsertUserEmailMutationParams{
				UserID:    userID,
				Email:     *payload.Email,
				CreatedAt: now,
			}); err != nil {
				return errors.WithStack(err)
			}
		}

		if payload.Password != nil {
			passwordHash, err := user.GenPasswordHash([]byte(*payload.Password))
			if err != nil {
				return errors.WithStack(err)
			}
			if err := q.UpdateUserAuthPasswordHash(ctx, sqlcgen.UpdateUserAuthPasswordHashParams{
				UserID:       userID,
				UpdatedAt:    now,
				PasswordHash: string(passwordHash),
			}); err != nil {
				return errors.WithStack(err)
			}
		}

		if payload.Username != nil || payload.Bio != nil || payload.Image != nil {
			current, err := q.GetUserProfileByUserID(ctx, userID)
			if err != nil {
				return errors.WithStack(err)
			}

			param := sqlcgen.UpdateUserProfileParams{
				UserID:    userID,
				UpdatedAt: now,
				Username:  current.Username,
				Bio:       current.Bio,
				ImageUrl:  current.ImageUrl,
			}

			if payload.Username != nil {
				param.Username = *payload.Username
			}
			if payload.Bio != nil {
				param.Bio = *payload.Bio
			}
			if payload.Image != nil {
				param.ImageUrl = *payload.Image
			}

			if err := q.UpdateUserProfile(ctx, param); err != nil {
				return errors.WithStack(err)
			}
			if err := q.InsertUserProfileMutation(ctx, sqlcgen.InsertUserProfileMutationParams{
				CreatedAt: now,
				UserID:    userID,
				Username:  param.Username,
				Bio:       param.Bio,
				ImageUrl:  param.ImageUrl,
			}); err != nil {
				return errors.WithStack(err)
			}
		}

		return nil
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	user, err := u.getUserByUserID(ctx, q, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &goa.UpdateResult{User: user}, nil
}

func (u User) getUserByUserID(ctx context.Context, q *sqlcgen.Queries, userID uuid.UUID) (*goa.User, error) {
	email, err := q.GetUserEmailByUserID(ctx, userID)
	if err != nil {
		// handle ErrNoRows as internal server error
		return nil, errors.WithStack(err)
	}

	profile, err := q.GetUserProfileByUserID(ctx, userID)
	if err != nil {
		// handle ErrNoRows as internal server error
		return nil, errors.WithStack(err)
	}

	token, err := user.IssueToken(userID, mytime.Now(ctx))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &goa.User{
		Email:    email,
		Token:    token,
		Username: profile.Username,
		Bio:      profile.Bio,
		Image:    profile.ImageUrl,
	}, nil
}
