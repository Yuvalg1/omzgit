package program

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	style := lipgloss.NewStyle().Border(lipgloss.DoubleBorder(), false, true, true).Height(m.Height - 2).Width(m.Width - 2).Background(lipgloss.Color("red"))

	titleStyle := lipgloss.NewStyle().Width(m.Width - 2).Align(lipgloss.Center)
	aligned := titleStyle.Render(m.title)

	if m.Popup.Visible {
		return style.Render(m.Popup.View())
	}

	return "╔" + replaceBorder(" "+m.title+" ", aligned) + "╗\n" + style.Render(m.Tabs[m.ActiveTab].View())
}

func replaceBorder(original string, aligned string) string {
	prefixLen := strings.Index(aligned, original)
	suffixLen := len(aligned) - prefixLen - len(original)

	return strings.Repeat("═", prefixLen) + original + strings.Repeat("═", suffixLen)
}
