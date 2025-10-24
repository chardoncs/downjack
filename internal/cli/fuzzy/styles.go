package fuzzy

import "github.com/charmbracelet/lipgloss/v2"

var (
	listItemStyle         = lipgloss.NewStyle().PaddingLeft(2)
	selectedListItemStyle = listItemStyle.Foreground(lipgloss.Color("170"))
)
