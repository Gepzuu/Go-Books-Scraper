# Description:
A Go-based web scraper designed to extract book information from the BooksToScrape website. Utilizing the powerful Colly library, this application efficiently navigates the website, collects data, and formats it into structured JSON output.

# Features:
Concurrent Scraping: Leverages Colly's asynchronous capabilities to scrape multiple pages concurrently, improving efficiency and speed.
Structured Data Extraction: Collects detailed information about each book, including the link, name, price, and stock status.
Error Handling: Includes robust error handling mechanisms to log and manage potential issues during the scraping process.
Performance Monitoring: Incorporates a timer function to measure and display the total execution time for the scraping process.
Thread Safety: Utilizes sync.Mutex to ensure thread-safe operations when accessing shared resources.
Prerequisites
Go 1.15 or higher
Colly library
Installation
Clone the Repository:

sh
Copy code
git clone https://github.com/your-username/books-to-scrape-web-scraper.git
cd books-to-scrape-web-scraper
Install Dependencies:
Ensure you have Go installed. Then, run:

sh
Copy code
go mod tidy
Install Colly:

sh
Copy code
go get -u github.com/gocolly/colly
Usage
Run the Scraper:
Execute the main Go file:

sh
Copy code
go run main.go
View the Results:
The output will be printed in JSON format, displaying the collected book details.
