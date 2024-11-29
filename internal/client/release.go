package client

import (
	"google.golang.org/grpc"

	"github.com/valek177/auth/grpc/pkg/auth_v1"
	"github.com/valek177/chat-client/grpc/pkg/chat_v1"
)

func chatClient(conn *grpc.ClientConn) chat_v1.ChatV1Client {
	return chat_v1.NewChatV1Client(conn)
}

func authClient(conn *grpc.ClientConn) auth_v1.AuthV1Client {
	return auth_v1.NewAuthV1Client(conn)
}
