package logolang

import (
	"fmt"
	"io"
)

const (
	// Logger levels
	LevelNoLog    = iota
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

// log is the internal function for logging messages
func (l *Logger) log(w io.Writer, levelName, levelColor, message string) {
	if l.Color {
		levelName = levelColor + levelName + colorDefault
	}

	_, err := fmt.Fprintln(w, l.Formatter(levelName, message))
	if err != nil {
		panic(err)
	}
}

// Critical logs a critical message in the critical interface when logger Level >= LevelCritical.
func (l *Logger) Critical(message string) {
	if l.Level < LevelCritical {
		return
	}
	l.log(l.critical, nameCritical, colorRed, message)
}

// Criticalf logs a critical message in the critical interface when logger Level >= LevelCritical.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Criticalf(format string, v ...interface{}) {
	if l.Level < LevelCritical {
		return
	}
	l.log(l.critical, nameCritical, colorRed, fmt.Sprintf(format, v...))
}

// Debug logs a debug message in the debug interface when logger Level >= LevelDebug.
func (l *Logger) Debug(message string) {
	if l.Level < LevelDebug {
		return
	}
	l.log(l.debug, nameDebug, colorLightBlue, message)
}

// Debug logs a debug message in the debug interface when logger Level >= LevelDebug.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.Level < LevelDebug {
		return
	}
	l.log(l.debug, nameDebug, colorLightBlue, fmt.Sprintf(format, v...))
}

// Error logs an error message in the error interface when logger Level >= LevelError.
func (l *Logger) Error(message string) {
	if l.Level < LevelError {
		return
	}
	l.log(l.error, nameError, colorYellow, message)
}

// Error logs an error message in the error interface when logger Level >= LevelError.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.Level < LevelError {
		return
	}
	l.log(l.error, nameError, colorYellow, fmt.Sprintf(format, v...))
}

// Info logs an info message in the info interface when logger Level >= LevelInfo.
func (l *Logger) Info(message string) {
	if l.Level < LevelInfo {
		return
	}
	l.log(l.info, nameInfo, colorDefault, message)
}

// Info logs an info message in the info interface when logger Level >= LevelInfo.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.Level < LevelInfo {
		return
	}
	l.log(l.info, nameInfo, colorDefault, fmt.Sprintf(format, v...))
}
