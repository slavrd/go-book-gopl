package main

import (
	"bufio"
	"bytes"
	"fmt"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

//!-bytecounter

//!+wordcounter

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	rd := bytes.NewReader(p)
	s := bufio.NewScanner(rd)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		*c++
	}

	return len(p), nil
}

//!-wordcounter

//!+linecounter

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	rd := bytes.NewReader(p)
	s := bufio.NewScanner(rd)
	for s.Scan() {
		*l++
	}

	return len(p), nil
}

//!-linecounter

func main() {

	var wc WordCounter
	bc, _ := fmt.Fprintf(&wc, "zero %s", "one two three four")
	fmt.Printf("Word count: %d\n", wc)
	fmt.Printf("Bytes count: %d\n", bc)

	var lc WordCounter
	bc, _ = fmt.Fprintf(&lc, "zero \n%s", "\none\ntwo\nthree\nfour\nfive")
	fmt.Printf("Line count: %d\n", lc)
	fmt.Printf("Bytes count: %d\n", bc)
}
