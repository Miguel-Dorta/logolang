package logolang

import (
	"golang.org/x/sys/windows"
	"os"
)

func init() {
	enableAnsiColor()
}

// enableAnsiColor enables virtual terminal processing in the console mode of the stdout and stderr
func enableAnsiColor() {
	// Get stdout & stderr references
	stdout := windows.Handle(os.Stdout.Fd())
	stderr := windows.Handle(os.Stderr.Fd())

	// Get their original modes
	var stdoutMode, stderrMode uint32
	_ = windows.GetConsoleMode(stdout, &stdoutMode)
	_ = windows.GetConsoleMode(stderr, &stderrMode)

	// Modify their modes to enable virtual terminal processing.
	// See: https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences
	_ = windows.SetConsoleMode(stdout, stdoutMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	_ = windows.SetConsoleMode(stderr, stderrMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
