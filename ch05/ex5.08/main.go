// Outline prints the outline of an HTML document tree up to a provided element id
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	outline(os.Args[1], os.Args[2])
}

func outline(url, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement, id)
	//!-call

	return nil
}

var isFound bool

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
// Function will stop traversal when element with attr id == provided id is found

func forEachNode(n *html.Node, pre func(n *html.Node, id string) bool, post func(n *html.Node), id string) {
	if pre != nil {
		isFound = pre(n, id)
	}

	for c := n.FirstChild; c != nil && !isFound; c = c.NextSibling {
		forEachNode(c, pre, post, id)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node, id string) bool {
	var found bool
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
			if a.Key == "id" && a.Val == id {
				found = true
			}
		}
		if n.FirstChild != nil {
			fmt.Printf(">\n")
		} else {
			fmt.Printf("/>\n")
		}
		depth++
	} else if n.Type == html.TextNode && n.Data != "\n" {
		for _, line := range strings.Split(n.Data, "\n") {
			if line != "" {
				fmt.Printf("%*s%s\n", depth*2, "", line)
			}
		}
	}
	return found
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

//!-startend
