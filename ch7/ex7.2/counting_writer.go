package ex7_2

import (
	"bufio"
	"bytes"
	"io"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {

}
