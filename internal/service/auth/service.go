package auth

import (
	"context"

	"github.com/valek177/chat-client/internal/service"
)

type serv struct {
	authClient Client
}

// NewService creates new service with settings
func NewService(authClient Client) service.AuthService {
	return &serv{
		authClient: authClient,
	}
}

// GetAccessToken returns access token
func (s *serv) GetAccessToken(ctx context.Context, username, password string) (string, error) {
	accessToken, _, err := s.authClient.Login(ctx, username, password)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
