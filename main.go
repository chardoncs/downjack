package main

import (
	"context"
	"os"

	"github.com/chardoncs/downjack/internal/cmd"
	"github.com/charmbracelet/fang"
)

func main() {
	if err := fang.Execute(context.Background(), cmd.RootCmd); err != nil {
		os.Exit(1)
	}
}
