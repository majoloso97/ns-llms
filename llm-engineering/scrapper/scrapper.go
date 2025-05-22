package scrapper

import (
	"github.com/gocolly/colly/v2"
)

type WebPageContent struct {
	Title   string
	Content string
}

func ScrapURL(url string, content *WebPageContent) {
	c := colly.NewCollector()
	c.OnHTML("title", func(e *colly.HTMLElement) {
		content.Title = e.Text
	})
	c.OnHTML("body", func(e *colly.HTMLElement) {
		content.Content = e.Text
	})
	c.Visit(url)
}
