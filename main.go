package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

const URL = "https://phillipsbeer.com/growlers/"

func main() {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}

	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)

	for {
		t := tokenizer.Next()

		switch {
		case t == html.ErrorToken:
			return

		case t == html.StartTagToken:
			token := tokenizer.Token()
			if isABeerName(token) {
				t = tokenizer.Next()

				if t == html.TextToken {
					fmt.Printf("- %s\n", tokenizer.Token())
				}
			}
		}
	}
}

func isABeerName(token html.Token) bool {
	if token.Data == "span" {
		for _, attr := range token.Attr {
			if attr.Key == "class" && attr.Val == "beer-name" {
				return true
			}
		}
	}

	return false
}
