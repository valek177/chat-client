package service

import (
	"context"

	"github.com/valek177/chat-client/grpc/pkg/chat_v1"
)

type ChatClient interface {
	CreateChat(ctx context.Context, chatname string, userIDs []int64)
}

type AuthService interface {
	GetAccessToken(ctx context.Context, username, password string) (string, error)
}

type ChatService interface {
	ConnectChat(ctx context.Context, chatname, username string) (
		chat_v1.ChatV1_ConnectChatClient, error,
	)
	CreateChat(ctx context.Context, chatname string, userIDs []int64) (int64, error)
	DeleteChat(ctx context.Context, chatID int64) error
}
