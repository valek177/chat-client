package chat

import (
	"context"
)

// SendMessage sends message
func (s *CommandService) SendMessage(ctx context.Context, chatname, from, message string) error {
	ctxWithToken, err := TokenCtx(ctx)
	if err != nil {
		return err
	}
	return s.chatService.SendMessage(ctxWithToken, chatname, from, message)
}
