package pointers

import (
	"omzgit/messages"
	"omzgit/popups/commit"
	"omzgit/popups/discard"
	"omzgit/popups/input"
	"omzgit/popups/reset"
	"omzgit/program"
	"omzgit/roller"
)

var (
	_ messages.Refresher = (*commit.Model)(nil)

	_ messages.Refresher = (*discard.Model)(nil)

	_ messages.Refresher = (*input.Model)(nil)

	_ messages.Refresher = (*program.Model)(nil)

	_ messages.Refresher = (*reset.Model)(nil)

	_ messages.Rollerer = (*roller.Model)(nil)
)
