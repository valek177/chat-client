package chat

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/valek177/chat-client/client/internal/client"
	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

var CreateChatCmd = &cobra.Command{
	Use:   "create",
	Short: "Create chat",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewChatV1Client()
		if err != nil {
			log.Fatalf("unable to create client for connect")
		}
		defer c.Close()

		chatID, err := createChat(cmd.Context(), c.C)
		if err != nil {
			log.Fatalf("failed to connect chat: %v", err)
		}

		log.Printf("was created chat with id %d", chatID)
	},
}

// input ids, name of chat
func init() {
}

func createChat(ctx context.Context, client chat_v1.ChatV1Client) (int64, error) {
	res, err := client.CreateChat(ctx, &chat_v1.CreateChatRequest{
		Name: "newchat", UserIds: []int64{1, 2, 3, 4},
	})
	if err != nil {
		return 0, err
	}

	return res.GetId(), nil
}
