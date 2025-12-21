package fuzzy

import (
	tea "charm.land/bubbletea/v2"
	"go.chardoncs.dev/downjack/internal/cli/fuzzy/ui"
)

// Fuzzy find from a bunch of options.
//
// Returns selected item ("" for N/A), if selected (for ensuring check), and possible error.
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
