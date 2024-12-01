package auth

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/valek177/chat-client/grpc/pkg/chat_v1"
)

// Client is interface for chat client
type Client interface {
	ConnectChat(ctx context.Context, chatname, username string) (
		chat_v1.ChatV1_ConnectChatClient, error)
	CreateChat(ctx context.Context, chatname string, userIDs []int64) (int64, error)
	SendMessage(ctx context.Context, chatname, from, message string) error
}

type client struct {
	client chat_v1.ChatV1Client
}

// NewClient creates new chat client
func NewClient(conn *grpc.ClientConn) *client {
	return &client{
		client: chat_v1.NewChatV1Client(conn),
	}
}

// ConnectChat connects user to chat
func (c *client) ConnectChat(ctx context.Context, chatname, username string) (
	chat_v1.ChatV1_ConnectChatClient, error,
) {
	resp, err := c.client.ConnectChat(ctx, &chat_v1.ConnectChatRequest{
		Chatname: chatname,
		Username: username,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateChat creates new chat
func (c *client) CreateChat(ctx context.Context, chatname string, userIDs []int64,
) (int64, error) {
	res, err := c.client.CreateChat(ctx, &chat_v1.CreateChatRequest{
		Name: chatname, UserIds: userIDs,
	})
	if err != nil {
		return 0, err
	}

	return res.GetId(), nil
}

// SendMessage sends message to chat
func (c *client) SendMessage(ctx context.Context, chatname, from, message string) error {
	_, err := c.client.SendMessage(ctx, &chat_v1.SendMessageRequest{
		Chatname: chatname,
		Message: &chat_v1.Message{
			From:      from,
			Text:      message,
			CreatedAt: timestamppb.Now(),
		},
	})
	if err != nil {
		return err
	}

	return nil
}
