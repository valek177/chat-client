package chat

import (
	"context"
)

func (s *CommandService) SendMessage(ctx context.Context, chatname, from, message string) error {
	// s.authService.Login(ctx, from, password)
	return s.chatService.SendMessage(ctx, chatname, from, message)
}
