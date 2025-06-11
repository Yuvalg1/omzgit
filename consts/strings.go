package consts

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func TrimRight(name string, width int) string {
	parts := strings.Split(name, "\n")

	var trimmed string
	for _, element := range parts {
		trimmed += element[:min(lipgloss.Width(element), width)] + "\n"
	}

	trimmed = trimmed[:len(trimmed)-1]
	return trimmed
}

func PadTitle(title string, width int) string {
	padding := width - len(title) - 4
	parity := len(title) % 2

	return "┌" + strings.Repeat("─", padding/2) + " " + title + " " + strings.Repeat("─", padding/2+parity) + "┐" + "\n"
}
