package quote

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const URL string = "https://andruxnet-random-famous-quotes.p.mashape.com/?cat="

type Quote struct {
	Quote    string `json:"quote"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

// Stringer function for `Quote` struct
func (q Quote) String() string {
	return fmt.Sprintf("Quote: %s, Author: %s, Category: %s", q.Quote, q.Author, q.Category)
}

type QuoteRequest struct {
	Key string
}

func (qr QuoteRequest) GetNewQuote(c Category) Quote {
	// Get the url
	url := GetUrl(c)

	// HTTP client
	client := GetHttpClient()

	// Get constructed HTTP Request
	request := GetRequest(url, qr)

	// HTTP Response
	response := GetResponse(client, request)
	defer response.Body.Close()

	randomQuote := Quote{}

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(response.Body).Decode(&randomQuote); err != nil {
		panic(err)
	}

	return randomQuote
}

// Construct the URL using the given category and return it
func GetUrl(c Category) string {
	return URL + c.Type
}

func GetHttpClient() *http.Client {
	return &http.Client{}
}

func GetRequest(url string, qr QuoteRequest) *http.Request {
	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Mashape-Key", qr.Key)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	if err != nil {
		panic(err)
	}

	return req
}

func GetResponse(client *http.Client, request *http.Request) *http.Response {
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	return response
}
