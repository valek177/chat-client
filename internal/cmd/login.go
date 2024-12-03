package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/valek177/chat-client/internal/app"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login command",
	Run: func(cmd *cobra.Command, _ []string) {
		ctx := cmd.Context()
		application, err := app.NewApp(cmd.Context())
		if err != nil {
			log.Fatalf("failed to login: %v", err)
		}
		cmdService, err := application.ServiceProvider.CommandService(ctx)
		if err != nil {
			log.Fatalf("failed to login: %v", err)
		}

		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get flag username: %s\n", err.Error())
		}

		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalf("failed to get flag password: %s\n", err.Error())
		}

		err = cmdService.Login(ctx, username, password)
		if err != nil {
			log.Fatalf("failed to login: %v", err)
		}
		log.Printf("login successful")
	},
}

func init() {
	loginCmd.Flags().StringP("username", "u", "", "User name")
	err := loginCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalf("failed to mark username flag as required: %s\n", err.Error())
	}
	loginCmd.Flags().StringP("password", "p", "", "Password")
	err = loginCmd.MarkFlagRequired("password")
	if err != nil {
		log.Fatalf("failed to mark password flag as required: %s\n", err.Error())
	}
}
