// logolang is a simple and thread-safe package for logging operations.
//
// It consists in a Logger object where you can configure a writer for each log level.
// There are 5 of those levels:
//
// - 0: no log
//
// - 1: critical
//
// - 2: error
//
// - 3: info
//
// - 4: debug
package logolang

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"time"
)

const (
	colorDefault = "\x1b[39m"
	colorLightBlue = "\x1b[94m"
	colorRed = "\x1b[31m"
	colorYellow = "\x1b[33m"
)

// Logger is the type used by logolang to perform logging operations.
// It have 4 writers, one for each logging level except 0 (no log).
// Every logging operation results in panic if there's an error when writing to one of those interfaces.
type Logger struct {
	level int
	debug, info, error, critical io.Writer
	mutex sync.Mutex
}

// NewLogger creates a new Logger object
func NewLogger(debug, info, error, critical io.Writer) (*Logger, error) {
	if debug == nil || info == nil || error == nil || critical == nil {
		return nil, errors.New("not all interfaces are defined")
	}
	return &Logger{
		debug:    debug,
		info:     info,
		error:    error,
		critical: critical,
	}, nil
}

// SetLevel sets the logger level to the value given
func (l *Logger) SetLevel(level int) error {
	if level < 0 || level > 4 {
		return errors.New("invalid value")
	}
	l.mutex.Lock()
	l.level = level
	l.mutex.Unlock()
	return nil
}

// Critical logs a critical message in the critical interface when the logger level is 1 or greater
func (l *Logger) Critical(message string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.level < 1 {
		return
	}
	l.log(l.critical, "CRITICAL", colorRed, message)
}

// Error logs an error message in the error interface when the logger level is 2 or greater
func (l *Logger) Error(message string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.level < 2 {
		return
	}
	l.log(l.error, "ERROR", colorYellow, message)
}

// Info logs an info message in the info interface when the logger level is 3 or greater
func (l *Logger) Info(message string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.level < 3 {
		return
	}
	l.log(l.info, "INFO", colorDefault, message)
}

// Debug logs a debug message in the debug interface when the logger level is 4 or greater
func (l *Logger) Debug(message string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.level < 4 {
		return
	}
	l.log(l.debug, "DEBUG", colorLightBlue, message)
}

// log is the internal function for logging messages
func (l *Logger) log(w io.Writer, levelName, levelColor, message string) {
	now := time.Now()
	_, err := fmt.Fprintf(w, "[%04d-%02d-%02d %02d:%02d:%02d] %s%s%s: %s\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(),
		levelColor, levelName, colorDefault, message,
	)
	if err != nil {
		panic(err)
	}
}
