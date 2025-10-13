package cmd

import (
	"github.com/chardoncs/downjack/cmd/gitignore"
	"github.com/chardoncs/downjack/cmd/license"
	"github.com/chardoncs/downjack/internal/cli"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "downjack [command]",
	Short: "Set up your gitignore and license files like using a lumberjack",
	Version: "0.1.0",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	// Use custom error display instead
	SilenceUsage: true,
	SilenceErrors: true,
}

func Execute() {
	err := rootCmd.Execute()
	cli.CheckErr(err)
}

func init() {
	rootCmd.AddCommand(gitignore.GitignoreCmd)
	rootCmd.AddCommand(license.LicenseCmd)
}
