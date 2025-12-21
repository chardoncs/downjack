package cmd

import (
	"go.chardoncs.dev/downjack/internal/cmd/gitignore"
	"go.chardoncs.dev/downjack/internal/cmd/license"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "downjack",
	Short: "Set up your gitignore and license files like using a lumberjack",
	// Use custom error display instead
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	RootCmd.AddCommand(gitignore.GitignoreCmd)
	RootCmd.AddCommand(license.LicenseCmd)
}
