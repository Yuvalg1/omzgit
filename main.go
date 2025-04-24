package main

import (
	"fmt"
	"os"
	"os/signal"
	"program/messages"
	"program/program"
	"program/program/branches"
	"program/program/commits"
	"program/program/files"
	"syscall"

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

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGWINCH)

	go func() {
		for range sigChan {
			width, height, _ := term.GetSize(os.Stdout.Fd())
			p.Send(messages.TerminalMsg{Width: width, Height: height})
		}
	}()

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
