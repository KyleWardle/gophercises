package parser

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func ParseLinks(reader io.Reader) []Link {
	doc, err := html.Parse(reader)
	if err != nil {
		panic(err)
	}

	var links []Link
	iterateNode(doc, &links, nil)

	return links
}

func iterateNode(node *html.Node, links *[]Link, text *string) {
	if text == nil && node.Type == html.ElementNode && node.Data == "a" {
		internalText := ""

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			iterateNode(child, links, &internalText)
		}

		for _, attr := range node.Attr {
			if attr.Key != "href" {
				continue
			}

			*links = append(*links, Link{
				Text: cleanupText(internalText),
				Href: attr.Val,
			})
		}

	} else {
		if text != nil && node.Type == html.TextNode {
			*text += node.Data
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			iterateNode(child, links, text)
		}
	}
}

func cleanupText(text string) string {
	text = strings.TrimSpace(text)
	return text
}
