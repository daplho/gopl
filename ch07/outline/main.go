package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type CustomReader struct {
	s        string
	i        int64 // current reading index
	prevRune int   //index of previous rune; or < 0
}

func (r *CustomReader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func main() {
	doc, err := html.Parse(newReader("<html></html>"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func newReader(s string) *CustomReader {
	return &CustomReader{s, 0, -1}
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
