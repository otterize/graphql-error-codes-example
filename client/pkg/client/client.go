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

func (c *Client) GetDog(ctx context.Context, name string, password string) (AllDogInfo, error) {
	res, err := getDog(ctx, c.gqlClient, name, password)
	return res.Dog.AllDogInfo, err
}

func (c *Client) AddDog(ctx context.Context, dogInput DogInput) error {
	_, err := addDog(ctx, c.gqlClient, dogInput)
	return err
}
