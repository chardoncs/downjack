package cmd

import (
	"fmt"
	"os"

	"github.com/chardoncs/downjack/cmd/gitignore"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: fmt.Sprintf("%s [commands] [flags] [values]", os.Args[0]),
	Short: "Setup your gitignore and license files like using a lumberjack",
	Version: "0.1.0",
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	rootCmd.AddCommand(gitignore.GitignoreCmd)
}
