package cmd

import (
	"github.com/cjfinnell/go-server/server"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.NewService().Run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
