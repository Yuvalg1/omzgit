package git

import (
	"os/exec"
	"strings"
)

func Exec(args ...string) bool {
	cmd := exec.Command("git", args...)

	_, err := cmd.Output()

	return err == nil
}

func GetExec(args ...string) string {
	cmd := exec.Command("git", args...)

	stdout, err := cmd.Output()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(stdout))
}
