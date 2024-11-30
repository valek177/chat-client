package chat

import (
	"io"
	"log"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/valek177/chat-client/internal/app"
)

var ConnectChatCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to chat",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get usernames: %s\n", err.Error())
		}

		chatname, err := cmd.Flags().GetString("chatname")
		if err != nil {
			log.Fatalf("failed to get chat-id: %s\n", err.Error())
		}

		wg := sync.WaitGroup{}
		wg.Add(1)

		go func() {
			defer wg.Done()

			stream, err := app.ConnectChat(cmd.Context(), chatname, username)
			if err != nil {
				log.Fatalf("unable to connect to chat with name %s", chatname)
			}

			log.Println("Connected to chat", chatname)

			for {
				message, errRecv := stream.Recv()
				if errRecv == io.EOF {
					log.Println("error receive")
				}
				if errRecv != nil {
					log.Println("failed to receive message from stream: ", errRecv)
				}

				log.Printf("[%v] - [from: %s]: %s\n",
					color.YellowString(message.GetCreatedAt().AsTime().Format(time.RFC3339)),
					color.BlueString(message.GetFrom()),
					message.GetText(),
				)
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

	ConnectChatCmd.Flags().StringP("chatname", "c", "", "Chat name")
	err = ConnectChatCmd.MarkFlagRequired("chatname")
	if err != nil {
		log.Fatalf("failed to mark chatname flag as required: %s\n", err.Error())
	}
}
