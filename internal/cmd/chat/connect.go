package chat

import (
	"context"
	"io"
	"log"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/valek177/chat-client/client/internal/client"
	"github.com/valek177/chat-client/grpc/pkg/chat_v1"
)

var ConnectChatCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to chat",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get usernames: %s\n", err.Error())
		}

		chatID, err := cmd.Flags().GetInt64("chat-id")
		if err != nil {
			log.Fatalf("failed to get chat-id: %s\n", err.Error())
		}

		c, err := client.NewChatV1Client()
		if err != nil {
			log.Fatalf("unable to create client")
		}
		defer c.Close()

		wg := sync.WaitGroup{}
		wg.Add(1)

		go func() {
			defer wg.Done()

			err = connectChat(cmd.Context(), c.C, chatID, username, 5*time.Second)
			if err != nil {
				log.Fatalf("failed to connect chat: %v", err)
			}
		}()

		wg.Wait()
	},
}

func init() {
	ConnectChatCmd.Flags().StringP("username", "u", "", "User name")
	err := ConnectChatCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalf("failed to mark username flag as required: %s\n", err.Error())
	}

	ConnectChatCmd.Flags().Int64("chat-id", 0, "Chat ID")
	err = ConnectChatCmd.MarkFlagRequired("chat-id")
	if err != nil {
		log.Fatalf("failed to mark chat-id flag as required: %s\n", err.Error())
	}
}

func connectChat(ctx context.Context, client chat_v1.ChatV1Client, chatID int64, username string, period time.Duration) error {
	stream, err := client.ConnectChat(ctx, &chat_v1.ConnectChatRequest{
		ChatId:   chatID,
		Username: username,
	})
	if err != nil {
		return err
	}

	log.Println("Connected to chat", chatID)

	for {
		message, errRecv := stream.Recv()
		if errRecv == io.EOF {
			log.Println("error receive")
			return nil
		}
		if errRecv != nil {
			log.Println("failed to receive message from stream: ", errRecv)
			return nil
		}

		log.Printf("[%v] - [from: %s]: %s\n",
			color.YellowString(message.GetCreatedAt().AsTime().Format(time.RFC3339)),
			color.BlueString(message.GetFrom()),
			message.GetText(),
		)
	}
}
