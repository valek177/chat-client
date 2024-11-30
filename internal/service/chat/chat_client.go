package auth

import (
	"context"

	"google.golang.org/grpc"

	"github.com/valek177/chat-client/grpc/pkg/chat_v1"
)

// AuthClient is interface for auth client
type ChatClient interface {
	ConnectChat(ctx context.Context, chatname, username string) (
		chat_v1.ChatV1_ConnectChatClient, error)
	CreateChat(ctx context.Context, chatname string, userIDs []int64) (int64, error)
	DeleteChat(ctx context.Context, chatID int64) error
	// Disconnect?
	SendMessage() // TODO
}

type chatClient struct {
	client chat_v1.ChatV1Client
}

func NewChatClient(conn *grpc.ClientConn) *chatClient {
	return &chatClient{
		client: chat_v1.NewChatV1Client(conn),
	}
}

func (c *chatClient) ConnectChat(ctx context.Context, chatname, username string) (
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

func (c *chatClient) CreateChat(ctx context.Context, chatname string, userIDs []int64,
) (int64, error) {
	res, err := c.client.CreateChat(ctx, &chat_v1.CreateChatRequest{
		Name: chatname, UserIds: userIDs,
	})
	if err != nil {
		return 0, err
	}

	return res.GetId(), nil
}

func (c *chatClient) DeleteChat(ctx context.Context, chatID int64) error {
	_, err := c.client.DeleteChat(ctx, &chat_v1.DeleteChatRequest{
		Id: chatID,
	})
	if err != nil {
		return err
	}

	return nil
}
