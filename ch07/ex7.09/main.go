// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"text/template"
	"time"
)

//!+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 3 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!-main

//!+sorting

// mapping of compare functions to strings
var sFMap = map[string]func(x, y *Track) int{
	"byyear":   cmpByYear,
	"bylength": cmpByLength,
	"byalbum":  cmpByAlbum,
	"byartist": cmpByArtist,
	"bytitle":  cmpByTitle,
}

// stack to keep sort columns order. Higher priority are at the top.
// values must be keys in the sFMap map.
var sStack = make([]string, 0)

type CascadeSort struct {
	data []*Track
}

func (c CascadeSort) Len() int {
	return len(c.data)
}

func (c CascadeSort) Swap(i, j int) {
	c.data[i], c.data[j] = c.data[j], c.data[i]
}

func (c CascadeSort) Less(i, j int) bool {
	for l := len(sStack) - 1; l >= 0; l-- {
		r := sFMap[sStack[l]](c.data[i], c.data[j])
		if r < 0 {
			return true
		} else if r > 0 {
			return false
		}
	}
	return false
}

// usStack updates the sorting stack sStack
func usStack(s string) bool {
	// check if s is a valid function according to sFMap
	exists := false
	for key, _ := range sFMap {
		if key == s {
			exists = true
			break
		}
	}
	if !exists {
		return false
	}

	// if sStack is empty add it directly
	if len(sStack) == 0 {
		sStack = append(sStack, s)
		return true
	}

	// check if element is already at the top if the stack
	if sStack[len(sStack)-1] == s {
		return false
	}

	// check if s is already present in sStack and remove it
	for i, val := range sStack[:len(sStack)-1] {
		if val == s {
			copy(sStack[i:], sStack[i+1:])
			sStack = sStack[:len(sStack)-1]

		}
	}

	// add value to the top of sStack
	sStack = append(sStack, s)
	return true
}

//!-sorting

//!+genertateHTML

func generateHTML(tracks []*Track, dst io.Writer, src io.Reader) error {

	// read source
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(src)
	if err != nil {
		return err
	}

	// parse template
	t, err := template.New("tracksweb").Parse(buf.String())

	if err != nil {
		return err
	}

	// write generated content
	err = t.Execute(dst, tracks)
	if err != nil {
		return err
	}

	return nil
}

//!-generateHTML

func main() {

	sdata := CascadeSort{data: tracks}

	handler := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm() // ignoring parse errors

		// sort according to provided parameter
		if usStack(r.FormValue("sort")) {
			sort.Sort(sdata)
		}

		// load template file
		tf, err := os.Open("index.tpl.html")
		if err != nil {
			log.Fatalf("error loading html template: %s", err)
			os.Exit(1)
		}
		defer tf.Close()

		// write response
		generateHTML(tracks, w, tf)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func cmpString(x, y string) int {
	if x > y {
		return 1
	} else if x == y {
		return 0
	} else {
		return -1
	}
}

func cmpByTitle(x, y *Track) int {
	return cmpString(x.Title, y.Title)
}

func cmpByArtist(x, y *Track) int {
	return cmpString(x.Artist, y.Artist)
}

func cmpByAlbum(x, y *Track) int {
	return cmpString(x.Album, y.Album)
}

func cmpByLength(x, y *Track) int {
	if x.Length > y.Length {
		return 1
	} else if x.Length == y.Length {
		return 0
	} else {
		return -1
	}
}

func cmpByYear(x, y *Track) int {
	if x.Year > y.Year {
		return 1
	} else if x.Year == y.Year {
		return 0
	} else {
		return -1
	}
}
