package cli

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

var (
	infoStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightCyan)
	warnStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightYellow)

	infoArrow = infoStyle.Render(">>")
	warnArrow = warnStyle.Render(">>")

	warningText = warnStyle.Bold(true).Render("Warning:")
)

func Infof(format string, a ...any) {
	lipgloss.Println(
		infoArrow,
		fmt.Sprintf(format, a...),
	)
}

func Warnf(format string, a ...any) {
	lipgloss.Println(
		warnArrow,
		warningText,
		fmt.Sprintf(format, a...),
	)
}

func InfoProgressf(format string, a ...any) {
	lipgloss.Print(
		infoArrow,
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
		fmt.Printf("%3d) %s\n", i+1, item)
	}
}
