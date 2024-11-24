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
	db myrdb.DB
}

func NewUser(rdb myrdb.DB) *User {
	return &User{db: rdb}
}

var _ goa.Service = &User{}

func (s *User) Login(ctx context.Context, payload *goa.LoginPayload) (res *goa.LoginResult, err error) {
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

	db := s.db

	userID, err := sqlcgen.Q.GetUserIDByEmail(ctx, db, payload.Email)
	if err != nil {
		if myrdb.IsErrNoRows(err) {
			return nil, user.ErrEmailNotFound
		}
		return nil, errors.WithStack(err)
	}

	storedPasswordHash, err := sqlcgen.Q.GetPasswordHashByUserID(ctx, db, userID)
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

	profile, err := sqlcgen.Q.GetUserProfileByUserID(ctx, db, userID)
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

func (s *User) Register(ctx context.Context, payload *goa.RegisterPayload) (res *goa.RegisterResult, err error) {
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

	passwordHash, err := user.GenPasswordHash([]byte(payload.Password))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userID := uuid.New()
	if err := myrdb.Tx(ctx, s.db, func(ctx context.Context, txdb myrdb.TxDB) error {
		db := txdb

		now := mytime.Now(ctx)
		if err := sqlcgen.Q.InsertUser(ctx, db, sqlcgen.InsertUserParams{
			CreatedAt: now,
			ID:        userID,
		}); err != nil {
			return errors.WithStack(err)
		}

		if err := sqlcgen.Q.InsertUserProfile(ctx, db, sqlcgen.InsertUserProfileParams{
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
		if err := sqlcgen.Q.InsertUserProfileMutation(ctx, db, sqlcgen.InsertUserProfileMutationParams{
			UserID:    userID,
			Username:  payload.Username,
			Bio:       "",
			ImageUrl:  "",
			CreatedAt: now,
		}); err != nil {
			return errors.WithStack(err)
		}

		if err := sqlcgen.Q.InsertUserEmail(ctx, db, sqlcgen.InsertUserEmailParams{
			UserID:    userID,
			Email:     payload.Email,
			CreatedAt: now,
		}); err != nil {
			if myrdb.IsErrUniqueViolation(err) {
				return user.ErrEmailAlreadyUsed
			}
			return errors.WithStack(err)
		}
		if err := sqlcgen.Q.InsertUserEmailMutation(ctx, db, sqlcgen.InsertUserEmailMutationParams{
			UserID:    userID,
			Email:     payload.Email,
			CreatedAt: now,
		}); err != nil {
			return errors.WithStack(err)
		}

		if err := sqlcgen.Q.InsertUserAuthPassword(ctx, db, sqlcgen.InsertUserAuthPasswordParams{
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

func (s *User) GetCurrent(ctx context.Context) (*goa.GetCurrentResult, error) {
	userID, err := myctx.ShouldGetAuthenticatedUserID(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user, err := s.getUserByUserID(ctx, s.db, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &goa.GetCurrentResult{User: user}, nil
}

func (s *User) Update(ctx context.Context, payload *goa.UpdatePayload) (res *goa.UpdateResult, err error) {
	userID, err := myctx.ShouldGetAuthenticatedUserID(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	now := mytime.Now(ctx)
	db := s.db

	if err := myrdb.Tx(ctx, db, func(ctx context.Context, txdb myrdb.TxDB) error {
		db := txdb

		if payload.Email != nil {
			if err := sqlcgen.Q.UpdateUserEmail(ctx, db, sqlcgen.UpdateUserEmailParams{
				UserID:    userID,
				UpdatedAt: now,
				Email:     *payload.Email,
			}); err != nil {
				return errors.WithStack(err)
			}

			if err := sqlcgen.Q.InsertUserEmailMutation(ctx, db, sqlcgen.InsertUserEmailMutationParams{
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
			if err := sqlcgen.Q.UpdateUserAuthPasswordHash(ctx, db, sqlcgen.UpdateUserAuthPasswordHashParams{
				UserID:       userID,
				UpdatedAt:    now,
				PasswordHash: string(passwordHash),
			}); err != nil {
				return errors.WithStack(err)
			}
		}

		if payload.Username != nil || payload.Bio != nil || payload.Image != nil {
			current, err := sqlcgen.Q.GetUserProfileByUserID(ctx, db, userID)
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

			if err := sqlcgen.Q.UpdateUserProfile(ctx, db, param); err != nil {
				return errors.WithStack(err)
			}
			if err := sqlcgen.Q.InsertUserProfileMutation(ctx, db, sqlcgen.InsertUserProfileMutationParams{
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

	user, err := s.getUserByUserID(ctx, db, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &goa.UpdateResult{User: user}, nil
}

func (s *User) getUserByUserID(ctx context.Context, db myrdb.DB, userID uuid.UUID) (*goa.User, error) {
	email, err := sqlcgen.Q.GetUserEmailByUserID(ctx, db, userID)
	if err != nil {
		// handle ErrNoRows as internal server error
		return nil, errors.WithStack(err)
	}

	profile, err := sqlcgen.Q.GetUserProfileByUserID(ctx, db, userID)
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
