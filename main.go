package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Item struct{
	Link string 'json:"Link"'
	Name string 'json:"name"'
	Price string 'json:"price"'
	Instock string 'json:"instock"'
}
func main() {
	c :=colly.NewCollector()

	c.OnHTML("li.next a", func(h *colly.HTMLElement) {
		c.Visit(h.Request.AbsoluteURL(h.Attr("href")))

	})

	c.OnHTML("", func(h *colly.HTMLElement){
		fmt.Println(h.ChildAttr("h3 a", "title"))

       
	})

	c.OnRequest(func(r *colly.Request){
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://books.toscrape.com/catalogue/page-1.html")
}