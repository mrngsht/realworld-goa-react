package service

import (
	"context"
	"database/sql"
	"log"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"

	"github.com/mrngsht/realworld-goa-react/ctxtime"
	"github.com/mrngsht/realworld-goa-react/domain/user"
	goa "github.com/mrngsht/realworld-goa-react/gen/user"
	"github.com/mrngsht/realworld-goa-react/rdb"
	"github.com/mrngsht/realworld-goa-react/rdb/sqlcgen"
)

type User struct {
	rdb *sql.DB
}

func NewUser(rdb *sql.DB) User {
	return User{rdb: rdb}
}

var _ goa.Service = User{}

func (u User) Login(ctx context.Context, payload *goa.LoginPayload) (res *goa.LoginResult, err error) {
	return &goa.LoginResult{
		User: &goa.UserType{
			Email:    payload.Email,
			Username: "taro",
		},
	}, nil
}

func (u User) Register(ctx context.Context, payload *goa.RegisterPayload) (res *goa.RegisterResult, err error) {
	defer func() {
		// FIXME:
		if err != nil {
			log.Default().Printf("err: %+v\n", err)
		}
	}()

	q := sqlcgen.New(u.rdb)

	passwordHash, err := user.GenPasswordHash([]byte(payload.Password))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := rdb.Tx(ctx, u.rdb, func(ctx context.Context, tx *sql.Tx) error {
		q = q.WithTx(tx)

		now := ctxtime.Now(ctx)
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
