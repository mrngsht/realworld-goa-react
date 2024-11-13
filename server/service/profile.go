package service

import (
	"context"

	goa "github.com/mrngsht/realworld-goa-react/gen/profile"
	"github.com/mrngsht/realworld-goa-react/myrdb"
)

type Profile struct {
	rdb myrdb.RDB
}

func NewProfile(rdb myrdb.RDB) *Profile {
	return &Profile{rdb: rdb}
}

var _ goa.Service = &Profile{}

func (s *Profile) FollowUser(context.Context, *goa.FollowUserPayload) (res *goa.FollowUserResult, err error) {
	return nil, nil
}

func (s *Profile) UnfollowUser(context.Context, *goa.UnfollowUserPayload) (res *goa.UnfollowUserResult, err error) {
	return nil, nil
}
