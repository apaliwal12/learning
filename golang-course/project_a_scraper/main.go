package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

// ─────── PROJECT A: CONCURRENT WEB SCRAPER ───────
// This project synthesizes:
// - Goroutines, Channels, WaitGroups (Module 9)
// - Context for timeout/cancellation (Module 10)
// - HTTP Client (Module 11)
// - Robust Error Handling (Module 14)

type Result struct {
	URL   string
	Title string
	Err   error
}

func fetchTitle(ctx context.Context, targetURL string) Result {
	res := Result{URL: targetURL}

	// 1. Create a request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		res.Err = fmt.Errorf("creating request: %w", err)
		return res
	}

	// 2. Make the HTTP call (using default client for simplicity, but customized timeout via ctx)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		res.Err = fmt.Errorf("fetching URL: %w", err)
		return res
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		res.Err = fmt.Errorf("bad status: %s", resp.Status)
		return res
	}

	// 3. Parse HTML to find <title>
	res.Title, res.Err = extractTitle(resp.Body)
	return res
}

func extractTitle(body io.Reader) (string, error) {
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			return "", fmt.Errorf("title not found")
		}
		
		if tt == html.StartTagToken {
			t := z.Token()
			if t.Data == "title" {
				tt = z.Next() // Move to the text inside the title
				if tt == html.TextToken {
					return strings.TrimSpace(z.Token().Data), nil
				}
			}
		}
	}
}

func main() {
	fmt.Println("--- Project A: Concurrent Web Scraper ---")

	urls := []string{
		"https://go.dev",
		"https://pkg.go.dev",
		"https://google.com",
		"https://github.com",
		"http://invalid.url.that.does.not.exist", // Will fail
	}

	// Set a global timeout for the entire scraping job
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	results := make(chan Result, len(urls))
	var wg sync.WaitGroup

	// FAN-OUT: Start a goroutine for each URL
	for _, u := range urls {
		// Basic validation
		if _, err := url.ParseRequestURI(u); err != nil {
			log.Printf("Skipping invalid URL: %s\n", u)
			continue
		}

		wg.Add(1)
		go func(target string) {
			defer wg.Done()
			results <- fetchTitle(ctx, target)
		}(u)
	}

	// Closer goroutine: wait for all to finish, then close the channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// FAN-IN: Read results from the channel
	for res := range results {
		if res.Err != nil {
			fmt.Printf("❌ ERROR fetching %s: %v\n", res.URL, res.Err)
		} else {
			fmt.Printf("✅ %s -> Title: %q\n", res.URL, res.Title)
		}
	}

	fmt.Println("Scraping complete.")
}
