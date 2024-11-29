package chat

import "github.com/spf13/cobra"

var DisconnectChatCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect from chat",
	Run: func(cmd *cobra.Command, args []string) {
		// usernamesStr, err := cmd.Flags().GetString("username")
		// if err != nil {
		// 	log.Fatalf("failed to get usernames: %s\n", err.Error())
		// }

		// log.Printf("user %s created\n", usernamesStr)
	},
}
