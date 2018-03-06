package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findelements: %v\n", err)
		os.Exit(1)
	}

	nodes := ElementsByTagName(doc, "h1", "h2", "h3")
	for _, node := range nodes {
		fmt.Printf("%s\n", node.Data)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	names := make(map[string]struct{})
	for _, n := range name {
		names[n] = struct{}{}
	}

	var nodes []*html.Node
	visitNode := func(n *html.Node, names map[string]struct{}) {
		if n.Type == html.ElementNode {
			if _, ok := names[n.Data]; ok {
				nodes = append(nodes, n)
			}
		}
	}
	forEachNode(doc, names, visitNode)
	return nodes
}

// forEachNode calls the function action for each node x in the tree rooted at
// n. Action is optional and is called before the children of the node are
// visited.
func forEachNode(n *html.Node, names map[string]struct{}, action func(n *html.Node, names map[string]struct{})) {
	if action != nil {
		action(n, names)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, names, action)
	}
}
