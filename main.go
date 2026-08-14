package main

import (
	"os"
	"os/exec"

	"omzgit/program"
	"omzgit/program/branches"
	"omzgit/program/commits"
	"omzgit/program/files"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
)

func main() {
	width, height, _ := term.GetSize(os.Stdout.Fd())

	command := make(chan []string, 1)
	page := make(chan string, 1)
	page <- "Files"

	for len(command) == 0 {
		m := program.InitialModel(
			[]program.ExtendedModel{
				{Title: "Files", Tab: files.InitialModel(width, height)},
				{Title: "Branches", Tab: branches.InitialModel(width, height, "Branches")},
				{Title: "Commits", Tab: commits.InitialModel(width, height, "Commits")},
			},
			width,
			height,
			command,
			page,
		)

		p := tea.NewProgram(m)

		_, _ = p.Run()

		if len(command) == 0 {
			break
		}

		command := <-command
		cmd := exec.Command("git", command...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		_ = cmd.Run()
	}
}
