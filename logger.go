package logolang

import (
	"fmt"
	"io"
	"os"
	"time"
)

// DefaultFormatter is the default function for formatting the logging messages.
// It will format them as "[YYYY-MM-DD hh:mm:ss.nssssssss] LEVEL: MESSAGE"
var DefaultFormatter = func(levelName, msg string) string {
	now := time.Now()
	return fmt.Sprintf(
		"[%04d-%02d-%02d %02d:%02d:%02d.%09d] %s: %s",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(), now.Nanosecond(),
		levelName, msg,
	)
}

// Logger is the type used by logolang to perform logging operations.
// This type is NOT thread-safe. Do not modify its content concurrently.
type Logger struct {
	// Will the logger print the level name with color?
	Color                        bool

	// Function for formatting the output.
	Formatter                    func(levelName, msg string) string

	// Logger level. It will log the levels with equal or less value.
	Level                        int

	// Writers for every level.
	debug, info, error, critical io.Writer
}

// NewLogger creates a new logger with the default values.
//
// The default values are:
//		Color:     true
//		Formatter: DefaultFormatter
//		Level:     LevelError
//
// Debug(f) and Info(f) will log to the os.Stdout interface.
// Error(f) and Critical(f) will log to the os.Stderr interface.
func NewLogger() *Logger {
	stdout := &SafeWriter{W: os.Stdout}
	stderr := &SafeWriter{W: os.Stderr}
	return &Logger{
		Color:     true,
		Formatter: DefaultFormatter,
		Level:     LevelError,
		debug:     stdout,
		info:      stdout,
		error:     stderr,
		critical:  stderr,
	}
}

// NewLoggerWriters creates a new Logger object.
// This function will let you define the writer for each log level.
// If any of the writers is set to nil, it will be assigned to its default value.
// It will use standard values of NewLogger.
func NewLoggerWriters(debug, info, error, critical io.Writer) *Logger {
	l := NewLogger()

	// Check writers
	if debug == nil {
		debug = l.debug
	}
	if info == nil {
		info = l.info
	}
	if error == nil {
		error = l.error
	}
	if critical == nil {
		critical = l.critical
	}

	l.debug = debug
	l.info = info
	l.error = error
	l.critical = critical

	return l
}

