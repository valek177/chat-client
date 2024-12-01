package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/valek177/chat-client/internal/app"
)

var sendMessageCmd = &cobra.Command{
	Use:   "send-message",
	Short: "Send message",
	Run: func(cmd *cobra.Command, _ []string) {
		ctx := cmd.Context()
		application, err := app.NewApp(cmd.Context())
		if err != nil {
			log.Fatalf("failed to send message to chat: %v", err)
		}
		cmdService, err := application.ServiceProvider.CommandService(ctx)
		if err != nil {
			log.Fatalf("failed to send message to chat: %v", err)
		}

		chatname, err := cmd.Flags().GetString("chatname")
		if err != nil {
			log.Fatalf("failed to get flag chatname: %s\n", err.Error())
		}

		message, err := cmd.Flags().GetString("message")
		if err != nil {
			log.Fatalf("failed to get flag message: %s\n", err.Error())
		}

		from, err := cmd.Flags().GetString("from")
		if err != nil {
			log.Fatalf("failed to get flag from: %s\n", err.Error())
		}

		err = cmdService.SendMessage(ctx, chatname, from, message)
		if err != nil {
			log.Fatalf("failed to send message: %v", err)
		}
		log.Printf("was sended message")
	},
}

func init() {
	sendMessageCmd.Flags().StringP("chatname", "c", "", "Chat name")
	err := sendMessageCmd.MarkFlagRequired("chatname")
	if err != nil {
		log.Fatalf("failed to mark chatname flag as required: %s\n", err.Error())
	}
	sendMessageCmd.Flags().StringP("message", "m", "", "Message text")
	err = sendMessageCmd.MarkFlagRequired("message")
	if err != nil {
		log.Fatalf("failed to mark message flag as required: %s\n", err.Error())
	}
	sendMessageCmd.Flags().StringP("from", "f", "", "From")
	err = sendMessageCmd.MarkFlagRequired("from")
	if err != nil {
		log.Fatalf("failed to mark from flag as required: %s\n", err.Error())
	}
}
