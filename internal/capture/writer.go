package capture

import (
	"io"
	"sync"
)

type Writer struct {
	inner io.Writer
	mu    sync.Mutex
	bytes []byte
	done  bool
}

func NewWriter(writer io.Writer) *Writer {
	return &Writer{
		inner: writer,
	}
}

func (w *Writer) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.bytes = append(w.bytes, p...)

	return w.inner.Write(p)

}

func (w *Writer) Read(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.done {
		return 0, io.EOF
	}

	for _, b := range w.bytes {
		p = append(p, b)
	}
	w.done = true

	return len(w.bytes), nil
}
