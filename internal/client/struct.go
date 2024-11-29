package client

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/valek177/auth/grpc/pkg/auth_v1"
	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

const (
	chatAddress = "localhost:50071"
	authAddress = "localhost:50061"
)

type ChatV1Client struct {
	conn *grpc.ClientConn
	C    chat_v1.ChatV1Client
}

type AuthV1Client struct {
	conn *grpc.ClientConn
	C    auth_v1.AuthV1Client
}

func NewChatV1Client() (*ChatV1Client, error) {
	// creds, err := credentials.NewClientTLSFromFile("/home/valek/microservices_course/auth/tls/service.pem", "")
	// if err != nil {
	// 	return nil, err
	// }
	conn, err := grpc.NewClient(
		chatAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &ChatV1Client{
		conn: conn,
		C:    chatClient(conn),
	}, nil
}

func (c *ChatV1Client) Close() {
	c.conn.Close()
}

func NewAuthV1Client() (*AuthV1Client, error) {
	creds, err := credentials.NewClientTLSFromFile("/home/valek/microservices_course/auth/tls/service.pem", "")
	if err != nil {
		return nil, err
	}
	conn, err := grpc.NewClient(
		authAddress,
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &AuthV1Client{
		conn: conn,
		C:    authClient(conn),
	}, nil
}

func (c *AuthV1Client) Close() {
	c.conn.Close()
}
