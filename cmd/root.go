package cmd

import (
	"github.com/chardoncs/downjack/cmd/gitignore"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "downjack [commands] [flags] [values]",
	Short: "Set up your gitignore and license files like a lumberjack",
	Version: "0.1.0",
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	rootCmd.AddCommand(gitignore.GitignoreCmd)
}
