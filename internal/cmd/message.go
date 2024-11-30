package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/valek177/auth/grpc/pkg/auth_v1"
	"github.com/valek177/chat-client/grpc/pkg/chat_v1"
	"github.com/valek177/chat-client/internal/client"
)

var sendMessageCmd = &cobra.Command{
	Use:   "send-message",
	Short: "Send message",
	Run: func(cmd *cobra.Command, args []string) {
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

		c, err := client.NewChatV1Client()
		if err != nil {
			log.Fatalf("unable to create client")
		}
		defer c.Close()

		authClient, err := client.NewAuthV1Client()
		if err != nil {
			log.Fatalf("unable to create auth client")
		}
		defer c.Close()

		// login to auth service; get token and use it for send message
		accessToken, err := login(cmd.Context(), authClient.C)
		if err != nil {
			log.Fatalf("failed to login: %v", err)
		}

		log.Println("token is ", accessToken)

		md := metadata.New(map[string]string{"authorization": "Bearer " + accessToken})
		newCtx := metadata.NewOutgoingContext(cmd.Context(), md)

		err = sendMessage(newCtx, c.C, from, chatname, message)
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

func sendMessage(ctx context.Context, c chat_v1.ChatV1Client, from string, chatname string, message string) error {
	_, err := c.SendMessage(ctx, &chat_v1.SendMessageRequest{
		Chatname: chatname,
		Message: &chat_v1.Message{
			From:      from,
			Text:      message,
			CreatedAt: timestamppb.Now(),
		},
	})
	if err != nil {
		log.Println("failed to send message: ", err)
		return err
	}

	return nil
}

func login(ctx context.Context, authClient auth_v1.AuthV1Client) (string, error) {
	resp, err := authClient.Login(ctx, &auth_v1.LoginRequest{
		Username: "valya",
		Password: "P@ssw0rd123",
	})
	if err != nil {
		return "", err
	}

	log.Println("token is ", resp.GetAccessToken())

	return resp.GetAccessToken(), nil
}
