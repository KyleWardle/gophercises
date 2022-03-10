package main

import (
	"flag"
	"github.com/KyleWardle/gophercises/html-link-parser/parser"
	"github.com/davecgh/go-spew/spew"
	"net/http"
	"strings"
)

func main() {
	url := flag.String("url", "https://calhoun.io", "")
	flag.Parse()

	sitemap := map[string]bool{}
	sitemap = checkSite(*url, "/", sitemap)
	spew.Dump(sitemap)
}

func checkSite(baseUrl string, href string, sitemap map[string]bool) map[string]bool {
	spew.Dump("Running " + baseUrl + href)
	resp, err := http.Get(baseUrl + href)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	links := parser.ParseLinks(resp.Body)

	for _, link := range links {
		//spew.Dump("Checking " + link.Href)

		localLink := getLocalDomain(link.Href, baseUrl)
		if localLink == "" {
			continue
		}

		if sitemap[baseUrl+localLink] {
			continue
		}

		sitemap[baseUrl+localLink] = true
		sitemap = checkSite(baseUrl, localLink, sitemap)
	}

	return sitemap
}

func getLocalDomain(newLink string, originalLink string) string {
	if newLink[len(newLink)-1:] == "/" {
		newLink = newLink[0:(len(newLink) - 1)]
	}

	if len(newLink) == 0 {
		return ""
	}

	if newLink[0:1] == "/" {
		return newLink
	}

	newLink = stripAllLinkStuff(newLink)
	originalLink = stripAllLinkStuff(originalLink)

	if len(newLink) < len(originalLink) {
		return ""
	}

	if newLink[0:len(originalLink)] == originalLink {
		return newLink[len(originalLink):]
	}

	return ""
}

func stripAllLinkStuff(link string) string {
	link = strings.ReplaceAll(link, "www.", "")
	link = strings.ReplaceAll(link, "https://", "")
	link = strings.ReplaceAll(link, "http://", "")
	return link
}
