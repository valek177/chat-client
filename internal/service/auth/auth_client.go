package auth

import (
	"context"

	"google.golang.org/grpc"

	"github.com/valek177/auth/grpc/pkg/auth_v1"
)

// Client is interface for auth client
type Client interface {
	Login(ctx context.Context, username, password string) (string, string, error)
}

type client struct {
	client auth_v1.AuthV1Client
}

// NewClient returns new AuthClient
func NewClient(conn *grpc.ClientConn) *client {
	return &client{
		client: auth_v1.NewAuthV1Client(conn),
	}
}

// Login executes user login and returns tokens
func (c *client) Login(ctx context.Context, username, password string) (string, string, error) {
	resp, err := c.client.Login(ctx, &auth_v1.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return "", "", err
	}

	return resp.GetAccessToken(), resp.GetRefreshToken(), nil
}
