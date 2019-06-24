package logolang

import (
	"fmt"
	"strings"
	"time"
)

const (
	// Default format for logolang
	DefaultFormat  = "[%YYYY%-%MM%-%DD% %hh%:%mm%:%ss%] %LEVEL%: %MESSAGE%"

	// CheatSheet for parsing format
	escapeChar         = "%"
	sequenceYear       = "YYYY"
	sequenceMonth      = "MM"
	sequenceDay        = "DD"
	sequenceHour       = "hh"
	sequenceMinute     = "mm"
	sequenceSecond     = "ss"
	sequenceNanosecond = "ns"
	sequenceLevel      = "LEVEL"
	sequenceMessage    = "MESSAGE"
)

type formatter struct {
	formattedString string
	sequences []string
}

func newFormatter(format string) *formatter {
	parts := strings.Split(format, escapeChar)

	var sb strings.Builder; sb.Grow(128)
	sequences := make([]string, 0, 16)

	for _, p := range parts {
		var toAppendBuilder, toAppendSequences string

		switch p {
		case sequenceYear:
			toAppendBuilder = "%04d"
			toAppendSequences = sequenceYear
		case sequenceMonth:
			toAppendBuilder = "%02d"
			toAppendSequences = sequenceMonth
		case sequenceDay:
			toAppendBuilder = "%02d"
			toAppendSequences = sequenceDay
		case sequenceHour:
			toAppendBuilder = "%02d"
			toAppendSequences = sequenceHour
		case sequenceMinute:
			toAppendBuilder = "%02d"
			toAppendSequences = sequenceMinute
		case sequenceSecond:
			toAppendBuilder = "%02d"
			toAppendSequences = sequenceSecond
		case sequenceNanosecond:
			toAppendBuilder = "%09d"
			toAppendSequences = sequenceNanosecond
		case sequenceLevel:
			toAppendBuilder = "%s"
			toAppendSequences = sequenceLevel
		case sequenceMessage:
			toAppendBuilder = "%s"
			toAppendSequences = sequenceMessage
		default:
			toAppendBuilder = p
		}

		sb.WriteString(toAppendBuilder)
		if toAppendSequences != "" {
			sequences = append(sequences, toAppendSequences)
		}
	}

	return &formatter{
		formattedString: sb.String(),
		sequences: sequences,
	}
}

func (f *formatter) format(level, msg string) string {
	now := time.Now()

	v := make([]interface{}, len(f.sequences))
	for i, seq := range f.sequences {
		switch seq {
		case sequenceYear:
			v[i] = now.Year()
		case sequenceMonth:
			v[i] = now.Month()
		case sequenceDay:
			v[i] = now.Day()
		case sequenceHour:
			v[i] = now.Hour()
		case sequenceMinute:
			v[i] = now.Minute()
		case sequenceSecond:
			v[i] = now.Second()
		case sequenceNanosecond:
			v[i] = now.Nanosecond()
		case sequenceLevel:
			v[i] = level
		case sequenceMessage:
			v[i] = msg
		}
	}

	return fmt.Sprintf(f.formattedString, v...)
}
