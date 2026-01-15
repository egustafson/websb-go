package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:               "websb <sub-command>",
	PersistentPreRunE: initAppHook,
}

var (
	// GitSummary = git describe --always --dirty --tags
	GitSummary string
	// BuildDate = date -u +%Y-%m-%dT%H:%M:%SZ
	BuildDate string
)

// flags initialized in flags.go

func Execute(gitSummary, buildDate string) error {
	GitSummary = gitSummary
	BuildDate = buildDate
	return rootCmd.Execute()
}

func initAppHook(_ *cobra.Command, _ []string) error {

	// TODO: app initialization here

	return nil
}
