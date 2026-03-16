package diff

import (
	"strings"

	"omzgit/default/colors"
	"omzgit/default/colors/bg"
	"omzgit/git"
	"omzgit/program/files/row"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	viewport viewport.Model

	Staged bool
	path   string
}

func InitialModel(row row.Model, width int, height int) Model {
	viewport := viewport.New(getWidth(width), getHeight(height))

	m := Model{
		viewport: viewport,
		Staged:   row.Staged,
		path:     row.Roller.Name,
	}

	if row.Active {
		m.viewport.SetContent(m.getDiff())
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width - 2
}

func getHeight(height int) int {
	return height - 2
}

func (m Model) colorLines(output string) string {
	contentLines := strings.Split(output, "\n")
	style := lipgloss.NewStyle().Background(bg.C[0]).Width(m.viewport.Width)

	removeStyle := lipgloss.NewStyle().Foreground(colors.Red).Inherit(style)
	addStyle := lipgloss.NewStyle().Foreground(colors.Green).Inherit(style)
	infoStyle := lipgloss.NewStyle().Foreground(colors.Blue).Inherit(style)

	content := ""

	for _, element := range contentLines[:len(contentLines)-1] {
		if strings.HasPrefix(element, "-") {
			content += removeStyle.Render(element) + "\n"
		} else if strings.HasPrefix(element, "+") {
			content += addStyle.Render(element) + "\n"
		} else if strings.HasPrefix(element, "@@") {
			parts := strings.Split(element[2:], "@@")
			content += infoStyle.Render("@@ "+parts[0]+" @@") + "\n"
			if parts[1] != "" {
				content += style.Render(parts[1]) + "\n"
			}
		} else {
			content += style.Render(element) + "\n"
		}
	}

	return content
}

func (m Model) getDiff() string {
	if m.Staged {
		output, _ := git.Exec("diff", "--staged", m.path)
		return m.colorLines(output)
	}

	_, err := git.Exec("ls-files", "--error-unmatch", m.path)

	if err == nil {
		output, _ := git.Exec("diff", "--", m.path)
		return m.colorLines(output)
	}

	output, _ := git.Exec("diff", "--no-index", "/dev/null", m.path)
	return m.colorLines(output)
}
