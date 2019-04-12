package main

import (
	"bytes"
	"fmt"
	"io"
)

type CountWrites struct {
	count int64
	w     io.Writer
}

func (c *CountWrites) Write(p []byte) (int, error) {
	count, err := c.w.Write(p)
	c.count += int64(count)
	return count, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var r CountWrites
	r.w = w
	return &r, &r.count
}

func main() {

	buf := bytes.NewBuffer([]byte{})
	cbuf, count := CountingWriter(buf)

	fmt.Printf("\nInitial content: %#v\n", buf.String())
	s := "12345678"
	fmt.Printf("Writing: \"%s\"\n", s)
	cbuf.Write([]byte(s))
	fmt.Printf("Counter: %d\n", *count)
	fmt.Printf("Content: %#v\n", buf.String())
	fmt.Printf("Writing: \"%s\"\n", s)
	cbuf.Write([]byte(s))
	fmt.Printf("Counter: %d\n", *count)
	fmt.Printf("Content: %#v\n", buf.String())

}
