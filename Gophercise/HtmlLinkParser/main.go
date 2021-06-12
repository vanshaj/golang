package main

import (
	"os"

	"github.com/vanshaj/golang/Gophercise/HtmlLinkParser/link"
)

func main() {
	r, _ := os.Open("Data/ex2.html")
	_, err := link.Parser(r)
	if err != nil {
		return
	}
}
