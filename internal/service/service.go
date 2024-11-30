package service

import (
	"context"

	"github.com/valek177/chat-client/grpc/pkg/chat_v1"
)

type AuthService interface {
	Login(ctx context.Context, username, password string) (string, error)
	// CreateUser()
	// DeleteUser()
}

type ChatService interface {
	ConnectChat(ctx context.Context, chatname, username string) (
		chat_v1.ChatV1_ConnectChatClient, error,
	)
	CreateChat(ctx context.Context, chatname string, userIDs []int64) (int64, error)
	DeleteChat(ctx context.Context, chatID int64) error
	SendMessage(ctx context.Context, chatname, from, message string) error
}

type CommandService interface {
	// CreateUser()
	// DeleteUser()
	ConnectChat(ctx context.Context, chatname, username string) (chat_v1.ChatV1_ConnectChatClient, error)
	// CreateChat(ctx context.Context, chatname string, userIDs []int64) (int64, error)
	// DeleteChat(ctx context.Context, chatID int64) error
}
