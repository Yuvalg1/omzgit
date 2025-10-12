package program

import (
	"strings"

	"omzgit/consts"
	"omzgit/git"
	"omzgit/messages"
	"omzgit/program/cokeline"
	"omzgit/program/popups"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = getWidth(msg.Width)
		m.Height = getHeight(msg.Height)

		msg.Width = getWidth(msg.Width)
		msg.Height = getHeight(msg.Height)

		res1, cmd1 := m.cokeline.Update(msg)
		m.cokeline = res1.(cokeline.Model)

		res2, cmd2 := m.Popup.Update(msg)
		m.Popup = res2.(popups.Model[popups.InnerModel])

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

	case messages.RollerMsg:
		res, cmd := m.Tabs[m.ActiveTab].Update(msg)
		m.Tabs[m.ActiveTab] = res

		return m, cmd

	case messages.TickMsg, messages.RefreshMsg:
		res1, cmd1 := m.Tabs[m.ActiveTab].Update(msg)
		m.Tabs[m.ActiveTab] = res1

		res2, cmd2 := m.Popup.Update(msg)
		m.Popup = res2.(popups.Model[popups.InnerModel])
		return m, tea.Batch(cmd1, cmd2)

	case messages.PopupMsg, messages.ApiMsg, spinner.TickMsg:
		res, cmd := m.Popup.Update(msg)
		m.Popup = res.(popups.Model[popups.InnerModel])
		return m, cmd

	case messages.ModeMsg:
		m.mode = msg.Mode
		return m, nil

	case tea.KeyMsg:
		current := m.Popup.GetCurrent()
		if current.GetVisible() {
			res, cmd := m.Popup.Update(msg)
			m.Popup = res.(popups.Model[popups.InnerModel])

			cmds := []tea.Cmd{cmd}

			return m, tea.Batch(cmds...)
		}

		if m.mode == "goto" {
			return pickTab(&m, msg)
		}

		if m.mode == "search" {
			res, cmd := m.Tabs[m.ActiveTab].Update(msg)
			m.Tabs[m.ActiveTab] = res
			return m, cmd
		}

		switch keypress := msg.String(); keypress {
		case "f":
			return m, m.PopupCmd("async", "", "fetching", func() tea.Cmd {
				output, err := git.Exec("fetch")
				if err == nil {
					return nil
				}

				return m.PopupCmd("alert", "Fetch Error", output, func() tea.Cmd { return nil })
			})

		case "g":
			return m, m.ModeCmd("goto")

		case "l":
			return m, m.PopupCmd("async", "", "pulling", func() tea.Cmd {
				output, err := git.Exec("pull")
				if err == nil {
					return nil
				}

				return m.PopupCmd("alert", "Pull Error", output, func() tea.Cmd { return nil })
			})

		case "p":
			return m, m.PopupCmd("async", "", "pushing", func() tea.Cmd {
				_, err := git.Exec("push")
				if err == nil {
					return nil
				}

				output, _ := git.Exec("rev-parse", "--abbrev-ref", "HEAD")

				return m.PopupCmd("discard", "upstream push", strings.TrimSpace(output), func() tea.Cmd {
					return m.PopupCmd("async", "", "upstream pushing", func() tea.Cmd {
						output, _ := git.Exec("rev-parse", "--abbrev-ref", "HEAD")

						output, err := git.Exec("push", "--set-upstream", "origin", strings.TrimSpace(output))
						if err == nil {
							return nil
						}
						return m.PopupCmd("alert", "Upstream Error", strings.TrimSpace(output), func() tea.Cmd { return nil })
					})
				})
			})

		case "P":
			output, _ := git.Exec("rev-parse", "--abbrev-ref", "HEAD")

			return m, m.PopupCmd("discard", "force push", strings.TrimSpace(output), func() tea.Cmd {
				return m.PopupCmd("async", "", "force pushing", func() tea.Cmd {
					output, err := git.Exec("push", "--force")
					if err == nil {
						return nil
					}

					return m.PopupCmd("alert", "Force Push Error", strings.TrimSpace(output), func() tea.Cmd {
						return nil
					})
				})
			})

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

func handlePick(m *Model, key int, msg tea.Msg) (tea.Model, tea.Cmd) {
	escMsg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("esc")}

	m.ActiveTab = key - 1

	res1, cmd1 := m.Tabs[m.ActiveTab].Update(escMsg)
	m.Tabs[m.ActiveTab] = res1

	res2, cmd2 := m.cokeline.Update(msg)
	m.cokeline = res2.(cokeline.Model)

	return m, tea.Batch(cmd1, cmd2)
}

func (m Model) RefreshCmd() tea.Cmd {
	return func() tea.Msg {
		return messages.RefreshMsg{}
	}
}

func pickTab(m *Model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	m.mode = ""

	switch keypress := msg.String(); keypress {
	case "b":
		return handlePick(m, consts.BRANCHES, msg)
	case "c":
		return handlePick(m, consts.COMMITS, msg)
	case "f":
		return handlePick(m, consts.FILES, msg)

	case "g":
		res, cmd := m.Tabs[m.ActiveTab].Update(msg)
		m.Tabs[m.ActiveTab] = res
		return m, cmd

	default:
		return m, nil
	}
}
