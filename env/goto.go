package env

type gitGoto struct {
	Actions  Option
	Branches Option
	Commits  Option
	Files    Option
	Issues   Option
	Prs      Option
	Top      Option
}

var Goto = gitGoto{
	Actions: Option{
		Msg:         "a",
		Description: "opens actions page in browser",
	},
	Branches: Option{
		Msg:         "b",
		Description: "switches to Branches page",
	},
	Files: Option{
		Msg:         "f",
		Description: "switches to Files page",
	},
	Commits: Option{
		Msg:         "c",
		Description: "switches to Commits page",
	},
	Issues: Option{
		Msg:         "i",
		Description: "opens create an issue page in browser",
	},
	Prs: Option{
		Msg:         "p",
		Description: "opens create a pull request page in browser",
	},
	Top: Option{
		Msg:         "g",
		Description: "goes to the top of the list",
	},
}
