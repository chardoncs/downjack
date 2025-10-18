package main

import (
	"context"
	"os"

	"github.com/chardoncs/downjack/internal/cmd"
	"github.com/chardoncs/downjack/internal/version"
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
