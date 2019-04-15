package main

import (
	"bytes"
	"fmt"

	"golang.org/x/net/html"
)

func main() {

	htmls := "<!DOCTYPE html><html><head><title>My Web Site</title></head><body>My Website's Body</body></html>"
	_, err := html.Parse(newReader(htmls))
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Println("HTML parsed successfully")
	}

}

func newReader(s string) *bytes.Reader {

	return bytes.NewReader([]byte(s))

}
