package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://books.toscrape.com/"

	// Make HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching URL: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("Unexpected status code: %d", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalf("Error loading HTML: %v", err)
	}

	doc.Find(".product_pod").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3 a").Text()
		price := s.Find(".price_color").Text()

		// Clean up price (remove non-breaking space and currency symbol)
		price = strings.TrimSpace(strings.Replace(price, "Â", "", -1))

		fmt.Printf("Book %d: Title - %s, Price - %s\n", i+1, title, price)
	})
}

apiKey := "API_KEY"     
	cseID := "CSE_ID" =
	query := "web scraping with Go lang" 

	url := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s", apiKey, cseID, url.QueryEscape(query))

	//HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making request to Google Custom Search API: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Google Custom Search API request failed with status code %d", response.StatusCode)
	}

	//response
	var data map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}

	// Extract
	items, ok := data["items"].([]interface{})
	if !ok {
		log.Fatal("No search results found")
	}

	// Print
	fmt.Println("Search Results:")
	for _, item := range items {
		if result, ok := item.(map[string]interface{}); ok {
			title := result["title"]
			link := result["link"]
			fmt.Printf("%s\n%s\n\n", title, link)
		}
	}
}
