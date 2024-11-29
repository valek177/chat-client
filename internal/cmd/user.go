package cmd

import (
	"github.com/spf13/cobra"

	"github.com/valek177/chat-client/client/internal/cmd/user"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Command for user management",
}

func init() {
	userCmd.AddCommand(user.CreateUserCmd)
	userCmd.AddCommand(user.DeleteUserCmd)
}
