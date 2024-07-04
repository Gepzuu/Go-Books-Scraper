package main

import (
	"github.com/gocolly/colly"
)

func main() {
	c :=colly.NewCollector()

	c.OnHTML("", func(h *colly.HTMLElement) {

		
	})
}