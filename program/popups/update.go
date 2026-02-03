package popups

import (
	"omzgit/messages/api"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		var cmds []tea.Cmd
		for index, element := range m.Popups {
			res, cmd := element.Update(msg)
			m.Popups[index] = res.(T)
			cmds = append(cmds, cmd)
		}
		return m, tea.Batch(cmds...)

	case Msg:
		m.current = msg.Type

		res, cmd := m.Popups[m.current].Update(msg)
		m.Popups[m.current] = res.(T)

		return m, cmd

	case api.Msg:
		res, cmd := m.Popups[m.current].Update(msg)
		m.Popups[m.current] = res.(T)

		return m, cmd

	case spinner.TickMsg:
		if m.current == "async" && m.Popups[m.current].GetVisible() {
			res, cmd := m.Popups[m.current].Update(msg)
			m.Popups[m.current] = res.(T)
			return m, cmd
		}
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		default:
			res, cmd := m.Popups[m.current].Update(msg)
			m.Popups[m.current] = res.(T)
			return m, cmd
		}
	}

	return m, nil
}
