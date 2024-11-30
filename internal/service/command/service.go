package cmd

import "github.com/valek177/chat-client/internal/service"

type CommandService struct {
	chatService service.ChatService
	authService service.AuthService
}

func NewChatCommandService(
	chatService service.ChatService,
	authService service.AuthService,
) *CommandService {
	return &CommandService{
		chatService: chatService,
		authService: authService,
	}
}
