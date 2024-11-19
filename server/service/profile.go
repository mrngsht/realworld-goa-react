package service

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/mrngsht/realworld-goa-react/domain/user"
	goa "github.com/mrngsht/realworld-goa-react/gen/profile"
	"github.com/mrngsht/realworld-goa-react/myctx"
	"github.com/mrngsht/realworld-goa-react/myerr"
	"github.com/mrngsht/realworld-goa-react/myrdb"
	"github.com/mrngsht/realworld-goa-react/myrdb/sqlcgen"
	"github.com/mrngsht/realworld-goa-react/mytime"
)

type Profile struct {
	rdb myrdb.Conn
}

func NewProfile(rdb myrdb.Conn) *Profile {
	return &Profile{rdb: rdb}
}

var _ goa.Service = &Profile{}

func (s *Profile) FollowUser(ctx context.Context, payload *goa.FollowUserPayload) (res *goa.FollowUserResult, err error) {
	defer func() {
		if apErr, ok := myerr.AsAppErr(err); ok {
			switch apErr {
			case user.ErrUserNotFound:
				err = goa.MakeUserNotFound(err)
			case user.ErrUserAlreadyFollowing:
				err = goa.MakeUserAlreadyFollowing(err)
			}
		}
	}()

	q := sqlcgen.New(s.rdb)
	requestUserID := myctx.MustGetRequestUserID(ctx)

	followingProfile, err := q.GetUserProfileByUsername(ctx, payload.Username)
	if err != nil {
		if myrdb.IsErrNoRows(err) {
			return nil, user.ErrUserNotFound
		}
		return nil, errors.WithStack(err)
	}

	if err := myrdb.Tx(ctx, s.rdb, func(ctx context.Context, tx myrdb.TxConn) error {
		q = sqlcgen.New(tx)
		now := mytime.Now(ctx)

		if err := q.InsertUserFollow(ctx, sqlcgen.InsertUserFollowParams{
			CreatedAt:      now,
			UserID:         requestUserID,
			FollowedUserID: followingProfile.UserID,
		}); err != nil {
			if myrdb.IsErrUniqueViolation(err) {
				return user.ErrUserAlreadyFollowing
			}
			return errors.WithStack(err)
		}

		if err := q.InsertUserFollowMutation(ctx, sqlcgen.InsertUserFollowMutationParams{
			CreatedAt:      now,
			UserID:         requestUserID,
			FollowedUserID: followingProfile.UserID,
			Type:           sqlcgen.UserFollowMutationTypeFollow,
		}); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return &goa.FollowUserResult{
		Profile: &goa.Profile{
			Username:  payload.Username,
			Bio:       followingProfile.Bio,
			Image:     followingProfile.ImageUrl,
			Following: true,
		},
	}, nil
}

func (s *Profile) UnfollowUser(ctx context.Context, payload *goa.UnfollowUserPayload) (res *goa.UnfollowUserResult, err error) {
	defer func() {
		if apErr, ok := myerr.AsAppErr(err); ok {
			switch apErr {
			case user.ErrUserNotFound:
				err = goa.MakeUserNotFound(err)
			case user.ErrUserNotFollowing:
				err = goa.MakeUserNotFollowing(err)
			}
		}
	}()

	q := sqlcgen.New(s.rdb)
	requestUserID := myctx.MustGetRequestUserID(ctx)

	followingProfile, err := q.GetUserProfileByUsername(ctx, payload.Username)
	if err != nil {
		if myrdb.IsErrNoRows(err) {
			return nil, user.ErrUserNotFound
		}
		return nil, errors.WithStack(err)
	}

	if err := myrdb.Tx(ctx, s.rdb, func(ctx context.Context, tx myrdb.TxConn) error {
		q = sqlcgen.New(tx)
		now := mytime.Now(ctx)

		rowsAffected, err := q.DeleteUserFollow(ctx, sqlcgen.DeleteUserFollowParams{
			UserID:         requestUserID,
			FollowedUserID: followingProfile.UserID,
		})
		if err != nil {
			return errors.WithStack(err)
		}
		if rowsAffected == 0 {
			return user.ErrUserNotFollowing
		}

		if err := q.InsertUserFollowMutation(ctx, sqlcgen.InsertUserFollowMutationParams{
			CreatedAt:      now,
			UserID:         requestUserID,
			FollowedUserID: followingProfile.UserID,
			Type:           sqlcgen.UserFollowMutationTypeUnfollow,
		}); err != nil {
			return errors.WithStack(err)
		}

		return nil
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return &goa.UnfollowUserResult{
		Profile: &goa.Profile{
			Username:  payload.Username,
			Bio:       followingProfile.Bio,
			Image:     followingProfile.ImageUrl,
			Following: false,
		},
	}, nil
}
