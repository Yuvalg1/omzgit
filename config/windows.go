//go:build windows

package config

import (
	"os"

	"golang.org/x/sys/windows"
)

func RestoreConsole() {
	windows.SetConsoleOutputCP(65001)
	windows.SetConsoleCP(65001)
	h := windows.Handle(os.Stdout.Fd())
	var mode uint32
	windows.GetConsoleMode(h, &mode)
	windows.SetConsoleMode(h, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
