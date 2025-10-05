package cmd

import (
	"github.com/chardoncs/downjack/cmd/gitignore"
	"github.com/chardoncs/downjack/internal/cli"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "downjack [commands] [flags] [values]",
	Short: "Set up your gitignore and license files like a lumberjack",
	Version: "0.1.0",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	err := rootCmd.Execute()
	cli.CheckErr(err)
}

func init() {
	rootCmd.AddCommand(gitignore.GitignoreCmd)
}
