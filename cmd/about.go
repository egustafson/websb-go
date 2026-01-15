package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Display build information about this binary",
	RunE:  doAbout,
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}

func doAbout(_ *cobra.Command, _ []string) error {
	fmt.Println("---")
	fmt.Printf("version: %s\n", GitSummary)
	fmt.Printf("build-date: %s\n", BuildDate)
	if flags.Verbose {
		fmt.Println("verbose: true")
	}
	fmt.Println("...")
	return nil
}
