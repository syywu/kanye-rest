package cmd

import (
	"github.com/spf13/cobra"
)

var (

	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "Random Kayne's quotes CLI app",
		Long: `Generates random quotes`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}