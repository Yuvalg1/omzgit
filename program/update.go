package program

import (
	"program/consts"
	"program/messages"
	"program/program/cokeline"
	"program/program/popup"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.TerminalMsg:
		m.Width = getWidth(msg.Width)
		m.Height = getHeight(msg.Height)

		msg.Width = getWidth(msg.Width)
		msg.Height = getHeight(msg.Height)

		res1, cmd1 := m.cokeline.Update(msg)
		m.cokeline = res1.(cokeline.Model)

		res2, cmd2 := m.Popup.Update(msg)
		m.Popup = res2.(popup.Model)

		cmds := []tea.Cmd{cmd1, cmd2}
		for index, element := range m.Tabs {
			res, cmd := element.Update(msg)
			m.Tabs[index] = res
			cmds = append(cmds, cmd)
		}

		return m, tea.Batch(cmds...)

	case messages.CokeMsg:
		res, cmd := m.cokeline.Update(msg)
		m.cokeline = res.(cokeline.Model)
		return m, cmd

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

	case messages.ModeMsg:
		m.mode = msg.Mode
		return m, nil

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

		if m.mode == "search" {
			res, cmd := m.Tabs[m.ActiveTab].Update(msg)
			m.Tabs[m.ActiveTab] = res
			return m, cmd
		}

		if m.mode == "goto" {
			return pickTab(&m, msg)
		}

		switch keypress := msg.String(); keypress {
		case "g":
			return m, m.ModeCmd("goto")

		case "]":
			index := (m.ActiveTab+1)%len(m.Tabs) + 1
			return handlePick(&m, index)

		case "[":
			index := (m.ActiveTab-1+len(m.Tabs))%len(m.Tabs) + 1
			return handlePick(&m, index)

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
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("esc")}

	m.ActiveTab = key - 1
	m.cokeline.ActiveCoke = m.ActiveTab

	res1, cmd1 := m.Tabs[m.ActiveTab].Update(msg)
	m.Tabs[m.ActiveTab] = res1

	return m, tea.Batch(cmd1)
}

func (m Model) DeleteCmd() tea.Cmd {
	return func() tea.Msg {
		return messages.DeletedMsg{}
	}
}

func pickTab(m *Model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	m.mode = ""

	switch keypress := msg.String(); keypress {
	case "b":
		return handlePick(m, consts.BRANCHES)
	case "c":
		return handlePick(m, consts.COMMITS)
	case "f":
		return handlePick(m, consts.FILES)

	case "g":
		res, cmd := m.Tabs[m.ActiveTab].Update(msg)
		m.Tabs[m.ActiveTab] = res
		return m, cmd

	default:
		return m, nil
	}
}
