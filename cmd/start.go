package cmd

import "github.com/spf13/cobra"

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		println("TODO: Start the server")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
