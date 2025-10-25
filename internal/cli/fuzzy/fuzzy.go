package fuzzy

import tea "github.com/charmbracelet/bubbletea/v2"

// / Fuzzy find from a bunch of options
// /
// / Returns selected item ("" for N/A), if selected (for ensuring check), and possible error
func Find(prompt string, options []string) (string, error) {
	if len(options) < 1 {
		return "", nil
	}

	program := tea.NewProgram(initialModel(prompt, options))
	m, err := program.Run()
	if err != nil {
		return "", err
	}

	mm, ok := m.(model)
	if !ok {
		return "", nil
	}

	return mm.SelectedItem(), nil
}
