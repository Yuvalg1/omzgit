package env

type branches struct {
	Checkout      Option
	CheckoutB     Option
	CheckoutForce Option
	Delete        Option
	DeleteForce   Option
	Origin        Option
	Rebase        Option
	Merge         Option

	Refresh Option
	Search  Option
	Yank    Option

	Up     Option
	Down   Option
	Bottom Option
}

var Branches = branches{
	Checkout: Option{
		Msg:         "c",
		Description: "checks out to a new branch",
		AltMsg:      "enter",
	},
	CheckoutB: Option{
		Msg:         "b",
		Description: "creates a new branch and checkouts to it",
	},
	CheckoutForce: Option{
		Msg:         "C",
		Description: "forces checkout to a new branch",
	},
	Delete: Option{
		Msg:         "d",
		Description: "opens option to delete branch",
	},
	DeleteForce: Option{
		Msg:         "D",
		Description: "opens option to force delete branch",
	},
	Origin: Option{
		Msg:         "o",
		Description: "switches between local and origin branches",
	},
	Rebase: Option{
		Msg:         "r",
		Description: "rebases to currently highlighted branch",
	},
	Merge: Option{
		Msg:         "m",
		Description: "merges to currently highlighted branch",
	},

	Refresh: Refresh,
	Search:  Search,
	Yank: Option{
		Msg:         "y",
		Description: "yanks the current branch name to clipboard",
	},

	Up: Option{
		Msg:         "up",
		Description: "highlights previous branch",
		AltMsg:      "k",
	},
	Down: Option{
		Msg:         "down",
		Description: "highlights next branch",
		AltMsg:      "j",
	},
	Bottom: Option{
		Msg:         "G",
		Description: "highlights last branch",
	},
}
