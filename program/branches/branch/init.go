package branch

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"omzgit/git"
	"omzgit/roller"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Active        bool
	Current       bool
	defaultBranch string
	diff          string
	lastUpdated   string
	Roller        roller.Model

	width int
}

func EmptyInitialModel(width int, height int, name string, defaultBranch string) Model {
	return Model{
		Active:        true,
		Roller:        roller.InitialModel(getWidth(width), name),
		lastUpdated:   "",
		Current:       false,
		defaultBranch: defaultBranch,
		diff:          "",

		width: getWidth(width),
	}
}

func InitialModel(width int, name string, defaultBranch string) Model {
	return Model{
		Active:        false,
		Roller:        roller.InitialModel(getWidth(width), name[2:]),
		Current:       strings.Contains(name[:2], "*"),
		defaultBranch: defaultBranch,
		diff:          "",
		lastUpdated:   "",

		width: getWidth(width),
	}
}

func (m Model) Init() tea.Cmd {
	return m.Roller.Init()
}

func getWidth(width int) int {
	return width
}

func (m Model) getLastUpdatedDate() string {
	if !m.Active {
		return ""
	}

	originName := "origin/" + m.Roller.Name
	output, err := git.Exec("log", "-1", "--format=%cd", originName)
	if err != nil {
		return "---"
	}

	layout := "Mon Jan 2 15:04:05 2006 -0700"
	parsedDate, err := time.Parse(layout, string(output)[:len(string(output))-1])
	if err != nil {
		return "---"
	}
	return formatUnixTime(time.Now().Unix() - parsedDate.Unix())
}

func formatUnixTime(unixTime int64) string {
	fTime := unixTime / 60
	if fTime < 60 {
		return fmt.Sprint(fTime, " Minutes Ago")
	}

	fTime /= 60
	if fTime < 24 {
		return fmt.Sprint(fTime, " Hours Ago")
	}

	fTime /= 24
	if fTime < 7 {
		return fmt.Sprint(fTime, " Days Ago")
	}

	fTime /= 7
	if fTime < 52 {
		return fmt.Sprint(fTime, " Weeks Ago")
	}

	fTime /= 365
	if fTime == 1 {
		return fmt.Sprint(fTime, "Year Ago")
	}
	return fmt.Sprint(fTime, " Years Ago")
}

func (m Model) getBranchDiff() string {
	if !m.Active {
		return ""
	}

	currentRemoteBranch := "origin/" + m.Roller.Name
	if currentRemoteBranch == m.defaultBranch {
		return "Default"
	}

	branchPath := currentRemoteBranch + "...HEAD"
	output, err := git.Exec("rev-list", "--left-right", "--count", branchPath)
	if err != nil {
		return "0 | 0"
	}

	fields := strings.Fields(string(output))
	slices.Reverse(fields)
	trimmed := strings.Join(fields, "|")
	return trimmed
}
