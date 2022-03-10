package main

import (
	"flag"
	"github.com/KyleWardle/gophercises/html-link-parser/parser"
	"github.com/davecgh/go-spew/spew"
	"os"
)

func main() {
	example := flag.String("example", "ex1", "")
	flag.Parse()

	file, err := os.Open("examples/" + *example + ".html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	links := parser.ParseLinks(file)
	spew.Dump(links)
}
