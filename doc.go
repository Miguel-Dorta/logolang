/*
	Package logolang is a simple and thread-safe library for logging operations.


	Logger levels

	The logger object from logolang have an internal logger level that is used to determine which
	log messages should log. It will log the messages whose log level is equal or less than the
	defined level. This levels are:
		0: no log
		1: critical
		2: error
		3: info
		4: debug


	Writers

	If you want to use a different writers for logging operations, it MUST be safe for concurrent use.
	You can get a thread-safe writer for an unsafe one by wrapping it in a SafeWriter. If you're not
	sure if your writer is safe or not, you should probably wrap it.

	The custom writers MUST also be reliable for writing, because a logging operation that founds an
	error while writing will end up in panic.


	Format

	You can set the way your message is logged by setting a custom format. The format is defined by a
	string where the following sequences are given the following values:
		%YYYY%    = current year
		%MM%      = current month
		%DD%      = current day of the month
		%hh%      = current hour
		%mm%      = current minute
		%ss%      = current second
		%ns%      = current nanosecond
		%LEVEL%   = level name (DEBUG, INFO, ERROR or CRITICAL)
		%MESSAGE% = message logged

	The default format is:
		DefaultFormat  = "[%YYYY%-%MM%-%DD% %hh%:%mm%:%ss%] %LEVEL%: %MESSAGE%"


	Colors

	The standard behaviour of logolang is to print log messages in the terminal coloring the level name
	of the logger. This uses special characters that the terminal will understand as colors, but text
	files and other things could not identify it as that. For that reason, you can disable color when
	using NewLogger.
 */
package logolang
