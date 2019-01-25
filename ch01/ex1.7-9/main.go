// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	urlPrefix := []string{"http://", "https://"}
	for _, url := range os.Args[1:] {
		hasPrefix := false
		for _, pfx := range urlPrefix {
			if strings.HasPrefix(url, pfx) {
				hasPrefix = true
			}
		}
		if !hasPrefix {
			url = urlPrefix[0] + url
		}
		println("Calling url:", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		respReader := bufio.NewReader(resp.Body)
		fmt.Printf("Status code: %v\n", resp.StatusCode)
		io.Copy(os.Stdout, respReader)
		resp.Body.Close()
	}
}

//!-
