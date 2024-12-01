package chat

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/valek177/chat-client/internal/app"
)

// CreateChatCmd creates chat
var CreateChatCmd = &cobra.Command{
	Use:   "create",
	Short: "Create chat",
	Run: func(cmd *cobra.Command, _ []string) {
		ctx := cmd.Context()
		application, err := app.NewApp(cmd.Context())
		if err != nil {
			log.Fatalf("failed to create chat: %v", err)
		}
		cmdService, err := application.ServiceProvider.CommandService(ctx)
		if err != nil {
			log.Fatalf("failed to create chat: %v", err)
		}

		chatname, err := cmd.Flags().GetString("chatname")
		if err != nil {
			log.Fatalf("failed to create chat: %v", err)
		}
		userIDs, err := cmd.Flags().GetInt64Slice("user-ids")
		if err != nil {
			log.Fatalf("failed to create chat: %v", err)
		}

		chatID, err := cmdService.CreateChat(ctx, chatname, userIDs)
		if err != nil {
			log.Fatalf("failed to create chat: %v", err)
		}

		log.Printf("was created chat with id %d", chatID)
	},
}

func init() {
	CreateChatCmd.Flags().StringP("chatname", "c", "", "Chat name")
	err := CreateChatCmd.MarkFlagRequired("chatname")
	if err != nil {
		log.Fatalf("failed to mark chatname flag as required: %s\n", err.Error())
	}

	CreateChatCmd.Flags().Int64SliceP("user-ids", "u", []int64{}, "User IDs list")
	err = CreateChatCmd.MarkFlagRequired("user-ids")
	if err != nil {
		log.Fatalf("failed to mark user-ids flag as required: %s\n", err.Error())
	}
}
