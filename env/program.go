package env

type program struct {
	Esc        Option
	Fetch      Option
	Goto       Option
	Pull       Option
	RebasePull Option
	Push       Option
	PushForce  Option
	CtrlC      Option
	Quit       Option
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
	Pull: Option{
		Msg:         "l",
		Description: "pulls from remote",
	},
	RebasePull: Option{
		Msg:         "L",
		Description: "rebase pulls from remote",
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
