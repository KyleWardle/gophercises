package builder

import (
	"github.com/KyleWardle/gophercises/html-link-parser/parser"
	"github.com/davecgh/go-spew/spew"
	"net/http"
	generator "sitemap-builder/Builder/XmlGenerator"
	"strings"
)

type Builder struct {
	baseUrl  string
	sitemap  *generator.Xml
	maxDepth int
}

func New(href string, maxDepth int) Builder {
	return Builder{
		baseUrl:  href,
		maxDepth: maxDepth,
		sitemap:  generator.New(),
	}
}

func (b Builder) GenerateSitemap() {
	b.checkSite("/", 0)
	b.sitemap.SaveToFile(b.generateFileName())
}

func (b Builder) checkSite(href string, depth int) {
	spew.Dump("Running " + b.baseUrl + href)
	resp, err := http.Get(b.baseUrl + href)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	links := parser.ParseLinks(resp.Body)

	for _, link := range links {
		localLink := b.getLocalDomain(link.Href)
		if localLink == "" {
			continue
		}

		if b.sitemap.HasUrl(b.baseUrl + localLink) {
			continue
		}

		b.sitemap.AddUrl(b.baseUrl + localLink)

		depth++
		if depth < b.maxDepth {
			b.checkSite(localLink, depth)
		}
	}
}

func (b Builder) getLocalDomain(newLink string) string {
	if newLink[len(newLink)-1:] == "/" {
		newLink = newLink[0:(len(newLink) - 1)]
	}

	if len(newLink) == 0 {
		return ""
	}

	if newLink[0:1] == "/" {
		return newLink
	}

	newLink = b.stripAllLinkStuff(newLink)
	originalLink := b.stripAllLinkStuff(b.baseUrl)

	if len(newLink) < len(originalLink) {
		return ""
	}

	if newLink[0:len(originalLink)] == originalLink {
		return newLink[len(originalLink):]
	}

	return ""
}

func (b Builder) stripAllLinkStuff(link string) string {
	link = strings.ReplaceAll(link, "www.", "")
	link = strings.ReplaceAll(link, "https://", "")
	link = strings.ReplaceAll(link, "http://", "")
	return link
}

func (b Builder) generateFileName() string {
	url := b.stripAllLinkStuff(b.baseUrl)
	url = strings.ReplaceAll(url, ".", "")
	url = strings.ReplaceAll(url, "/", "")
	return "./SiteMaps/" + url + ".xml"
}
