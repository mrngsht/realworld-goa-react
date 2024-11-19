// Code generated by goa v3.19.1, DO NOT EDIT.
//
// article client
//
// Command:
// $ goa gen github.com/mrngsht/realworld-goa-react/design

package article

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "article" service client.
type Client struct {
	CreateEndpoint goa.Endpoint
}

// NewClient initializes a "article" service client given the endpoints.
func NewClient(create goa.Endpoint) *Client {
	return &Client{
		CreateEndpoint: create,
	}
}

// Create calls the "create" endpoint of the "article" service.
func (c *Client) Create(ctx context.Context, p *CreatePayload) (res *CreateResult, err error) {
	var ires any
	ires, err = c.CreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*CreateResult), nil
}
