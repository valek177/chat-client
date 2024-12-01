package chat

import (
	"context"
)

// Login logins user to server
func (s *CommandService) Login(ctx context.Context, username, password string) error {
	accessToken, err := s.authService.GetAccessToken(ctx, username, password)
	if err != nil {
		return err
	}

	return saveTokenToFile(accessToken)
}
