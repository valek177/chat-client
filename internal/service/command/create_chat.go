package chat

import (
	"context"
)

// CreateChat creates chat
func (s *CommandService) CreateChat(ctx context.Context, chatname string, userIDs []int64,
) (int64, error) {
	ctx, err := TokenCtx(ctx)
	if err != nil {
		return 0, err
	}
	return s.chatService.CreateChat(ctx, chatname, userIDs)
}
