package env

type program struct {
	Esc         Option
	Fetch       Option
	Goto        Option
	Help        Option
	Pull        Option
	PullOptions Option
	Push        Option
	PushForce   Option
	CtrlC       Option
	Quit        Option
}

var Program = program{
	Fetch: Option{
		Msg:         "f",
		Description: "fetches from remote",
	},
	Goto: Option{
		Msg:         "g",
		Description: "goto/git",
	},
	Help: Option{
		Msg:         "h",
		AltMsg:      "?",
		Description: "shows current available actions",
	},
	Pull: Option{
		Msg:         "l",
		Description: "pulls from remote",
	},
	PullOptions: Option{
		Msg:         "L",
		Description: "more options for pulling",
	},
	Push: Option{
		Msg:         "p",
		Description: "pushes commits to remote",
	},
	PushForce: Option{
		Msg:         "P",
		Description: "force pushes commits to remote",
	},
	CtrlC: Quit.CtrlC,
	Quit:  Quit.Quit,
}
