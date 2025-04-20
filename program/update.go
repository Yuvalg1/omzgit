package program

import (
	"program/consts"
	"program/messages"
	"program/program/popup"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.TerminalMsg:
		m.Height = msg.Height
		m.Width = msg.Width

		msg.Height = msg.Height - 2
		msg.Width = msg.Width - 2

		res1, cmd1 := m.Popup.Update(msg)
		m.Popup = res1.(popup.Model)

		res2, cmd2 := m.Tabs[m.ActiveTab].Update(msg)
		m.Tabs[m.ActiveTab] = res2
		return m, tea.Batch(cmd1, cmd2)

	case messages.TickMsg, messages.DeletedMsg:
		res, cmd := m.Tabs[m.ActiveTab].Update(msg)
		m.Tabs[m.ActiveTab] = res
		return m, cmd

	case messages.PopupMsg:
		m.Popup.Fn = msg.Fn
		m.Popup.Name = msg.Name
		m.Popup.Visible = true

		res, cmd := m.Tabs[m.ActiveTab].Update(msg)
		m.Tabs[m.ActiveTab] = res
		return m, cmd

	case tea.KeyMsg:
		if m.Popup.Visible {
			res, cmd := m.Popup.Update(msg)
			m.Popup = res.(popup.Model)

			cmds := []tea.Cmd{cmd}

			if msg.String() == "y" {
				cmds = append(cmds, m.DeleteCmd())
			}

			return m, tea.Batch(cmds...)
		}
		switch keypress := msg.String(); keypress {

		case "b":
			if handledInner(m, msg) {
				return m, nil
			}

			m.ActiveTab = consts.BRANCHES
			return m, nil

		case "c":
			if handledInner(m, msg) {
				return m, nil
			}

			m.ActiveTab = consts.COMMITS
			return m, nil

		case "ctrl+c", "q":
			return m, tea.Quit

		case "esc":
			m.ActiveTab = consts.FILES
			res, cmd := m.Tabs[consts.FILES].Update(msg)
			m.Tabs[consts.FILES] = res

			return m, cmd

		default:
			res, cmd := m.Tabs[m.ActiveTab].Update(msg)
			m.Tabs[m.ActiveTab] = res
			return m, cmd
		}
	}

	return m, nil
}

func handledInner(m Model, msg tea.Msg) bool {
	if m.ActiveTab != consts.FILES {
		m.Tabs[m.ActiveTab], _ = m.Tabs[m.ActiveTab].Update(msg)
		return true
	}
	return false
}

func (m Model) DeleteCmd() tea.Cmd {
	return func() tea.Msg {
		return messages.DeletedMsg{}
	}
}
