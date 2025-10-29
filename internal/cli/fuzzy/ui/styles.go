package ui

import "github.com/charmbracelet/lipgloss/v2"

const (
	listWidth  = 25
	listHeight = 14
)

var (
	titleStyle = lipgloss.NewStyle().PaddingLeft(2)

	listFrameStyle = lipgloss.NewStyle().
			Height(listHeight + 2).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("69")).
			Width(listWidth)

	listItemStyle         = lipgloss.NewStyle()
	selectedListItemStyle = listItemStyle.Foreground(lipgloss.Color("170"))

	listFooterStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
)
