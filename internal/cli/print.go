package cli

import (
	"fmt"

	"github.com/fatih/color"
)

func Info(format string, a ...any) {
	fmt.Println(
		color.CyanString(">>"),
		fmt.Sprintf(format, a...),
	)
}

func Warning(format string, a ...any) {
	fmt.Println(
		color.YellowString(">>"),
		color.New(color.Bold).Sprint("Warning:"),
		fmt.Sprintf(format, a...),
	)
}

func InfoProgress(format string, a ...any) {
	fmt.Print(
		color.CyanString(">>"),
		" ",
		fmt.Sprintf(format, a...),
		"... ",
	)
}

func Done() {
	fmt.Println("done")
}

func PrintItems(items []string) {
	for i, item := range items {
		fmt.Printf("%3d) %s\n", i + 1, item)
	}
}
