package auth

import (
	"context"

	"github.com/valek177/chat-client/grpc/pkg/chat_v1"
	"github.com/valek177/chat-client/internal/service"
)

type serv struct {
	chatClient ChatClient
}

// NewService creates new service with settings
func NewService(chatClient ChatClient) service.ChatService {
	return &serv{
		chatClient: chatClient,
	}
}

func (s *serv) ConnectChat(ctx context.Context, chatname, username string) (
	chat_v1.ChatV1_ConnectChatClient, error,
) {
	stream, err := s.chatClient.ConnectChat(ctx, chatname, username)
	if err != nil {
		return nil, err
	}

	return stream, nil
}

func (s *serv) CreateChat(ctx context.Context, chatname string, userIDs []int64) (int64, error) {
	id, err := s.chatClient.CreateChat(ctx, chatname, userIDs)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *serv) DeleteChat(ctx context.Context, chatID int64) error {
	return s.chatClient.DeleteChat(ctx, chatID)
}

func (s *serv) SendMessage(ctx context.Context, chatname, from, message string) error {
	err := s.chatClient.SendMessage(ctx, chatname, from, message)
	if err != nil {
		return err
	}

	return nil
}
