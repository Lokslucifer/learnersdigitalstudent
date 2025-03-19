package service

import (
	"net/url"
	"strings"
	"golang.org/x/net/html"
)

type PageContent struct {
	Paragraphs []string
	Links      []string
	Baseurl    string
}

func resolveURL(base, href string) string {
	parsedBase, err := url.Parse(base)
	if err != nil {
		return ""
	}

	parsedHref, err := url.Parse(href)
	if err != nil {
		return ""
	}

	resolvedURL := parsedBase.ResolveReference(parsedHref).String()
	return resolvedURL
}

func extractText(n *html.Node, content *PageContent, baseurl string) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "p":
			paraText := getText(n)
			if paraText != "" {
				content.Paragraphs = append(content.Paragraphs, paraText)
			}
		case "a":
			linkText := getText(n)
			var linkHref string
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					linkHref = attr.Val
				}
			}
			if linkText != "" && linkHref != "" {
				content.Paragraphs = append(content.Paragraphs, linkText)
				absoluteURL := resolveURL(baseurl, linkHref)
				content.Links = append(content.Links, absoluteURL)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractText(c, content, baseurl)
	}
}

func getText(n *html.Node) string {
	var textContent strings.Builder

	if n.Type == html.TextNode {
		textContent.WriteString(strings.TrimSpace(n.Data))
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		textContent.WriteString(" " + getText(c))
	}
	result := strings.TrimSpace(textContent.String())

	return result
}
