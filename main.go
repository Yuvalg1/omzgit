package main

import (
	"fmt"
	"omzgit/program"
	"omzgit/program/branches"
	"omzgit/program/commits"
	"omzgit/program/files"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
)

func main() {
	width, height, _ := term.GetSize(os.Stdout.Fd())

	m := program.InitialModel(
		[]program.ExtendedModel{
			{Title: "Files", Tab: files.InitialModel(width, height)},
			{Title: "Branches", Tab: branches.InitialModel(width, height, "Branches")},
			{Title: "Commits", Tab: commits.InitialModel(width, height, "Commits")},
		},
		width,
		height,
	)

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
