package branch

import (
	"fmt"
	"os/exec"
	"slices"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Active        bool
	Current       bool
	defaultBranch string
	diff          string
	lastUpdated   string
	Name          string

	width  int
	height int
}

func InitialModel(width int, height int, name string, defaultBranch string, empty bool) Model {
	if empty {
		return Model{
			Active:        true,
			Name:          name,
			lastUpdated:   "",
			Current:       false,
			defaultBranch: defaultBranch,
			diff:          "",

			width:  getWidth(width),
			height: getHeight(height),
		}
	}

	return Model{
		Active:        false,
		Name:          name[2:],
		Current:       strings.Contains(name[:2], "*"),
		defaultBranch: defaultBranch,
		diff:          "",
		lastUpdated:   "",

		width:  getWidth(width),
		height: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func getWidth(width int) int {
	return width
}

func getHeight(height int) int {
	return 1
}

func (m Model) getLastUpdatedDate() string {
	if !m.Active {
		return ""
	}

	originName := "origin/" + m.Name
	cmd := exec.Command("git", "log", "-1", "--format=%cd", originName)

	stdout, err := cmd.Output()
	if err != nil {
		return "---"
	}

	layout := "Mon Jan 2 15:04:05 2006 -0700"
	parsedDate, err := time.Parse(layout, string(stdout)[:len(string(stdout))-1])
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

	currentRemoteBranch := "origin/" + m.Name
	if currentRemoteBranch == m.defaultBranch {
		return "Default"
	}

	branchPath := currentRemoteBranch + "...HEAD"
	cmd := exec.Command("git", "rev-list", "--left-right", "--count", branchPath)

	stdout, err := cmd.Output()
	if err != nil {
		return "0 | 0"
	}

	fields := strings.Fields(string(stdout))
	slices.Reverse(fields)
	trimmed := strings.Join(fields, "|")
	return trimmed
}
