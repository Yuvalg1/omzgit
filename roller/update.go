package roller

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = getWidth(msg.Width)
		return m, m.RollerCmd()

	case Msg:
		m.Offset = (m.Offset + 1) % (len(m.Name) + 1)

		return m, m.RollerCmd()

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case "j", "k", "down", "up", "g", "G", "/", "esc":
			m.Offset = 0
			return m, m.InitRollerCmd()

		default:
			return m, nil
		}
	}

	return m, nil
}
