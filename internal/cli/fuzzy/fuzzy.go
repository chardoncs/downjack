package fuzzy

import (
	"github.com/chardoncs/downjack/internal/cli/fuzzy/ui"
	tea "github.com/charmbracelet/bubbletea/v2"
)

// Fuzzy find from a bunch of options
//
// Returns selected item ("" for N/A), if selected (for ensuring check), and possible error
func Find(prompt string, options []string, initialInput string) (string, error) {
	if len(options) < 1 {
		return "", nil
	}

	program := tea.NewProgram(ui.InitialModel(prompt, options, initialInput))
	m, err := program.Run()
	if err != nil {
		return "", err
	}

	mm, ok := m.(ui.MainModel)
	if !ok {
		return "", nil
	}

	selected, prs := mm.SelectedItem()
	if prs {
		return selected, nil
	}

	return "", nil
}
