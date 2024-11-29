package chat

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/valek177/chat-server/client/internal/client"
)

var DeleteChatCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete chat",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewChatV1Client()
		if err != nil {
			log.Fatalf("unable to create client for connect")
		}
		defer c.Close()

		// chatID, err := deleteChat(cmd.Context(), c.C)
		// if err != nil {
		// 	log.Fatalf("failed to delete chat: %v", err)
		// }

		// log.Printf("chat with id %d was deleted", chatID)
	},
}

func init() {
	DeleteChatCmd.Flags().Int64("chat-id", 0, "Chat ID")
}
