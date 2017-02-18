package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	URL         string = "https://andruxnet-random-famous-quotes.p.mashape.com/?cat="
	TYPE_MOVIES string = "movies"
	TYPE_FAMOUS string = "famous"
)

type Quote struct {
	Quote    string `json:"quote"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

func (q Quote) String() string {
	return fmt.Sprintf("Quote:%s, Author:%s, Category:%s", q.Quote, q.Author, q.Category)
}

func main() {
	// Get the url
	url := getURL(TYPE_MOVIES)
	fmt.Println("url:", url)

	// Create a new HTTP client
	client := &http.Client{}

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Mashape-Key", getKey())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Callers should close resp.Body when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	var randomQuote Quote

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&randomQuote); err != nil {
		log.Println(err)
	}

	fmt.Println(randomQuote)
}

func getURL(category string) string {
	return URL + category
}

func getKey() string {
	key := os.Getenv("RANDOM_QUOTES_KEY")
	if key == "" {
		panic("'RANDOM_QUOTES_KEY' not set.")
	}

	return key
}
