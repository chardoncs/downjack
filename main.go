package main

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"go.chardoncs.dev/downjack/internal/cmd"
	"go.chardoncs.dev/downjack/internal/version"
)

func main() {
	if err := fang.Execute(
		context.Background(),
		cmd.RootCmd,
		fang.WithVersion(version.Version),
	); err != nil {
		os.Exit(1)
	}
}
