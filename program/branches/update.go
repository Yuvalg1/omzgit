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
		m.list.ActiveRow = msg.Active
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

		case env.Branches.Checkout.Msg, env.Branches.Checkout.AltMsg:
			output, err := git.Exec("checkout", m.list.GetCurrent().Roller.Name)
			if err != nil {
				return m, popups.Cmd("alert", "checkout error", output, func(name string) {})
			}

			current := slices.IndexFunc(m.list.Children, func(branch branch.Model) bool { return branch.Current })

			if current != -1 {
				m.list.Children[current].Current = false
			}

			m.list.Children[m.list.ActiveRow].Current = true
			return m, nil

		case env.Branches.CheckoutForce.Msg:
			return m, popups.Cmd("discard", "force checkout", m.list.GetCurrent().Roller.Name, func() tea.Cmd {
				output, err := git.Exec("checkout", "-f", m.list.GetCurrent().Roller.Name)
				if err != nil {
					return popups.Cmd("alert", "checkout error", output, func(name string) {})
				}

				current := slices.IndexFunc(m.list.Children, func(branch branch.Model) bool { return branch.Current })

				if current != -1 {
					m.list.Children[current].Current = false
				}

				m.list.Children[m.list.ActiveRow].Current = true
				return nil
			})

		case env.Branches.Delete.Msg:
			return m, popups.Cmd("discard", "delete", m.list.GetCurrent().Roller.Name, func() tea.Cmd {
				output, err := git.Exec("branch", "-d", m.list.GetCurrent().Roller.Name)
				if err == nil {
					return nil
				}

				return popups.Cmd("alert", "Delete Error", output, func(name string) {})
			})

		case env.Branches.DeleteForce.Msg:
			return m, popups.Cmd("discard", "force delete", m.list.GetCurrent().Roller.Name, func() tea.Cmd {
				output, err := git.Exec("branch", "-D", m.list.GetCurrent().Roller.Name)
				if err == nil {
					return nil
				}

				return popups.Cmd("alert", "Force Delete Error", output, func(name string) {})
			})

		case env.Branches.Origin.Msg:
			m.list.TextInput.SetValue("")
			m.remote = !m.remote

			snapshot := m.getSnapshot()
			return m, list.Cmd(func() []branch.Model { return getBranches(snapshot) }, m.list.ActiveRow, msg, m.CokeCmd())

		case "?":
			return m, popups.Cmd("help", "", "", func() []env.Option {
				return append(help.GetEnvOptions(env.Branches), help.GetEnvOptions(env.Program)...)
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
