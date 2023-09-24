package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println("Document Outline")
	outline(nil, doc)

	fmt.Println("\nLinks in the Document")
	links := visit(nil, doc)
	for _, link := range links {
		fmt.Println("*", link)
	}
}

// visit finds the links in HTML document
func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}

	if node.FirstChild != nil {
		// A recursive call to first child
		links = visit(links, node.FirstChild)
	}

	if node.NextSibling != nil {
		// A recursive call to next sibling
		links = visit(links, node.NextSibling)
	}

	return links
}

// outline prints the HTML document outline
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// A recursive call to each child
		outline(stack, c)
	}
}
