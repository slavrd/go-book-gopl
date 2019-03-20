package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("Error parsing html: %s", err)
		os.Exit(1)
	}

	counts := make(map[string]int)
	countWordsAndImages(counts, doc)

	fmt.Printf("\n%v\n", counts)

}

func countWordsAndImages(m map[string]int, node *html.Node) {

	if node.Type == html.ElementNode && node.Data == "img" {
		m["images"]++
	} else if node.Type == html.TextNode {
		m["words"] += countWords(node.Data)
	}

	if node.FirstChild != nil {
		countWordsAndImages(m, node.FirstChild)
	}
	if node.NextSibling != nil {
		countWordsAndImages(m, node.NextSibling)
	}
}

func countWords(s string) int {
	buf := bytes.NewBufferString(s)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)

	var result int
	for ; scanner.Scan(); result++ {
	}
	return result

}
