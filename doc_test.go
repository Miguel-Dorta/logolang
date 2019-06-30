package logolang_test

import (
	"fmt"
	"github.com/Miguel-Dorta/logolang"
	"os"
)

func ExampleLogger_color() {
	log := logolang.NewLogger()
	log.Level = logolang.LevelDebug

	// Printed with color
	log.Debug("debug test")
	log.Info("info test")
	log.Error("error test")
	log.Critical("critical test")

	log.Color = false

	// Printed with no color
	log.Debug("debug test")
	log.Info("info test")
	log.Error("error test")
	log.Critical("critical test")
}

func ExampleLogger_formatter() {
	log := logolang.NewLogger()
	log.Level = logolang.LevelDebug

	// Set Formatter to be "LEVEL: MESSAGE"
	log.Formatter = func(levelName, msg string) string {
		return fmt.Sprintf("%s: %s", levelName, msg)
	}

	log.Debug("debug test")
	log.Info("info test")
	log.Error("error test")
	log.Critical("critical test")
}

func ExampleLogger_level() {
	log := logolang.NewLogger()
	log.Level = logolang.LevelInfo // Print info or lower

	log.Debug("debug test")       // Will not be printed
	log.Info("info test")         // Will be printed
	log.Error("error test")       // Will be printed
	log.Critical("critical test") // Will be printed
}

func ExampleNewLoggerWriters() {
	// Create output files
	outFile, err := os.Create("/tmp/out.log")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	errFile, err := os.Create("/tmp/err.log")
	if err != nil {
		panic(err)
	}
	defer errFile.Close()

	// Create safe writers
	outSafe := &logolang.SafeWriter{W: outFile}
	errSafe := &logolang.SafeWriter{W: errFile}

	// Assign safe writers to new Logger
	log := logolang.NewLoggerWriters(outSafe, outSafe, errSafe, errSafe)
	log.Color = false    // Disable color (read documentation)
	log.Level = logolang.LevelDebug

	log.Debug("debug test")
	log.Info("info test")
	log.Error("error test")
	log.Critical("critical test")
}

func ExampleSafeWriter() {
	// Create output files
	outFile, err := os.Create("/tmp/out.log")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	errFile, err := os.Create("/tmp/err.log")
	if err != nil {
		panic(err)
	}
	defer errFile.Close()

	// Create safe writers
	outSafe := &logolang.SafeWriter{W: outFile}
	errSafe := &logolang.SafeWriter{W: errFile}

	// Assign safe writers to new Logger
	log := logolang.NewLoggerWriters(outSafe, outSafe, errSafe, errSafe)
	log.Color = false    // Disable color (read documentation)
	log.Level = logolang.LevelDebug

	log.Debug("debug test")
	log.Info("info test")
	log.Error("error test")
	log.Critical("critical test")
}
