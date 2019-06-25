package logolang

import (
	"errors"
	"fmt"
	"io"
	"os"
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

// Logger is the type used by logolang to perform logging operations. .
type Logger struct {
	color bool
	format *formatter
	level int
	debug, info, error, critical io.Writer
}

// NewLoggerStandard creates a new logger with the default values.
//
// The default values are:
//		color:    true
//		format:   DefaultFormat
//		level:    LevelError
//		debug:    os.Stdout
//		info:     os.Stdout
//		error:    os.Stderr
//		critical: os.Stderr
func NewLoggerStandard() *Logger {
	l, _ := NewLogger(true, "", LevelError, nil, nil, nil, nil)
	return l
}

// NewLogger creates a new Logger object.
// It's the function that lets you define every aspect of the new logger.
// It lets you define whether you want the output to be printed with colors,
// the format you want the messages to be logged, the logger level,
// and the writers where you want to print your log.
//
// If the format provided is a empty string, it will be set to DefaultFormat.
// If any of the writers provided is nil, it will be assigned to its default value.
//
// Default values:
//		debug:    os.Stdout
//		info:     os.Stdout
//		error:    os.Stderr
//		critical: os.Stderr
func NewLogger(color bool, format string, level int, debug, info, error, critical io.Writer) (*Logger, error) {
	// Check level
	if level < LevelNoLog || level > LevelDebug {
		return nil, errors.New("invalid level")
	}

	// Check format
	if format == "" {
		format = DefaultFormat
	}

	// Check writers
	stdout := &SafeWriter{W: os.Stdout}
	stderr := &SafeWriter{W: os.Stderr}

	if debug == nil {
		debug = stdout
	}
	if info == nil {
		info = stdout
	}
	if error == nil {
		error = stderr
	}
	if critical == nil {
		critical = stderr
	}

	// Create and return logger
	return &Logger{
		color: color,
		format: newFormatter(format),
		level: level,
		debug: debug,
		info: info,
		error: error,
		critical: critical,
	}, nil
}

// Critical logs a critical message in the critical interface when logger level >= LevelCritical.
func (l *Logger) Critical(message string) {
	if l.level < LevelCritical {
		return
	}
	l.log(l.critical, nameCritical, colorRed, message)
}

// Criticalf logs a critical message in the critical interface when logger level >= LevelCritical.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Criticalf(format string, v ...interface{}) {
	if l.level < LevelCritical {
		return
	}
	l.log(l.critical, nameCritical, colorRed, fmt.Sprintf(format, v...))
}

// Error logs an error message in the error interface when logger level >= LevelError.
func (l *Logger) Error(message string) {
	if l.level < LevelError {
		return
	}
	l.log(l.error, nameError, colorYellow, message)
}

// Error logs an error message in the error interface when logger level >= LevelError.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level < LevelError {
		return
	}
	l.log(l.error, nameError, colorYellow, fmt.Sprintf(format, v...))
}

// Info logs an info message in the info interface when logger level >= LevelInfo.
func (l *Logger) Info(message string) {
	if l.level < LevelInfo {
		return
	}
	l.log(l.info, nameInfo, colorDefault, message)
}

// Info logs an info message in the info interface when logger level >= LevelInfo.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level < LevelInfo {
		return
	}
	l.log(l.info, nameInfo, colorDefault, fmt.Sprintf(format, v...))
}

// Debug logs a debug message in the debug interface when logger level >= LevelDebug.
func (l *Logger) Debug(message string) {
	if l.level < LevelDebug {
		return
	}
	l.log(l.debug, nameDebug, colorLightBlue, message)
}

// Debug logs a debug message in the debug interface when logger level >= LevelDebug.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level < LevelDebug {
		return
	}
	l.log(l.debug, nameDebug, colorLightBlue, fmt.Sprintf(format, v...))
}

// log is the internal function for logging messages
func (l *Logger) log(w io.Writer, levelName, levelColor, message string) {
	if l.color {
		levelName = levelColor + levelName + colorDefault
	}
	formattedMsg := l.format.format(levelName, message)

	_, err := io.WriteString(w, formattedMsg)
	if err != nil {
		panic(err)
	}
}
