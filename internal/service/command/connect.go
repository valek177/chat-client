package chat

import (
	"context"

	"github.com/valek177/chat-client/grpc/pkg/chat_v1"
)

func (s *CommandService) ConnectChat(ctx context.Context, chatname, username string) (chat_v1.ChatV1_ConnectChatClient, error) {
	return s.chatService.ConnectChat(ctx, chatname, username)
}
