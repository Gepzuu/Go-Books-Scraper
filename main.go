package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"sync"
	"time"
)

// Define a struct to hold the scraped data
type Item struct {
	Link    string `json:"link"`
	Name    string `json:"name"`
	Price   string `json:"price"`
	Instock string `json:"instock"`
}

// Create a timer function to measure execution time
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	// Start the timer for the main function
	defer timer("main")()

	// Create a new Colly collector with async mode enabled
	c := colly.NewCollector(colly.Async(true))

	// Initialize variables to hold the scraped data and a mutex for synchronization
	var (
		items []Item
		mu    sync.Mutex
	)

	// Define a helper function to visit a URL with logging
	visitWithLogging := func(url string) {
		err := c.Visit(url)
		if err!= nil {
			log.Printf("Error visiting %s: %v", url, err)
		}
	}

	// Define a callback function to handle HTML elements with class "side_categories"
	c.OnHTML("div.side_categories li ul li", func(h *colly.HTMLElement) {
		// Extract the link from the HTML element
		link := h.ChildAttr("a", "href")
		// Visit the extracted link with logging
		visitWithLogging(h.Request.AbsoluteURL(link))
	})

	// Define a callback function to handle HTML elements with class "next"
	c.OnHTML("li.next a", func(h *colly.HTMLElement) {
		// Extract the link from the HTML element
		visitWithLogging(h.Request.AbsoluteURL(h.Attr("href")))
	})

	// Define a callback function to handle HTML elements with class "product_pod"
	c.OnHTML("article.product_pod", func(h *colly.HTMLElement) {
		// Extract the data from the HTML element
		i := Item{
			Link:    h.ChildAttr("a", "href"),
			Name:    h.ChildAttr("h3 a", "title"),
			Price:   h.ChildText("p.price_color"),
			Instock: h.ChildText("p.instock"),
		}

		// Lock the mutex to ensure thread safety
		mu.Lock()
		// Append the extracted data to the items slice
		items = append(items, i)
		// Unlock the mutex
		mu.Unlock()
	})

	// Define a callback function to handle requests
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Start the scraping process by visiting the initial URL
	visitWithLogging("https://books.toscrape.com/catalogue/category/books/travel_2/index.html")

	// Wait for the scraping process to complete
	c.Wait()

	// Marshal the scraped data to JSON
	data, err := json.MarshalIndent(items, "", "  ")
	if err!= nil {
		log.Fatalf("Error marshalling items to JSON: %v", err)
	}

	// Print the JSON data
	fmt.Println(string(data))
}