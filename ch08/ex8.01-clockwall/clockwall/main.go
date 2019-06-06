// ClockWall connects to several clock server and displays the time
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type clock struct {
	loc, addr string
	ctime     time.Time
}

// clockWatch connects to a clock server at address a
func (c *clock) clockWatch() {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	s := bufio.NewScanner(conn)
	for s.Scan() {
		c.ctime, err = time.Parse("15:04:05", s.Text())
		if err != nil {
			log.Fatal(err)
		}
	}
}

// showClockWall displays the clocks
func showClockWall(clocks []*clock) {
	for _, c := range clocks {
		fmt.Printf("%s: %s\n", c.loc, c.ctime.Format("15:04:05"))
	}
}

func main() {

	var clocks = make([]*clock, 0, 0)

	// parse input
	for _, arg := range os.Args[1:] {
		fields := strings.Split(arg, "=")
		if len(fields) != 2 {
			log.Fatalf("invalid argument %s", arg)
		}
		clocks = append(clocks, &clock{fields[0], fields[1], *new(time.Time)})
	}

	// initiate clock watchers
	for _, c := range clocks {
		log.Printf("connecting to clock @%s", c.addr)
		go c.clockWatch()
	}

	// display the clock wall
	for {
		showClockWall(clocks)
		time.Sleep(1 * time.Second)
		fmt.Println()
	}

}
