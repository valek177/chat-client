package auth

import (
	"context"

	"google.golang.org/grpc"

	"github.com/valek177/auth/grpc/pkg/auth_v1"
)

// AuthClient is interface for auth client
type AuthClient interface {
	Login(ctx context.Context, username, password string) (string, string, error)
}

type authClient struct {
	client auth_v1.AuthV1Client
}

// NewAuthClient returns new AuthClient
func NewAuthClient(conn *grpc.ClientConn) *authClient {
	return &authClient{
		client: auth_v1.NewAuthV1Client(conn),
	}
}

// Login executes user login and returns tokens
func (c *authClient) Login(ctx context.Context, username, password string) (string, string, error) {
	resp, err := c.client.Login(ctx, &auth_v1.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return "", "", err
	}

	return resp.GetAccessToken(), resp.GetRefreshToken(), nil
}
