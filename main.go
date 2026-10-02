package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"

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
	error := make(chan string, 1)
	error <- ""

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
			error,
		)

		p := tea.NewProgram(m)

		_, _ = p.Run()

		if len(command) == 0 {
			break
		}

		var stderr bytes.Buffer

		command := <-command
		cmd := exec.Command("git", command...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = io.MultiWriter(os.Stderr, &stderr)

		_ = cmd.Run()
		error <- strings.TrimSpace(stderr.String())
	}
}
