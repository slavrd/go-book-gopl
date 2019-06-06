// ClockWall connects to several clock server and displays the time
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"text/tabwriter"
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
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 12, 2, ' ', 0)
	for _, c := range clocks {
		fmt.Fprintf(tw, "%v\t", c.loc)
	}
	fmt.Fprintf(tw, "\n")
	for _, c := range clocks {
		for i := 0; i < len(c.loc); i++ {
			fmt.Fprint(tw, "-")
		}
		fmt.Fprintf(tw, "\t")
	}
	fmt.Fprintf(tw, "\n")
	for _, c := range clocks {
		fmt.Fprintf(tw, "%v\t", c.ctime.Format("15:04:05"))
	}
	tw.Flush()

	for {
		fmt.Fprintf(tw, "\r")
		for _, c := range clocks {
			fmt.Fprintf(tw, "%v\t", c.ctime.Format("15:04:05"))
		}
		fmt.Fprintf(tw, "\t")
		tw.Flush()
		time.Sleep(100 * time.Millisecond)
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
		log.Printf("connecting to clock %s@%s", c.loc, c.addr)
		go c.clockWatch()
	}

	// display the clock wall
	fmt.Println("")
	showClockWall(clocks)
}
