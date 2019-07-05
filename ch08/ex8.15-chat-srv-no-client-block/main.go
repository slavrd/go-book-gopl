// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster
type client struct {
	channel chan string // an outgoing message channel
	name    string      // client string identification
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)     // all incoming client messages
	clients  = make(map[client]bool) // all connected clients
)

func broadcaster() {
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				select {
				case cli.channel <- msg:
				default:
					// skip message if client channel is busy
				}
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.channel)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	defer conn.Close()

	cli := client{make(chan string), conn.RemoteAddr().String()}
	go clientWriter(conn, cli.channel)

	active := make(chan struct{})
	go clientTimeOut(conn, active)

	input := bufio.NewScanner(conn)

	cli.channel <- "Enter your nickname:"
	if input.Scan() {
		active <- struct{}{}
		cli.name = input.Text()
	} else {
		return
	}

	cli.channel <- "You are " + cli.name
	clilist := "Connected clients:\n"
	for ccli := range clients {
		clilist += ccli.name + "\n"
	}
	cli.channel <- clilist
	messages <- cli.name + " has arrived"
	entering <- cli

	for input.Scan() {
		active <- struct{}{}
		messages <- cli.name + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- cli.name + " has left"
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func clientTimeOut(conn net.Conn, ch <-chan struct{}) {
loop:
	for {
		select {
		case <-ch:
		case <-time.After(1 * time.Minute):
			conn.Close()
			break loop
		}
	}

}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
