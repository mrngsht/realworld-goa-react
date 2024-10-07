package service

import (
	"context"

	"github.com/mrngsht/realworld-goa-react/gen/user"
)

type User struct{}

var _ user.Service = User{}

func (u User) Login(ctx context.Context, payload *user.LoginPayload) (res *user.LoginResult, err error) {
	return &user.LoginResult{
		User: &user.UserType{
			Email:    payload.Email,
			Username: "taro",
		},
	}, nil
}

func (u User) Register(ctx context.Context, payload *user.RegisterPayload) (res *user.RegisterResult, err error) {
	return &user.RegisterResult{
		User: &user.UserType{
			Email:    payload.Email,
			Username: "taro",
		},
	}, nil
}
