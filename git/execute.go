package git

import (
	"os/exec"
	"strings"
)

func Exec(args ...string) (string, error) {
	cmd := exec.Command("git", args...)

	output, err := cmd.CombinedOutput()

	return strings.TrimSpace(string(output)), err
}
