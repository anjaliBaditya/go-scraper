package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/PuerkitoBio/goquery"
)

func main() {
    // Specify the URL you want to scrape
    url := "https://example.com"

    // Make HTTP GET request
    response, err := http.Get(url)
    if err != nil {
        log.Fatal("Error fetching URL: ", err)
    }
    defer response.Body.Close()

    // Load HTML document
    doc, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTML: ", err)
    }

    // Find and extract data
    doc.Find("h1").Each(func(i int, s *goquery.Selection) {
        // Extract text from the selected element
        title := s.Text()
        fmt.Printf("Title %d: %s\n", i+1, title)
    })

    // You can further process or store the extracted data as needed
}
