package chat

import (
	"bufio"
	"context"
	"os"

	"google.golang.org/grpc/metadata"

	"github.com/valek177/chat-client/internal/service"
)

const (
	authHeaderName = "authorization"
	authPrefix     = "Bearer "

	accessTokenFilename = ".access_token"
)

// CommandService is struct for command logic
type CommandService struct {
	chatService service.ChatService
	authService service.AuthService
}

// NewChatCommandService creates new chat command service
func NewChatCommandService(
	chatService service.ChatService,
	authService service.AuthService,
) *CommandService {
	return &CommandService{
		chatService: chatService,
		authService: authService,
	}
}

// TokenCtx returns new context with access token
func TokenCtx(ctx context.Context) (context.Context, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	accessToken, err := os.ReadFile(homedir + "/" + accessTokenFilename) //nolint:gosec
	if err != nil {
		return nil, err
	}

	md := metadata.New(map[string]string{authHeaderName: authPrefix + string(accessToken)})

	return metadata.NewOutgoingContext(ctx, md), nil
}

func saveTokenToFile(accessToken string) error {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	f, err := os.Create(dirname + "/" + accessTokenFilename) //nolint:gosec
	if err != nil {
		return err
	}
	defer f.Close() //nolint:errcheck

	w := bufio.NewWriter(f)
	_, err = w.WriteString(accessToken)
	if err != nil {
		return err
	}
	w.Flush() //nolint:errcheck

	return nil
}
