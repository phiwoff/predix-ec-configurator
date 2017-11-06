package helpers

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
)

// ParseMarkdown reads a markdownfile and
// returns as an HTML template
func ParseMarkdown(filename string) template.HTML {
	// load markdown file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// convert markdown to html
	unsafe := blackfriday.MarkdownCommon(content)
	// replace code-parts with syntax-highlighted parts
	replaced, err := replaceCodeParts(unsafe)
	if err != nil {
		log.Fatal(err)
	}
	// sanitize html
	//html := bluemonday.UGCPolicy().SanitizeBytes(replaced)

	return template.HTML(string(replaced))
}

func replaceCodeParts(mdFile []byte) (string, error) {
	byteReader := bytes.NewReader(mdFile)
	doc, err := goquery.NewDocumentFromReader(byteReader)
	if err != nil {
		return "", err
	}
	// find code-parts via css selector and replace them with highlighted versions
	doc.Find("code[class*=\"language-\"]").Each(func(i int, s *goquery.Selection) {
		oldCode := s.Text()
		formatted, err := syntaxhighlight.AsHTML([]byte(oldCode))
		if err != nil {
			log.Fatal(err)
		}
		s.SetHtml(string(formatted))
	})
	new, err := doc.Html()
	if err != nil {
		return "", err
	}
	// replace unnecessarily added html tags
	new = strings.Replace(new, "<html><head></head><body>", "", 1)
	new = strings.Replace(new, "</body></html>", "", 1)
	return new, nil
}
