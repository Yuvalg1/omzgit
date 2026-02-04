package branches

import (
	"slices"

	"omzgit/git"
	"omzgit/lib/list"
	"omzgit/messages/refresh"
	"omzgit/program/branches/branch"
	"omzgit/program/popups"
	"omzgit/roller"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case refresh.Msg:
		m.list.SetContent(m.getBranches())

		res, cmd := m.list.Update(msg)
		m.list = res.(list.Model[branch.Model])

		prev := m.list.GetCurrent()
		prev.Active = false

		m.list.ActiveRow = max(slices.IndexFunc(m.list.Children, func(branch branch.Model) bool { return branch.Current }), 0)
		current := m.list.GetCurrent()
		current.Active = true

		return m, cmd

	case tea.WindowSizeMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)

		msg.Width = getWidth(msg.Width)
		msg.Height = getHeight(msg.Height)

		res, cmd := m.list.Update(msg)
		m.list = res.(list.Model[branch.Model])

		return m, cmd

	case list.Msg:
		m.list.SetContent(m.getBranches())
		m.list.Children[m.list.ActiveRow].Active = true

		return m, m.CokeCmd()

	case roller.Msg:
		res, cmd := m.list.UpdateCurrent(msg)
		m.list = res

		return m, cmd

	case tea.KeyMsg:
		if m.list.TextInput.Focused() {
			res, cmd := m.list.Update(msg)
			m.list = res.(list.Model[branch.Model])
			return m, cmd
		}

		switch keypress := msg.String(); keypress {

		case "b":
			return m, popups.Cmd("input", "Name", "Enter A new Branch Name", func(name string) {
				git.Exec("checkout", "-b", name)
			})

		case "c":
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

		case "C":
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

		case "d":
			return m, popups.Cmd("discard", "delete", m.list.GetCurrent().Roller.Name, func() tea.Cmd {
				output, err := git.Exec("branch", "-d", m.list.GetCurrent().Roller.Name)
				if err == nil {
					return nil
				}

				return popups.Cmd("alert", "Delete Error", output, func(name string) {})
			})

		case "D":
			return m, popups.Cmd("discard", "force delete", m.list.GetCurrent().Roller.Name, func() tea.Cmd {
				output, err := git.Exec("branch", "-D", m.list.GetCurrent().Roller.Name)
				if err == nil {
					return nil
				}

				return popups.Cmd("alert", "Force Delete Error", output, func(name string) {})
			})

		case "o":
			m.list.TextInput.SetValue("")

			m.remote = !m.remote
			m.list.SetContent(m.getBranches())

			m.list.ActiveRow = min(m.list.ActiveRow, len(m.list.Children))
			m.list.Children[m.list.ActiveRow].Active = true

			return m, m.CokeCmd()

		case "esc", "/":
			m.list.SetContent(m.getBranches())

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
