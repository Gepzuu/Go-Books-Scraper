package main

import (
	"github.com/gocolly/colly"
)

func main() {
	c :=colly.NewCollector()

	c.OnHTML("li.next a", func(h *colly.HTMLElement) {
		c.Visit(h.Request.AbsoluteURL(h.Attr("href")))

	})
}