package app

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/valek177/chat-client/internal/config"
	"github.com/valek177/chat-client/internal/config/env"
	service "github.com/valek177/chat-client/internal/service"
	"github.com/valek177/chat-client/internal/service/auth"
	chat "github.com/valek177/chat-client/internal/service/chat"
	cmd "github.com/valek177/chat-client/internal/service/command"
	"github.com/valek177/platform-common/pkg/closer"
)

// ServiceProvider is service provider struct
type ServiceProvider struct {
	config config.ClientConfig

	authClient auth.Client
	chatClient chat.Client

	authService service.AuthService
	chatService service.ChatService

	authConn *grpc.ClientConn
	chatConn *grpc.ClientConn

	commandService *cmd.CommandService
}

// NewServiceProvider returns new service provider
func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

// ClientConfig returns client config object
func (s *ServiceProvider) ClientConfig() (config.ClientConfig, error) {
	if s.config == nil {
		cfg, err := env.NewClientConfig()
		if err != nil {
			return nil, err
		}

		s.config = cfg
	}

	return s.config, nil
}

// ChatConn returns chat connection
func (s *ServiceProvider) ChatConn() (*grpc.ClientConn, error) {
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

// ChatClient returns chat client
func (s *ServiceProvider) ChatClient(_ context.Context) (chat.Client, error) {
	if s.chatClient == nil {
		chatConn, err := s.ChatConnection()
		if err != nil {
			return nil, err
		}
		s.chatClient = chat.NewClient(chatConn)
	}

	return s.chatClient, nil
}

// AuthClient returns auth client
func (s *ServiceProvider) AuthClient(_ context.Context) (auth.Client, error) {
	if s.authClient == nil {
		authConn, err := s.AuthConnection()
		if err != nil {
			return nil, err
		}
		s.authClient = auth.NewClient(authConn)
	}

	return s.authClient, nil
}

// AuthConnection returns AuthConnection
func (s *ServiceProvider) AuthConnection() (*grpc.ClientConn, error) {
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

// ChatConnection returns connection to chat server
func (s *ServiceProvider) ChatConnection() (*grpc.ClientConn, error) {
	if s.chatConn == nil {
		var err error
		cfg, err := s.ClientConfig()
		if err != nil {
			return nil, err
		}
		conn, err := grpc.NewClient(
			cfg.ChatServerAddress(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return nil, err
		}

		closer.Add(conn.Close)

		s.chatConn = conn
	}

	return s.chatConn, nil
}

// AuthService returns auth service
func (s *ServiceProvider) AuthService(ctx context.Context) (service.AuthService, error) {
	if s.authService == nil {
		authClient, err := s.AuthClient(ctx)
		if err != nil {
			return nil, err
		}
		s.authService = auth.NewService(authClient)
	}

	return s.authService, nil
}

// ChatService returns chat service
func (s *ServiceProvider) ChatService(ctx context.Context) (service.ChatService, error) {
	if s.chatService == nil {
		chatClient, err := s.ChatClient(ctx)
		if err != nil {
			return nil, err
		}
		s.chatService = chat.NewService(chatClient)
	}

	return s.chatService, nil
}

// CommandService returns command service
func (s *ServiceProvider) CommandService(ctx context.Context) (*cmd.CommandService, error) {
	if s.commandService == nil {
		chatService, err := s.ChatService(ctx)
		if err != nil {
			return nil, err
		}
		authService, err := s.AuthService(ctx)
		if err != nil {
			return nil, err
		}
		s.commandService = cmd.NewChatCommandService(chatService, authService)

	}

	return s.commandService, nil
}
