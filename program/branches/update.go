package branches

import (
	"slices"

	"omzgit/env"
	"omzgit/git"
	"omzgit/lib/list"
	"omzgit/messages/mode"
	"omzgit/messages/refresh"
	"omzgit/popups/help"
	"omzgit/program/branches/branch"
	"omzgit/program/popups"
	"omzgit/roller"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case refresh.Msg:
		snapshot := m.getSnapshot()
		return m, list.Cmd(func() []branch.Model { return getBranches(snapshot) }, m.list.ActiveRow, "", m.CokeCmd())

	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)

		msg.Width = getWidth(msg.Width)
		msg.Height = getHeight(msg.Height)

		res, cmd := m.list.Update(msg)
		m.list = res.(list.Model[branch.Model])

		return m, cmd

	case list.Msg[branch.Model]:
		m.list.Children = msg.Children
		m.total = msg.Total
		m.total = len(m.list.Children)
		m.list.ActiveRow = min(len(m.list.Children)-1, msg.Active)
		m.list.Children[m.list.ActiveRow].Active = true

		res, cmd := m.list.Update(msg.Msg)
		m.list = res.(list.Model[branch.Model])

		return m, tea.Batch(cmd, msg.Cmd)

	case roller.Msg:
		res, cmd := m.list.UpdateCurrent(msg)
		m.list = res

		return m, cmd

	case mode.Msg:
		res, cmd := m.list.Update(msg)
		m.list = res.(list.Model[branch.Model])
		return m, cmd

	case tea.KeyMsg:
		if m.list.TextInput.Focused() {
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[branch.Model])

			return m, cmd
		}

		switch keypress := msg.String(); keypress {
		case env.Branches.CheckoutB.Msg:
			return m, popups.Cmd("input", "Name", "Enter A new Branch Name", func(name string) {
				git.Exec("checkout", "-b", name)
			})

		case env.Branches.Checkout.Msg, env.Branches.CheckoutForce.Msg:
			index := slices.IndexFunc(m.list.Children, func(branch branch.Model) bool { return branch.Current })

			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[branch.Model])

			if index != -1 && index != m.list.ActiveRow && m.list.GetCurrent().Current {
				m.list.Children[index].Current = false
			}

			return m, tea.Batch(cmd, m.CokeCmd())

		case env.Branches.Origin.Msg:
			m.list.TextInput.SetValue("")
			m.remote = !m.remote

			snapshot := m.getSnapshot()
			return m, list.Cmd(func() []branch.Model { return getBranches(snapshot) }, m.list.ActiveRow, msg, m.CokeCmd())

		case "?":
			return m, popups.Cmd("help", "", "", func() ([]env.Option, func() tea.Cmd) {
				return append(help.GetEnvOptions(env.Branches), help.GetEnvOptions(env.Program)...),
					func() tea.Cmd {
						return nil
					}
			})

		case env.Branches.Refresh.Msg, env.Branches.Search.Msg:
			m.list.SetContent(getBranches(m.getSnapshot()))
			m.total = len(m.list.Children)

			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[branch.Model])

			return m, tea.Batch(cmd, m.CokeCmd())

		default:
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[branch.Model])

			return m, tea.Batch(cmd, m.CokeCmd())
		}
	}

	return m, nil
}
