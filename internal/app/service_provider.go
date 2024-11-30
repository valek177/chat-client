package app

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/valek177/chat-client/internal/config"
	"github.com/valek177/chat-client/internal/config/env"
	"github.com/valek177/chat-client/internal/service/auth"
	auth "github.com/valek177/chat-client/internal/service/chat"
	cmd "github.com/valek177/chat-client/internal/service/command"
	"github.com/valek177/platform-common/pkg/closer"
)

type serviceProvider struct {
	config config.ClientConfig

	authClient auth.AuthClient
	chatClient chat.AuthClient

	authService service.AuthService
	chatService service.ChatService
	cmdService  service.CommandService

	authConn *grpc.ClientConn
	chatConn *grpc.ClientConn

	commandService *cmd.CommandService
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) ClientConfig() (config.ClientConfig, error) {
	if s.config == nil {
		cfg, err := env.NewClientConfig()
		if err != nil {
			return nil, err
		}

		s.config = cfg
	}

	return s.config, nil
}

func (s *serviceProvider) ChatConn() (*grpc.ClientConn, error) {
	if s.chatConn == nil {
		creds := insecure.NewCredentials()

		cfg, err := s.ClientConfig()
		if err != nil {
			return nil, err
		}

		conn, err := grpc.NewClient(
			cfg.ChatServerAddress(),
			grpc.WithTransportCredentials(creds),
		)
		if err != nil {
			return nil, err
		}

		closer.Add(conn.Close)

		s.chatConn = conn
	}

	return s.chatConn, nil
}

func (s *serviceProvider) ChatClient(ctx context.Context) (chat.ChatClient, error) {
	if s.chatClient == nil {
		chatConn, err := s.ChatConnection()
		if err != nil {
			return nil, err
		}
		s.chatClient = chat.NewChatClient(chatConn, chatService)
	}

	return s.chatClient, nil
}

func (s *serviceProvider) AuthClient(ctx context.Context) (auth.AuthClient, error) {
	if s.authClient == nil {
		authConn, err := s.AuthConnection()
		if err != nil {
			return nil, err
		}
		s.authClient = auth.NewAuthClient(authConn)
	}

	return s.authClient, nil
}

// AuthConnection returns AuthConnection
func (s *serviceProvider) AuthConnection() (*grpc.ClientConn, error) {
	if s.authConn == nil {
		var err error
		cfg, err := s.ClientConfig()
		if err != nil {
			return nil, err
		}
		creds, err := credentials.NewClientTLSFromFile(cfg.TLSCertFile(), "")
		if err != nil {
			return nil, err
		}
		conn, err := grpc.NewClient(
			cfg.AuthServerAddress(),
			grpc.WithTransportCredentials(creds),
		)
		if err != nil {
			return nil, err
		}

		closer.Add(conn.Close)

		s.authConn = conn
	}

	return s.authConn, nil
}

func (s *serviceProvider) AuthService(ctx context.Context) (service.AuthService, error) {
	if s.authService == nil {
		authClient, err := s.AuthClient(ctx)
		if err != nil {
			return nil, err
		}
		s.authService = auth.NewService(authClient)
	}

	return s.authService, nil
}

func (s *serviceProvider) CommandService(ctx context.Context) (*cmd.CommandService, error) {
	if s.commandService == nil {
		chatService, err := s.chatService()
		if err != nil {
			return nil, err
		}
		authService, err := s.authService()
		if err != nil {
			return nil, err
		}
		s.commandService = cmd.NewChatCommandService(chatService, authService)

	}

	return s.commandService, nil
}
