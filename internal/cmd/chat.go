package cmd

import (
	"github.com/spf13/cobra"

	"github.com/valek177/chat-client/internal/cmd/chat"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Command for chat management",
}

func init() {
	chatCmd.AddCommand(chat.ConnectChatCmd)
	chatCmd.AddCommand(chat.CreateChatCmd)
	chatCmd.AddCommand(chat.DeleteChatCmd)
	chatCmd.AddCommand(chat.DisconnectChatCmd)
}
