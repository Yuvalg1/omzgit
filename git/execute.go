package git

import "os/exec"

func Exec(args ...string) bool {
	cmd := exec.Command("git", args...)

	_, err := cmd.Output()

	return err == nil
}
