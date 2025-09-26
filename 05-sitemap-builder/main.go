package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"

	parser "example.com/04-html-link-parser"
)

type UrlSet struct {
	XMLName xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
	Urls    []Url    `xml:"url"`
}

type Url struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

var visited map[string]struct{}

func processUrl(url string, baseUrl string, depth int, maxDepth int) []Url {
	visited[url] = struct{}{}
	urls := make([]Url, 0)
	if depth > maxDepth {
		return urls
	}
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}

	links, err := parser.ParseHTML(body)

	for _, link := range links {
		link.Href = strings.TrimSuffix(link.Href, "/")
		if strings.HasPrefix(link.Href, url) || strings.HasPrefix(link.Href, "/") {
			if strings.HasPrefix(link.Href, "/") {
				link.Href = baseUrl + link.Href
			}
			if _, ok := visited[link.Href]; !ok {
				newUrl := Url{}
				newUrl.Loc = link.Href
				urls = append(urls, newUrl)
				urls = append(urls, processUrl(link.Href, baseUrl, depth+1, maxDepth)...)
			}
		}
	}
	return urls
}

func main() {
	visited = make(map[string]struct{})
	url := flag.String("url", "https://www.calhoun.io", "the url to create a sitemap for")
	flag.Parse()

	urls := processUrl(*url, *url, 0, 2)
	urlSet := UrlSet{}
	urlSet.Urls = urls
	str, err := xml.MarshalIndent(urlSet, " ", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(xml.Header + string(str))
}
