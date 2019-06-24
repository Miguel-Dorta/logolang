package logolang

import (
	"io"
	"sync"
)

type SafeWriter struct {
	w io.Writer
	mutex sync.Mutex
}

func NewSafeWriter(w io.Writer) io.Writer {
	return &SafeWriter{w: w}
}

func (sf *SafeWriter) Write(p []byte) (n int, err error) {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()
	return sf.w.Write(p)
}
