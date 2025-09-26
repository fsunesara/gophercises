package parser

import (
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Link struct {
	Href string
	Text string
}

func processAnchor(node *html.Node) Link {
	link := Link{}

	for _, a := range node.Attr {
		if a.Key == "href" {
			link.Href = a.Val
			break
		}
	}

	for d := range node.Descendants() {
		if d.Type == html.TextNode {
			text := strings.TrimSpace(d.Data)
			if link.Text != "" && text != "" {
				link.Text += " "
			}
			link.Text += text
		}
	}
	return link
}

func ParseFromFile(fileName string) ([]Link, error) {
	rawHTML, err := os.ReadFile(fileName)
	if err != nil {
		return []Link{}, err
	}
	return ParseHTML(rawHTML)
}

func ParseHTML(rawHTML []byte) ([]Link, error) {
	doc, err := html.Parse(strings.NewReader(string(rawHTML)))
	if err != nil {
		return []Link{}, err
	}

	links := make([]Link, 0)
	for node := range doc.Descendants() {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			links = append(links, processAnchor(node))
		}
	}
	return links, nil
}
