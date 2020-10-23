package main

import (
	"bytes"
	"fmt"
)

func main() {
	var bwc WriterCloser = NewBufferedWriteCloser()
	bwc.Write([]byte("Hello, this is my composer interface  yay."))
	bwc.Flush()
}

func NewBufferedWriteCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: &bytes.Buffer{},
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type Flusher interface {
	Flush() error
}

type WriterCloser interface {
	Writer
	Flusher
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Flush() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}
