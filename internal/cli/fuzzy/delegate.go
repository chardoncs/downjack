package fuzzy

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/v2/list"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type itemDelegate struct{}

func (self itemDelegate) Height() int {
	return 1
}

func (self itemDelegate) Spacing() int {
	return 0
}

func (self itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}

func (self itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	itemStr, ok := listItem.(item)
	if !ok {
		return
	}

	selectionMark := "  "
	style := listItemStyle

	if index == m.Index() {
		selectionMark = "> "
		style = selectedListItemStyle
	}

	fmt.Fprint(w, style.Render(fmt.Sprintf("%s %s", selectionMark, itemStr)))
}
