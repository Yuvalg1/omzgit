package border

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func GetTopBorder(title string, width int) string {
	titleStyle := lipgloss.NewStyle().Width(width).Align(lipgloss.Center)
	aligned := titleStyle.Render(title)

	return "╔" + ReplaceBorder(" "+title+" ", aligned) + "╗\n"
}

func ReplaceBorder(original string, aligned string) string {
	prefixLen := strings.Index(aligned, original)
	suffixLen := len(aligned) - prefixLen - len(original)

	return strings.Repeat("═", prefixLen) + original + strings.Repeat("═", suffixLen)
}
