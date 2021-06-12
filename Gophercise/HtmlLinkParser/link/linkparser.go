package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type linker struct {
	Href string
	Text string
}

func DFS(doc *html.Node, padding string) {
	msg := doc.Data
	if doc.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)

	// what an implementation of DFS(Inorder)
	for n := doc.FirstChild; n != nil; {
		DFS(n, padding+" ")
		n = n.NextSibling
	}
}

func Parser(r io.Reader) ([]*linker, error) {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println("unable to parse html doc")
		return nil, err
	}
	DFS(doc, "")
	return nil, nil
}
