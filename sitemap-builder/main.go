package main

import (
	"flag"
	builder "sitemap-builder/Builder"
)

func main() {
	url := flag.String("url", "https://calhoun.io", "")
	maxDepth := flag.Int("max-depth", 3, "")
	flag.Parse()

	b := builder.New(*url, *maxDepth)
	b.GenerateSitemap()
	println("Generated sitemap!")
}
