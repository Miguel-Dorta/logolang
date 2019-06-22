// Package logolang is a simple and thread-safe library for logging operations.
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
	"os"
	"sync"
	"time"
)

const (
	// Logger levels
	LevelNoLog = iota
	LevelCritical
	LevelError
	LevelInfo
	LevelDebug

	// Level names
	nameCritical = "CRITICAL"
	nameError = "ERROR"
	nameInfo = "INFO"
	nameDebug = "DEBUG"

	// ANSI escape sequences
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

// NewLogger creates a new Logger object.
// It will write the log in the writers provided.
// If one of the writers is set to nil, it will be set to its default value.
//
// Default values:
//
// - debug: os.Stdout
//
// - info: os.Stdout
//
// - error: os.Stderr
//
// - critical: os.Stderr
func NewLogger(debug, info, error, critical io.Writer) *Logger {
	l := Logger{
		debug:    os.Stdout,
		info:     os.Stdout,
		error:    os.Stderr,
		critical: os.Stderr,
	}

	if debug != nil {
		l.debug = debug
	}
	if info != nil {
		l.info = info
	}
	if error != nil {
		l.error = error
	}
	if critical != nil {
		l.critical = critical
	}
	return &l
}

// SetLevel sets the logger level to the value given
func (l *Logger) SetLevel(level int) error {
	if level < LevelNoLog || level > LevelDebug {
		return errors.New("invalid value")
	}
	l.mutex.Lock()
	l.level = level
	l.mutex.Unlock()
	return nil
}

// Critical logs a critical message in the critical interface when logger level >= LevelCritical.
func (l *Logger) Critical(message string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.level < LevelCritical {
		return
	}
	l.log(l.critical, nameCritical, colorRed, message)
}

// Criticalf logs a critical message in the critical interface when logger level >= LevelCritical.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Criticalf(format string, v ...interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.level < LevelCritical {
		return
	}
	l.log(l.critical, nameCritical, colorRed, fmt.Sprintf(format, v...))
}

// Error logs an error message in the error interface when logger level >= LevelError.
func (l *Logger) Error(message string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.level < LevelError {
		return
	}
	l.log(l.error, nameError, colorYellow, message)
}

// Error logs an error message in the error interface when logger level >= LevelError.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.level < LevelError {
		return
	}
	l.log(l.error, nameError, colorYellow, fmt.Sprintf(format, v...))
}

// Info logs an info message in the info interface when logger level >= LevelInfo.
func (l *Logger) Info(message string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.level < LevelInfo {
		return
	}
	l.log(l.info, nameInfo, colorDefault, message)
}

// Info logs an info message in the info interface when logger level >= LevelInfo.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.level < LevelInfo {
		return
	}
	l.log(l.info, nameInfo, colorDefault, fmt.Sprintf(format, v...))
}

// Debug logs a debug message in the debug interface when logger level >= LevelDebug.
func (l *Logger) Debug(message string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.level < LevelDebug {
		return
	}
	l.log(l.debug, nameDebug, colorLightBlue, message)
}

// Debug logs a debug message in the debug interface when logger level >= LevelDebug.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.level < LevelDebug {
		return
	}
	l.log(l.debug, nameDebug, colorLightBlue, fmt.Sprintf(format, v...))
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
