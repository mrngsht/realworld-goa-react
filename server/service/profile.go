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
	db myrdb.DB
}

func NewProfile(rdb myrdb.DB) *Profile {
	return &Profile{db: rdb}
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

	db := s.db
	requestUserID := myctx.MustGetRequestUserID(ctx)

	followingProfile, err := sqlcgen.Q.GetUserProfileByUsername(ctx, db, payload.Username)
	if err != nil {
		if myrdb.IsErrNoRows(err) {
			return nil, user.ErrUserNotFound
		}
		return nil, errors.WithStack(err)
	}

	if err := myrdb.Tx(ctx, db, func(ctx context.Context, txdb myrdb.TxDB) error {
		db := txdb
		now := mytime.Now(ctx)

		if err := sqlcgen.Q.InsertUserFollow(ctx, db, sqlcgen.InsertUserFollowParams{
			CreatedAt:      now,
			UserID:         requestUserID,
			FollowedUserID: followingProfile.UserID,
		}); err != nil {
			if myrdb.IsErrUniqueViolation(err) {
				return user.ErrUserAlreadyFollowing
			}
			return errors.WithStack(err)
		}

		if err := sqlcgen.Q.InsertUserFollowMutation(ctx, db, sqlcgen.InsertUserFollowMutationParams{
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

	db := s.db
	requestUserID := myctx.MustGetRequestUserID(ctx)

	followingProfile, err := sqlcgen.Q.GetUserProfileByUsername(ctx, db, payload.Username)
	if err != nil {
		if myrdb.IsErrNoRows(err) {
			return nil, user.ErrUserNotFound
		}
		return nil, errors.WithStack(err)
	}

	if err := myrdb.Tx(ctx, db, func(ctx context.Context, txdb myrdb.TxDB) error {
		db := txdb
		now := mytime.Now(ctx)

		rowsAffected, err := sqlcgen.Q.DeleteUserFollow(ctx, db, sqlcgen.DeleteUserFollowParams{
			UserID:         requestUserID,
			FollowedUserID: followingProfile.UserID,
		})
		if err != nil {
			return errors.WithStack(err)
		}
		if rowsAffected == 0 {
			return user.ErrUserNotFollowing
		}

		if err := sqlcgen.Q.InsertUserFollowMutation(ctx, db, sqlcgen.InsertUserFollowMutationParams{
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
