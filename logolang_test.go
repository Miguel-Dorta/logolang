package logolang_test

import (
	"fmt"
	"github.com/Miguel-Dorta/logolang"
	"strings"
	"testing"
)

type LoggerTest struct {
	Log *logolang.Logger
	Builder *strings.Builder
	SafeW *logolang.SafeWriter
}

func NewLoggerTest() LoggerTest {
	b := &strings.Builder{}
	sw := &logolang.SafeWriter{W: b}
	return LoggerTest{
		Log: logolang.NewLoggerWriters(sw, sw, sw, sw),
		Builder: b,
		SafeW: sw,
	}
}

func (lt *LoggerTest) Equals(s string) bool {
	return lt.Builder.String() == s
}

func TestLogger_Color(t *testing.T) {
	log := logolang.NewLogger()
	log.Level = logolang.LevelDebug

	fmt.Println("Color:")
	log.Debug("debug test")
	log.Info("info test")
	log.Error("error test")
	log.Critical("critical test")

	log.Color = false
	fmt.Println("\nNo color:")
	log.Debug("debug test")
	log.Info("info test")
	log.Error("error test")
	log.Critical("critical test")
}

func TestLogger_Formatter(t *testing.T) {
	lt := NewLoggerTest()
	lt.Log.Color = false
	lt.Log.Formatter = func(levelName, msg string) string {
		return levelName + " " + msg
	}
	lt.Log.Level = logolang.LevelDebug

	lt.Log.Debug("debug test")
	lt.Log.Info("info test")
	lt.Log.Error("error test")
	lt.Log.Critical("critical test")

	expectedResult :=
			"DEBUG debug test\n" +
			"INFO info test\n" +
			"ERROR error test\n" +
			"CRITICAL critical test\n"

	if !lt.Equals(expectedResult) {
		t.Fatalf("Formatter test failed.\nExpected result: \"%s\"\nObtained result: \"%s\"", expectedResult, lt.Builder.String())
	}
}

func TestLogger_Level(t *testing.T) {
	lt := NewLoggerTest()
	lt.Log.Color = false
	lt.Log.Formatter = func(levelName, msg string) string {
		return levelName + " " + msg
	}

	// Debug
	if _, err := fmt.Fprintln(lt.SafeW, "LevelDebug"); err != nil {
		t.Fatalf("Cannot write in lt.SafeW: %s", err.Error())
	}
	lt.Log.Level = logolang.LevelDebug
	lt.Log.Debug("debug test")
	lt.Log.Info("info test")
	lt.Log.Error("error test")
	lt.Log.Critical("critical test")

	// Info
	if _, err := fmt.Fprintln(lt.SafeW, "LevelInfo"); err != nil {
		t.Fatalf("Cannot write in lt.SafeW: %s", err.Error())
	}
	lt.Log.Level = logolang.LevelInfo
	lt.Log.Debug("debug test")
	lt.Log.Info("info test")
	lt.Log.Error("error test")
	lt.Log.Critical("critical test")

	// Error
	if _, err := fmt.Fprintln(lt.SafeW, "LevelError"); err != nil {
		t.Fatalf("Cannot write in lt.SafeW: %s", err.Error())
	}
	lt.Log.Level = logolang.LevelError
	lt.Log.Debug("debug test")
	lt.Log.Info("info test")
	lt.Log.Error("error test")
	lt.Log.Critical("critical test")

	// Critical
	if _, err := fmt.Fprintln(lt.SafeW, "LevelCritical"); err != nil {
		t.Fatalf("Cannot write in lt.SafeW: %s", err.Error())
	}
	lt.Log.Level = logolang.LevelCritical
	lt.Log.Debug("debug test")
	lt.Log.Info("info test")
	lt.Log.Error("error test")
	lt.Log.Critical("critical test")

	// No Log
	if _, err := fmt.Fprintln(lt.SafeW, "LevelNoLog"); err != nil {
		t.Fatalf("Cannot write in lt.SafeW: %s", err.Error())
	}
	lt.Log.Level = logolang.LevelNoLog
	lt.Log.Debug("debug test")
	lt.Log.Info("info test")
	lt.Log.Error("error test")
	lt.Log.Critical("critical test")

	expectedResult :=
			"LevelDebug\n" +
			"DEBUG debug test\n" +
			"INFO info test\n" +
			"ERROR error test\n" +
			"CRITICAL critical test\n" +
			"LevelInfo\n" +
			"INFO info test\n" +
			"ERROR error test\n" +
			"CRITICAL critical test\n" +
			"LevelError\n" +
			"ERROR error test\n" +
			"CRITICAL critical test\n" +
			"LevelCritical\n" +
			"CRITICAL critical test\n" +
			"LevelNoLog\n"

	if !lt.Equals(expectedResult) {
		t.Fatalf("Formatter test failed.\nExpected result: \"%s\"\nObtained result: \"%s\"", expectedResult, lt.Builder.String())
	}
}
