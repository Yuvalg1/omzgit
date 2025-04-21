package program

import (
	"fmt"
	"program/consts"
	"program/messages"
	"program/program/popup"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.TerminalMsg:
		m.Width = GetWidth(msg.Width)
		m.Height = GetHeight(msg.Height)

		msg.Width = GetWidth(msg.Width)
		msg.Height = GetHeight(msg.Height)

		res1, cmd1 := m.Popup.Update(msg)
		m.Popup = res1.(popup.Model)

		res2, cmd2 := m.Tabs[m.ActiveTab].Update(msg)
		m.Tabs[m.ActiveTab] = res2
		return m, tea.Batch(cmd1, cmd2)

	case messages.TitleMsg:
		m.title = msg.Title
		return m, nil

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

		if m.pickMode {
			return pickTab(&m, msg)
		}

		switch keypress := msg.String(); keypress {
		case " ", "esc", "backspace":
			m.pickMode = true
			return m, nil

		case "ctrl+c", "q":
			return m, tea.Quit

		default:
			res, cmd := m.Tabs[m.ActiveTab].Update(msg)
			m.Tabs[m.ActiveTab] = res
			return m, cmd
		}
	}

	return m, nil
}

func handlePick(m *Model, key int) (tea.Model, tea.Cmd) {
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(fmt.Sprint(key))}

	m.ActiveTab = key - 1
	res, cmd := m.Tabs[m.ActiveTab].Update(msg)
	m.Tabs[m.ActiveTab] = res
	return m, cmd
}

func (m Model) DeleteCmd() tea.Cmd {
	return func() tea.Msg {
		return messages.DeletedMsg{}
	}
}

func pickTab(m *Model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	m.pickMode = false

	switch keypress := msg.String(); keypress {
	case "b":
		m.title = "Branches"
		return handlePick(m, consts.BRANCHES)
	case "c":
		m.title = "Commits"
		return handlePick(m, consts.COMMITS)
	case "f":
		m.title = "Filed Changed"
		return handlePick(m, consts.FILES)

	case " ", "esc", "backspace":
		m.pickMode = true
		return m, nil

	default:
		return m, nil
	}
}
