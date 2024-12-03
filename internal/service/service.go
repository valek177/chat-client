package service

import (
	"context"

	"github.com/valek177/chat-client/grpc/pkg/chat_v1"
)

// AuthService is interface for auth logic
type AuthService interface {
	GetAccessToken(ctx context.Context, username, password string) (string, error)
}

// ChatService is interface for chat logic
type ChatService interface {
	ConnectChat(ctx context.Context, chatname, username string) (
		chat_v1.ChatV1_ConnectChatClient, error,
	)
	CreateChat(ctx context.Context, chatname string, userIDs []int64) (int64, error)
	SendMessage(ctx context.Context, chatname, from, message string) error
}

// CommandService is interface for command logic
type CommandService interface {
	ConnectChat(ctx context.Context, chatname, username string) (chat_v1.ChatV1_ConnectChatClient, error)
	CreateChat(ctx context.Context, chatname string, userIDs []int64) (int64, error)
	Login(ctx context.Context, username, password string) (string, error)
	SendMessage(ctx context.Context, chatname, from, message string) error
}
