// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
// Modified to close the write half of the connection
// Modified to disconnect clients after a timeout of 10s
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	if wg != nil {
		wg.Done()
	}
}

//!+
func handleConn(c *net.TCPConn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	userInput := make(chan struct{})

	go func() {
		for input.Scan() {
			userInput <- struct{}{}
			wg.Add(1)
			go echo(c, input.Text(), 5*time.Second, &wg)
		}
	}()

	for to := false; !to; {
		select {
		case <-userInput:
			// do nothing
		case <-time.After(10 * time.Second):
			to = true
		}
	}

	// NOTE: ignoring potential errors from input.Err()
	wg.Wait()
	c.CloseWrite()
}

//!-

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8000") // ignoring error - the address is hardcoded
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
