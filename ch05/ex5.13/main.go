// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(u string) []string {
	fmt.Println(u)
	list, err := Extract(u)
	if err != nil {
		log.Print(err)
	}

	ulink, _ := url.Parse(u) //Ignoring possible error
	if checkIsInputHost(*ulink) {
		err := saveURL(*ulink)
		if err != nil {
			log.Print(err)
		}

	}

	return list
}

func saveURL(u url.URL) error {

	var dir string
	var fileName string

	isFile, _ := regexp.MatchString(".+\\..+", path.Base(u.Path)) // Ignore error

	if isFile {
		dir = path.Join(u.Host, path.Dir(u.Path))
		fileName = strings.Trim(path.Base(u.Path), "/")
	} else {
		dir = path.Join(u.Host, u.Path)
		fileName = "content.html"
	}

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("error saving url %s : %s", u.String(), err)
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return fmt.Errorf("error saving url %s : %s", u.String(), err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("error saving url %s : status code %v", u.String(), resp.StatusCode)
	}

	file, err := os.Create(path.Join(dir, fileName))
	if err != nil {
		resp.Body.Close()
		return fmt.Errorf("error saving url %s : %s", u.String(), err)
	}
	writer := bufio.NewWriter(file)

	_, err = io.Copy(writer, resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("error saving url %s : %s", u.String(), err)
	}

	writer.Flush()
	file.Close()

	return nil
}

//!-crawl

var inputHosts []string

func checkIsInputHost(u url.URL) bool {

	for _, inHost := range inputHosts {
		if u.Hostname() == inHost {
			return true
		}
	}

	return false

}

//!+main
func main() {

	// Store the hosts of the command line args urls
	for _, u := range os.Args[1:] {
		link, err := url.Parse(u)
		if err != nil {
			log.Printf("could not parse url: %s", u)
		} else {
			inputHosts = append(inputHosts, link.Hostname())
		}
	}

	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

//!-main
