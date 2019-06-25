package logolang

import (
	"io"
	"sync"
)

// SafeWriter is a simple wrapper for making a writer safe for concurrent access.
type SafeWriter struct {
	W io.Writer
	mutex sync.Mutex
}

// Write follows the implementation of the io.Writer interface.
func (sf *SafeWriter) Write(p []byte) (n int, err error) {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()
	return sf.W.Write(p)
}
