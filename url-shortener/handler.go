package main

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		for shortUrl, intendedUrl := range pathsToUrls {
			if shortUrl == request.URL.Path {
				http.Redirect(writer, request, intendedUrl, http.StatusSeeOther)
			}
		}

		fallback.ServeHTTP(writer, request)
	})
}

type ShorterUrl struct {
	Path string `yaml:"path,omitempty"`
	Url  string `yaml:"url,omitempty"`
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	urls, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	mapped := convertToStringMap(urls)
	return MapHandler(mapped, fallback), nil
}

func parseYAML(yml []byte) ([]ShorterUrl, error) {
	var urls []ShorterUrl

	err := yaml.Unmarshal(yml, &urls)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func convertToStringMap(urls []ShorterUrl) map[string]string {
	stringMap := map[string]string{}

	for _, url := range urls {
		stringMap[url.Path] = url.Url
	}

	return stringMap
}
