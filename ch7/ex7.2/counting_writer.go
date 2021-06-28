package ex7_2

import (
	"io"
)

type writeCounter struct {
	w io.Writer
	c int64
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.w.Write(p)
	wc.c += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wc := writeCounter{w, 0}

	return &wc, &(wc.c)

}
