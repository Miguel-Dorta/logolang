# logolang
Package logolang is a simple and thread-safe library for logging operations.

## Example
```go
package main

import (
	log "github.com/Miguel-Dorta/logolang"
)

func main() {
	log.DefaultLogger.Level = log.LevelDebug

	log.Debug("debug test")
	log.Info("info test")
	log.Error("error test")
	log.Critical("critical test")
}
```

Output

![Example of logolang](https://i.nth.sh/media/4mM4w8KV46/xdiR4R4tiz.png)


## Documentation

### Logger levels

The logger object from logolang have an internal logger level that is used to determine which
log messages should log. It will log the messages whose log level is equal or less than the
defined level. This levels are:
* 0: no log
* 1: critical
* 2: error
* 3: info
* 4: debug


### Formatter

The standard behaviour of logolang is to print log message with the following format:
* [YYYY-MM-DD hh:mm:ss.nssssssss] LEVEL: MESSAGE

You may not want this, so you can set a custom formatter function. This function will receive the
name of the level logged (levelName) and the message to display (msg). This function will return
just a string with the formatted text.


### Colors

The standard behaviour of logolang is to print log messages in the terminal coloring the level name
of the logger. This uses special characters that the terminal will understand as colors, but text
files and other things could not identify it as that. For that reason, you can disable it.


### Writers

If you want to use a different writers for logging operations, it MUST be safe for concurrent use.
You can get a thread-safe writer for an unsafe one by wrapping it in a SafeWriter. If you're not
sure if your writer is safe or not, you should probably wrap it.

The custom writers MUST also be reliable for writing, because a logging operation that founds an
error while writing will end up in panic.


You can read the entire documentation in https://godoc.org/github.com/Miguel-Dorta/logolang

