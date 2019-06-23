package logolang

import (
	"fmt"
	"strings"
	"time"
)

const (
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

	// Default format for logolang
	DefaultFormat  = "[%YYYY%-%MM%-%DD% %hh%:%mm%:%ss%] %LEVEL%: %MESSAGE%"
)

func format(format, level, msg string) string {
	now := time.Now()

	parts := strings.Split(format, escapeChar)

	var sb strings.Builder; sb.Grow(50)
	v := make([]interface{}, 0, 8)

	for _, p := range parts {
		var toAppendFormat string
		var toAppendV interface{}

		switch p {
		case sequenceYear:
			toAppendFormat = "%04d"
			toAppendV = now.Year()
		case sequenceMonth:
			toAppendFormat = "%02d"
			toAppendV = now.Month()
		case sequenceDay:
			toAppendFormat = "%02d"
			toAppendV = now.Day()
		case sequenceHour:
			toAppendFormat = "%02d"
			toAppendV = now.Hour()
		case sequenceMinute:
			toAppendFormat = "%02d"
			toAppendV = now.Minute()
		case sequenceSecond:
			toAppendFormat = "%02d"
			toAppendV = now.Second()
		case sequenceNanosecond:
			toAppendFormat = "%09d"
			toAppendV = now.Nanosecond()
		case sequenceLevel:
			toAppendFormat = "%s"
			toAppendV = level
		case sequenceMessage:
			toAppendFormat = "%s"
			toAppendV = msg
		default:
			toAppendFormat = "%s"
			toAppendV = nil
		}

		sb.WriteString(toAppendFormat)
		if toAppendV != nil {
			v = append(v, toAppendV)
		}
	}

	return fmt.Sprintf(sb.String(), v...)
}
