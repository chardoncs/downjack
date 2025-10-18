package cmd

import (
	"github.com/chardoncs/downjack/internal/cmd/gitignore"
	"github.com/chardoncs/downjack/internal/cmd/license"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "downjack",
	Short:   "Set up your gitignore and license files like using a lumberjack",
	Version: "0.1.0", // TODO: make version configurable
	// Use custom error display instead
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	RootCmd.AddCommand(gitignore.GitignoreCmd)
	RootCmd.AddCommand(license.LicenseCmd)
}
