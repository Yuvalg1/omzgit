package program

import (
	"fmt"
	"strings"

	"omzgit/git"
	"omzgit/messages"
	"omzgit/messages/api"
	"omzgit/messages/mode"
	"omzgit/messages/tick"
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

	case cokeline.Msg:
		res, cmd := m.cokeline.Update(msg)
		m.cokeline = res.(cokeline.Model)
		return m, cmd

	case messages.RollerMsg:
		res, cmd := m.Tabs[m.ActiveTab].Update(msg)
		m.Tabs[m.ActiveTab] = res

		return m, cmd

	case tick.Msg, messages.RefreshMsg:
		res1, cmd1 := m.Tabs[m.ActiveTab].Update(msg)
		m.Tabs[m.ActiveTab] = res1

		res2, cmd2 := m.Popup.Update(msg)
		m.Popup = res2.(popups.Model[popups.InnerModel])
		return m, tea.Batch(cmd1, cmd2)

	case popups.Msg, api.Msg, spinner.TickMsg:
		res, cmd := m.Popup.Update(msg)
		m.Popup = res.(popups.Model[popups.InnerModel])
		return m, cmd

	case mode.Msg:
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
			return m, popups.Cmd("async", "", "fetching", func() tea.Cmd {
				output, err := git.Exec("fetch")
				if err == nil {
					return nil
				}

				return popups.Cmd("alert", "Fetch Error", output, func() tea.Cmd { return nil })
			})

		case "g":
			return m, mode.Cmd("goto")

		case "l":
			return m, popups.Cmd("async", "", "pulling", func() tea.Cmd {
				output, err := git.Exec("pull")
				if err == nil {
					return nil
				}

				return popups.Cmd("alert", "Pull Error", output, func() tea.Cmd { return nil })
			})

		case "L":
			return m, popups.Cmd("async", "", "rebase pulling", func() tea.Cmd {
				output, err := git.Exec("pull", "--rebase")
				if err == nil {
					return nil
				}

				return popups.Cmd("alert", "Rebase Pull Error", output, func() tea.Cmd { return nil })
			})

		case "p":
			return m, popups.Cmd("async", "", "pushing", func() tea.Cmd {
				_, err := git.Exec("push")
				if err == nil {
					return nil
				}

				output, _ := git.Exec("rev-parse", "--abbrev-ref", "HEAD")

				return popups.Cmd("discard", "upstream push", strings.TrimSpace(output), func() tea.Cmd {
					return popups.Cmd("async", "", "upstream pushing", func() tea.Cmd {
						output, _ := git.Exec("rev-parse", "--abbrev-ref", "HEAD")

						output, err := git.Exec("push", "--set-upstream", "origin", strings.TrimSpace(output))
						if err == nil {
							return nil
						}
						return popups.Cmd("alert", "Upstream Error", strings.TrimSpace(output), func() tea.Cmd { return nil })
					})
				})
			})

		case "P":
			output, _ := git.Exec("rev-parse", "--abbrev-ref", "HEAD")

			return m, popups.Cmd("discard", "force push", strings.TrimSpace(output), func() tea.Cmd {
				return popups.Cmd("async", "", "force pushing", func() tea.Cmd {
					output, err := git.Exec("push", "--force")
					if err == nil {
						return nil
					}

					return popups.Cmd("alert", "Force Push Error", strings.TrimSpace(output), func() tea.Cmd {
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

func handlePick(m *Model, key string, msg tea.Msg) (tea.Model, tea.Cmd) {
	escMsg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("esc")}

	m.ActiveTab = key

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
	case "a":
		return m, popups.Cmd("async", "", "opening actions", func() tea.Cmd {
			remote, _ := git.Exec("remote", "get-url", "origin")
			remote = strings.TrimSpace(remote)
			remote = strings.TrimSuffix(remote, ".git")

			url := fmt.Sprintf("%s/actions", remote)

			git.ExecNoOutput("web--browse", url)
			return nil
		})

	case "b":
		return handlePick(m, "Branches", msg)
	case "c":
		return handlePick(m, "Commits", msg)
	case "f":
		return handlePick(m, "Files", msg)

	case "g":
		res, cmd := m.Tabs[m.ActiveTab].Update(msg)
		m.Tabs[m.ActiveTab] = res
		return m, cmd

	case "i":
		return m, popups.Cmd("async", "", "creating issue", func() tea.Cmd {
			remote, _ := git.Exec("remote", "get-url", "origin")
			remote = strings.TrimSpace(remote)
			remote = strings.TrimSuffix(remote, ".git")

			url := fmt.Sprintf("%s/issues/new", remote)

			git.ExecNoOutput("web--browse", url)
			return nil
		})

	case "p":
		return m, popups.Cmd("async", "", "creating pr", func() tea.Cmd {
			remote, _ := git.Exec("remote", "get-url", "origin")
			remote = strings.TrimSpace(remote)
			remote = strings.TrimSuffix(remote, ".git")

			branch, _ := git.Exec("rev-parse", "--abbrev-ref", "HEAD")
			branch = strings.TrimSpace(branch)

			url := fmt.Sprintf("%s/compare/%s?expand=1", remote, branch)

			git.ExecNoOutput("web--browse", url)
			return nil
		})

	default:
		return m, nil
	}
}
