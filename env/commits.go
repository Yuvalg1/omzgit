package env

type commits struct {
	Checkout      Option
	CheckoutForce Option
	Reset         Option
	CherryPick    Option
	Yank          Option

	Up     Option
	Down   Option
	Bottom Option

	Refresh Option
	Search  Option
}

var Commits = commits{
	Checkout: Option{
		Msg:         "c",
		Description: "checks out to highlighted commit",
	},
	CheckoutForce: Option{
		Msg:         "C",
		Description: "forces checkout to highlighted commit",
	},
	Reset: Option{
		Msg:         "r",
		Description: "opens menu to choose reset option for highlighted commit",
	},
	CherryPick: Option{
		Msg:         "ctrl+p",
		Description: "cherry picks highlighted commit",
	},
	Yank: Option{
		Msg:         "y",
		Description: "copies highlighted commit hash to clipboard",
	},

	Up: Option{
		Msg:         "up",
		Description: "highlights previous log",
		AltMsg:      "k",
	},
	Down: Option{
		Msg:         "down",
		Description: "highlights next log",
		AltMsg:      "j",
	},
	Bottom: Option{
		Msg:         "G",
		Description: "highlights last log",
	},

	Refresh: Refresh,
	Search:  Search,
}
