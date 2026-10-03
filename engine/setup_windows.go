//go:build windows

package engine

import (
	"os"

	"golang.org/x/sys/windows"
)

func setupTerminal() func() {
	handle := windows.Handle(os.Stdout.Fd())
	var origMode uint32
	_ = windows.GetConsoleMode(handle, &origMode)
	_ = windows.SetConsoleMode(handle, origMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)

	origCP, _ := windows.GetConsoleCP()
	origOutputCP, _ := windows.GetConsoleOutputCP()
	_ = windows.SetConsoleCP(65001)
	_ = windows.SetConsoleOutputCP(65001)

	return func() {
		_ = windows.SetConsoleMode(handle, origMode)
		_ = windows.SetConsoleCP(origCP)
		_ = windows.SetConsoleOutputCP(origOutputCP)
	}
}
