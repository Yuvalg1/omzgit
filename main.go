package main

import (
	"fmt"
	"os"
	"program/messages"
	"program/program"
	"program/program/branches"
	"program/program/commits"
	"program/program/files"

	tsize "github.com/kopoli/go-terminal-size"

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

	listener, err := tsize.NewSizeListener()
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	go func() {
		for size := range listener.Change {
			p.Send(messages.TerminalMsg{Width: size.Width, Height: size.Height})
		}
	}()

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
