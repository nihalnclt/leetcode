package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetchURL(url string, results chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}

	defer resp.Body.Close()
	elapsed := time.Since(start).Seconds()
	results <- fmt.Sprintf("Fetched %s in %.2f seconds with status code %d", url, elapsed, resp.StatusCode)
}

func main() {
	urls := []string{
		"https://aerodeals-admin.zapie.dev",
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
	}

	results := make(chan string, len(urls))
	defer close(results)

	for _, url := range urls {
		go fetchURL(url, results)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-results)
	}
}
