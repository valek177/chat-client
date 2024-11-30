package chat

import (
	"context"
	"io"
	"log"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/valek177/chat-client/grpc/pkg/chat_v1"
	"github.com/valek177/chat-client/internal/client"
)

var ConnectChatCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to chat",
	Run: func(cmd *cobra.Command, args []string) {
		// application.commandService.ConnectChat
		// app.ConnectChat(ctx, chatID)
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get usernames: %s\n", err.Error())
		}

		chatname, err := cmd.Flags().GetString("chatname")
		if err != nil {
			log.Fatalf("failed to get chat-id: %s\n", err.Error())
		}

		// execute chatService logic

		c, err := client.NewChatV1Client()
		if err != nil {
			log.Fatalf("unable to create client")
		}
		defer c.Close()

		wg := sync.WaitGroup{}
		wg.Add(1)

		go func() {
			defer wg.Done()

			err = connectChat(cmd.Context(), c.C, chatname, username)
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

	ConnectChatCmd.Flags().StringP("chatname", "c", "", "Chat name")
	err = ConnectChatCmd.MarkFlagRequired("chatname")
	if err != nil {
		log.Fatalf("failed to mark chatname flag as required: %s\n", err.Error())
	}
}

func connectChat(ctx context.Context, c chat_v1.ChatV1Client, chatname string, username string) error {
	/// execute chat service connect
	stream, err := c.ConnectChat(ctx, &chat_v1.ConnectChatRequest{
		Chatname: chatname,
		Username: username,
	})
	if err != nil {
		return err
	}
	///

	log.Println("Connected to chat", chatname)

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
