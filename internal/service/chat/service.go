package auth

import (
	"context"

	"github.com/valek177/chat-client/grpc/pkg/chat_v1"
	"github.com/valek177/chat-client/internal/service"
)

type serv struct {
	chatClient Client
}

// NewService creates new service with settings
func NewService(chatClient Client) service.ChatService {
	return &serv{
		chatClient: chatClient,
	}
}

// ConnectChat connects user to chat
func (s *serv) ConnectChat(ctx context.Context, chatname, username string) (
	chat_v1.ChatV1_ConnectChatClient, error,
) {
	stream, err := s.chatClient.ConnectChat(ctx, chatname, username)
	if err != nil {
		return nil, err
	}

	return stream, nil
}

// CreateChat creates chat
func (s *serv) CreateChat(ctx context.Context, chatname string, userIDs []int64) (int64, error) {
	id, err := s.chatClient.CreateChat(ctx, chatname, userIDs)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// SendMessage sends message
func (s *serv) SendMessage(ctx context.Context, chatname, from, message string) error {
	err := s.chatClient.SendMessage(ctx, chatname, from, message)
	if err != nil {
		return err
	}

	return nil
}
