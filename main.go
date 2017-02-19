package main

import (
	"fmt"
	"os"
	"random-quotes/quote"
)

func main() {
	quoteRequest := quote.QuoteRequest{Key: getKey()}
	fmt.Println(quoteRequest.GetNewQuote(quote.NewRandomCategory()))
	fmt.Println(quoteRequest.GetNewQuote(quote.NewRandomCategory()))
	fmt.Println(quoteRequest.GetNewQuote(quote.NewRandomCategory()))
}

func getKey() string {
	key := os.Getenv("RANDOM_QUOTES_KEY")
	if key == "" {
		panic("'RANDOM_QUOTES_KEY' not set.")
	}

	return key
}
