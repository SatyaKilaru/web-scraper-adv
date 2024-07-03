package scraper

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func extractTitle(doc *html.Node, titleXPath string) string {
	titleNode, err := htmlquery.Query(doc, titleXPath)
	if err != nil || titleNode == nil {
		return ""
	}
	return htmlquery.InnerText(titleNode)
}

func extractBody(doc *html.Node, bodyXPath string) string {
	bodyNode, err := htmlquery.Query(doc, bodyXPath)
	if err != nil || bodyNode == nil {
		return ""
	}
	return htmlquery.OutputHTML(bodyNode, true)
}
