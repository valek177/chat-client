package env

import (
	"net"
	"os"

	"github.com/pkg/errors"

	"github.com/valek177/chat-client/internal/config"
)

var _ config.ClientConfig = (*clientConfig)(nil)

const (
	grpcAuthHostEnvName = "AUTH_SERVER_HOST"
	grpcAuthPortEnvName = "AUTH_SERVER_PORT"
	grpcChatHostEnvName = "CHAT_SERVER_HOST"
	grpcChatPortEnvName = "CHAT_SERVER_PORT"
	serviceTLSCertFile  = "GRPC_TLS_CERT_FILE"
	serviceTLSKeyFile   = "GRPC_TLS_KEY_FILE"
)

type clientConfig struct {
	authServerHost string
	authServerPort string
	chatServerHost string
	chatServerPort string
	tlsCertFile    string
	tlsKeyFile     string
}

// NewClientConfig creates new clientConfig
func NewClientConfig() (*clientConfig, error) {
	authServerHost := os.Getenv(grpcAuthHostEnvName)
	if len(authServerHost) == 0 {
		return nil, errors.New("grpc auth server host not found")
	}

	authServerPort := os.Getenv(grpcAuthPortEnvName)
	if len(authServerPort) == 0 {
		return nil, errors.New("grpc auth port not found")
	}

	chatServerHost := os.Getenv(grpcChatHostEnvName)
	if len(chatServerHost) == 0 {
		return nil, errors.New("grpc chat server host not found")
	}

	chatServerPort := os.Getenv(grpcChatPortEnvName)
	if len(chatServerPort) == 0 {
		return nil, errors.New("grpc chat port not found")
	}

	tlsServiceCertFile := os.Getenv(serviceTLSCertFile)
	if tlsServiceCertFile == "" {
		return nil, errors.New("grpc tls cert file not found")
	}

	tlsServiceKeyFile := os.Getenv(serviceTLSKeyFile)
	if tlsServiceKeyFile == "" {
		return nil, errors.New("grpc tls key file not found")
	}

	return &clientConfig{
		authServerHost: authServerHost,
		authServerPort: authServerPort,
		chatServerHost: chatServerHost,
		chatServerPort: chatServerPort,
		tlsCertFile:    tlsServiceCertFile,
		tlsKeyFile:     tlsServiceKeyFile,
	}, nil
}

// AuthServerAddress returns address of auth server from config
func (cfg *clientConfig) AuthServerAddress() string {
	return net.JoinHostPort(cfg.authServerHost, cfg.authServerPort)
}

// AuthServerAddress returns address of auth server from config
func (cfg *clientConfig) ChatServerAddress() string {
	return net.JoinHostPort(cfg.chatServerHost, cfg.chatServerPort)
}

// TLSCertFile returns path to TLS cert file from config
func (cfg *clientConfig) TLSCertFile() string {
	return cfg.tlsCertFile
}

// TLSKeyFile returns path to TLS key file from config
func (cfg *clientConfig) TLSKeyFile() string {
	return cfg.tlsKeyFile
}
