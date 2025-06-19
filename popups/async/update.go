package async

import (
	"omzgit/messages"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.PopupMsg:
		m.callbackFn = msg.Fn.(func() tea.Cmd)
		m.title = msg.Name
		m.visible = true
		return m, tea.Batch(m.spinner.Tick, m.ApiCmd())

	case messages.ApiMsg:
		m.visible = false
		return m, msg.Response

	case messages.TerminalMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)
		return m, nil

	case spinner.TickMsg:
		res, cmd := m.spinner.Update(msg)
		m.spinner = res
		return m, cmd

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "esc":
			m.visible = false
			return m, nil

		default:
			return m, nil
		}

	}
	return m, nil
}
