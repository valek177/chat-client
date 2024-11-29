package user

import (
	"log"

	"github.com/spf13/cobra"
)

var DeleteUserCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete chat user",
	Run: func(cmd *cobra.Command, args []string) {
		usernamesStr, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get usernames: %s\n", err.Error())
		}

		log.Printf("user %s deleted\n", usernamesStr)
	},
}

func init() {
	DeleteUserCmd.Flags().StringP("username", "u", "", "User name")
	err := DeleteUserCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalf("failed to mark username flag as required: %s\n", err.Error())
	}
}
