# logolang
Package logolang is a simple and thread-safe library for logging operations.

It consists in a Logger object where you can configure a writer for each log level. There are 5 of those levels:
- 0: no log
- 1: critical
- 2: error
- 3: info
- 4: debug

Example:
```go
package main

import (
	"github.com/Miguel-Dorta/logolang"
	"os"
)

func main() {
	log := logolang.NewLogger(nil, nil, nil, nil)
	if err := log.SetLevel(4); err != nil {
		panic(err)
	}
	
	log.Debug("debug test")
	log.Info("info test")
	log.Error("error test")
	log.Critical("critical test")
}
```

![Example of logolang](https://i.nth.sh/media/4mM4w8KV46/nU17GQ50q1.png)
