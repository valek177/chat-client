package user

import (
	"log"

	"github.com/spf13/cobra"
)

var CreateUserCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new chat user",
	Run: func(cmd *cobra.Command, args []string) {
		usernamesStr, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get usernames: %s\n", err.Error())
		}

		log.Printf("user %s created\n", usernamesStr)
	},
}

func init() {
	CreateUserCmd.Flags().StringP("username", "u", "", "User name")
	err := CreateUserCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalf("failed to mark username flag as required: %s\n", err.Error())
	}
}
