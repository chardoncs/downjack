package main

import (
	"context"
	"os"

	"go.chardoncs.dev/downjack/internal/cmd"
	"go.chardoncs.dev/downjack/internal/version"
	"github.com/charmbracelet/fang"
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
