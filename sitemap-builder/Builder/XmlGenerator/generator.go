package generator

import (
	"encoding/xml"
	"io"
	"os"
)

type Xml struct {
	XMLName xml.Name `xml:"urlset"`
	Urls    []Url    `xml:"url"`
}

type Url struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

func New() *Xml {
	return &Xml{}
}

func (x *Xml) AddUrl(url string) {
	x.Urls = append(x.Urls, Url{
		Loc: url,
	})
}

func (x *Xml) HasUrl(url string) bool {
	for _, existingUrl := range x.Urls {
		if existingUrl.Loc != url {
			continue
		}

		return true
	}
	return false
}

func (x *Xml) SaveToFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	xmlWriter := io.Writer(file)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")
	if err := enc.Encode(x); err != nil {
		panic(err)
	}
}
