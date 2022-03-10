package main

import (
	"flag"
	"github.com/KyleWardle/gophercises/html-link-parser/parser"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	example := flag.String("example", "ex1", "")
	flag.Parse()

	links := parser.ParseLinks("examples/" + *example + ".html")
	spew.Dump(links)
}
