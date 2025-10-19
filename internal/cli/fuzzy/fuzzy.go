package fuzzy

import tea "github.com/charmbracelet/bubbletea/v2"

/// Fuzzy find from a bunch of options
///
/// Returns selected index (-1 for N/A), if selected (for ensuring check), and possible error
func Find(prompt string, options []string) (int, bool, error) {
	if len(options) < 1 {
		return -1, false, nil
	}

	program := tea.NewProgram(initialModel(prompt, options))
	m, err := program.Run()
	if err != nil {
		return -1, false, err
	}

	mm, ok := m.(model)
	if !ok {
		return -1, false, nil
	}

	return mm.SelectedIndex(), true, nil
}
