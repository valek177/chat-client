package auth

import (
	"context"

	"github.com/valek177/chat-client/internal/service"
)

type serv struct {
	authClient AuthClient
}

// NewService creates new service with settings
func NewService(authClient AuthClient) service.AuthService {
	return &serv{
		authClient: authClient,
	}
}

func (s *serv) GetAccessToken(ctx context.Context, username, password string) (string, error) {
	accessToken, _, err := s.authClient.Login(ctx, username, password)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
