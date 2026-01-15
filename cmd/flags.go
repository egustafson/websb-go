package cmd

import "github.com/egustafson/websb-go/pkg/config"

var (
	flags = config.Flags{
		Verbose: false,
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&flags.Verbose, "verbose", "v", false,
		"verbose output")
}
