package cmd

import (
	"github.com/spf13/cobra"

	"github.com/egustafson/websb-go/pkg/server"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run the web server daemon",
	RunE:  doDaemon,
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}

func doDaemon(cmd *cobra.Command, args []string) error {
	return server.Start(flags)
}
