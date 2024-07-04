# Description:
<div align="center">
A Go-based web scraper designed to extract book information from the BooksToScrape website. Utilizing the powerful Colly library, this application efficiently navigates the website, collects data, and formats it into structured JSON output.</div>

## Features:
Concurrent Scraping: Leverages Colly's asynchronous capabilities to scrape multiple pages concurrently, improving efficiency and speed. <br>
Structured Data Extraction: Collects detailed information about each book, including the link, name, price, and stock status.<br>
Error Handling: Includes robust error handling mechanisms to log and manage potential issues during the scraping process.<br>
Performance Monitoring: Incorporates a timer function to measure and display the total execution time for the scraping process.<br>
Thread Safety: Utilizes sync.Mutex to ensure thread-safe operations when accessing shared resources.

## Prerequisites:

Go 1.15 or higher<br>
Colly library<br>


## Installation:
Clone the Repository:


Copy code
```sh
git clone https://github.com/Gepzuu/books-to-scrape-web-scraper.git
cd books-to-scrape-web-scraper
```


Install Dependencies:
Ensure you have Go installed. Then, run:
```sh
go mod tidy
```
Install Colly:
```sh
go get -u github.com/gocolly/colly
```

## Usage:

Execute the main Go file:
```sh
go run main.go
```
View the Results:
The output will be printed in JSON format, displaying the collected book details.
