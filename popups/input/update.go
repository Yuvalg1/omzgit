package input

import (
	"omzgit/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.textinput.Width = getWidth(msg.Width) - 5

		m.Width = getWidth(msg.Width)
		m.Height = getHeight(msg.Height)
		return m, nil

	case messages.PopupMsg:
		m.textinput.SetValue("")
		m.CallbackFn = msg.Fn.(func(string))
		m.Name = msg.Name
		m.visible = true
		m.textinput.Placeholder = msg.Verb
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "esc":
			m.visible = false
			return m, nil

		case "enter":
			m.CallbackFn(m.textinput.Value())
			m.visible = false
			return m, m.RefreshCmd()

		case "ctrl+c":
			return m, tea.Quit

		default:
			res, cmd := m.textinput.Update(msg)
			m.textinput = res
			return m, cmd
		}
	}

	return m, nil
}
