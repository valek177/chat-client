package chat

import (
	"context"

	"github.com/valek177/chat-client/grpc/pkg/chat_v1"
)

// ConnectChat connects to chat
func (s *CommandService) ConnectChat(ctx context.Context, chatname, username string) (
	chat_v1.ChatV1_ConnectChatClient, error,
) {
	ctx, err := TokenCtx(ctx)
	if err != nil {
		return nil, err
	}
	return s.chatService.ConnectChat(ctx, chatname, username)
}
