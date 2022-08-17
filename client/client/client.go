package client

import (
	"context"
	"github.com/Khan/genqlient/graphql"
	"net/http"
)

type Client struct {
	serverAddress string
	gqlClient     graphql.Client
}

func NewClient(serverAddress string) *Client {
	return &Client{
		serverAddress: serverAddress,
		gqlClient:     graphql.NewClient(serverAddress, http.DefaultClient),
	}
}

func (c *Client) ErrorTypesDemo(ctx context.Context, firstName string, lastName string) (string, error) {
	res, err := errorTypesDemo(ctx, c.gqlClient, firstName, lastName)
	return res.ErrorTypesDemo, err
}
