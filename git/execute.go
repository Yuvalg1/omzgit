package git

import (
	"bytes"
	"os"
	"os/exec"
)

func Exec(args ...string) (string, error) {
	var out bytes.Buffer

	cmd := exec.Command("git", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()

	return out.String(), err
}

func ExecNoOutput(args ...string) {
	cmd := exec.Command("git", args...)
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil

	cmd.Run()
}
