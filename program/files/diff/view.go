package diff

import (
	"bytes"
	"strings"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	var buf bytes.Buffer
	quick.Highlight(&buf, m.Content, "go", "terminal", "github-dark")

	block := m.viewport.Block

	rows := strings.Split(buf.String(), "\n")
	var styledRows string
	for _, element := range rows {
		styledRows += element + "\n"
	}

	m.viewport.SetContent(styledRows)
	m.viewport.SetBlock(block[0], block[1], m.width, m.height)

	return lipgloss.NewStyle().Width(m.width).Height(m.height).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#414B53")).
		Render(m.viewport.View())
}
