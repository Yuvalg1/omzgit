package cokeline

import (
	"omzgit/consts"
	"omzgit/default/colors"
	"omzgit/default/colors/bg"
	"omzgit/default/colors/gray"

	"github.com/charmbracelet/lipgloss"
)

var partStyle = lipgloss.NewStyle().
	Background(bg.C[3]).
	Foreground(bg.C[0])

var colorTones = map[string][]lipgloss.Color{
	"Branches": {colors.Purple, colors.Yellow, colors.Blue},
	"Commits":  {colors.Aqua, colors.Purple, colors.Yellow},
	"Files":    {colors.Yellow, colors.Green, colors.Red},
}

func (m Model) View() string {
	semitone := 1
	if !m.Primary {
		semitone += 1
	}

	left := lipgloss.NewStyle().
		Background(colorTones[m.Left][0]).
		Bold(true).
		Foreground(bg.C[0]).
		Padding(0, 1).
		Render(m.Left)

	right := lipgloss.NewStyle().
		Background(gray.C[1]).
		Padding(0, 1).
		Render(m.Right)

	center := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Background(bg.C[3]).
		Bold(true).
		Foreground(colorTones[m.Left][semitone]).
		Padding(0, 1).
		Width(m.width - lipgloss.Width(left) - lipgloss.Width(right)).
		Render(consts.TrimRight(m.Center, m.width-lipgloss.Width(left)-lipgloss.Width(right)-2))

	return lipgloss.NewStyle().
		Width(m.width).
		Height(1).
		Padding(0, 0, 1).
		Render(left + center + right)
}
