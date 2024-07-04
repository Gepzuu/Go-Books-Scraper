package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"sync"
	"time"
)

type Item struct {
	Link    string `json:"link"`
	Name    string `json:"name"`
	Price   string `json:"price"`
	Instock string `json:"instock"`
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")()

	c := colly.NewCollector(colly.Async(true))

	var (
		items []Item
		mu    sync.Mutex
	)

	visitWithLogging := func(url string) {
		err := c.Visit(url)
		if err != nil {
			log.Printf("Error visiting %s: %v", url, err)
		}
	}

	c.OnHTML("div.side_categories li ul li", func(h *colly.HTMLElement) {
		link := h.ChildAttr("a", "href")
		visitWithLogging(h.Request.AbsoluteURL(link))
	})

	c.OnHTML("li.next a", func(h *colly.HTMLElement) {
		visitWithLogging(h.Request.AbsoluteURL(h.Attr("href")))
	})

	c.OnHTML("article.product_pod", func(h *colly.HTMLElement) {
		i := Item{
			Link:    h.ChildAttr("a", "href"),
			Name:    h.ChildAttr("h3 a", "title"),
			Price:   h.ChildText("p.price_color"),
			Instock: h.ChildText("p.instock"),
		}

		mu.Lock()
		items = append(items, i)
		mu.Unlock()
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	visitWithLogging("https://books.toscrape.com/catalogue/category/books/travel_2/index.html")
	c.Wait()

	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling items to JSON: %v", err)
	}

	fmt.Println(string(data))
}
