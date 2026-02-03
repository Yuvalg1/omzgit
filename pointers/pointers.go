package pointers

import (
	"omzgit/lib/list"
	"omzgit/messages"
	"omzgit/popups/async"
	"omzgit/popups/commit"
	"omzgit/popups/discard"
	"omzgit/popups/input"
	"omzgit/popups/reset"
	"omzgit/program"
	"omzgit/program/branches"
	"omzgit/program/files"
	"omzgit/roller"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	_ messages.Apier = (*async.Model)(nil)

	_ messages.Cokerer = (*branches.Model)(nil)

	_ messages.Refresher = (*commit.Model)(nil)

	_ messages.Refresher = (*discard.Model)(nil)

	_ messages.Cokerer = (*files.Model)(nil)
	_ messages.Ticker  = (*files.Model)(nil)

	_ messages.Refresher = (*input.Model)(nil)

	_ messages.Moderer = (*list.Model[tea.Model])(nil)

	_ messages.Moderer   = (*program.Model)(nil)
	_ messages.Refresher = (*program.Model)(nil)

	_ messages.Refresher = (*reset.Model)(nil)

	_ messages.Rollerer = (*roller.Model)(nil)
)
