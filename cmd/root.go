package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "Random Facts CLI app",
		Long: `Generates random facts`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}