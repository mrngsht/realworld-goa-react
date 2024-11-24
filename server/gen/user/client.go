// Code generated by goa v3.19.1, DO NOT EDIT.
//
// user client
//
// Command:
// $ goa gen github.com/mrngsht/realworld-goa-react/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "user" service client.
type Client struct {
	LoginEndpoint      goa.Endpoint
	RegisterEndpoint   goa.Endpoint
	GetCurrentEndpoint goa.Endpoint
	UpdateEndpoint     goa.Endpoint
}

// NewClient initializes a "user" service client given the endpoints.
func NewClient(login, register, getCurrent, update goa.Endpoint) *Client {
	return &Client{
		LoginEndpoint:      login,
		RegisterEndpoint:   register,
		GetCurrentEndpoint: getCurrent,
		UpdateEndpoint:     update,
	}
}

// Login calls the "login" endpoint of the "user" service.
// Login may return the following errors:
//   - "UserLoginBadRequest" (type *UserLoginBadRequest)
//   - error: internal error
func (c *Client) Login(ctx context.Context, p *LoginPayload) (res *LoginResult, err error) {
	var ires any
	ires, err = c.LoginEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*LoginResult), nil
}

// Register calls the "register" endpoint of the "user" service.
// Register may return the following errors:
//   - "UserRegisterBadRequest" (type *UserRegisterBadRequest)
//   - error: internal error
func (c *Client) Register(ctx context.Context, p *RegisterPayload) (res *RegisterResult, err error) {
	var ires any
	ires, err = c.RegisterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*RegisterResult), nil
}

// GetCurrent calls the "getCurrent" endpoint of the "user" service.
func (c *Client) GetCurrent(ctx context.Context) (res *GetCurrentResult, err error) {
	var ires any
	ires, err = c.GetCurrentEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*GetCurrentResult), nil
}

// Update calls the "update" endpoint of the "user" service.
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (res *UpdateResult, err error) {
	var ires any
	ires, err = c.UpdateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UpdateResult), nil
}
