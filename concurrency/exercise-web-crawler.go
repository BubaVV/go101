package main

import (
	"fmt"
	"sync"
)

var cache map[string]bool
var mutex sync.Mutex

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan string) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer close(ch)
	if depth <= 0 {
		return
	}

	// fmt.Println("Checking url:", url)
	mutex.Lock()
	if _, ok := cache[url]; ok {
		// fmt.Println("Found in cache, exiting")
		defer mutex.Unlock()
		return
	} else {
		cache[url] = true
	}
	mutex.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch <- body
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		// fmt.Println("Cache:", cache)
		// fmt.Println("Running", u, depth-1)
		child_ch := make(chan string)
		go Crawl(u, depth-1, fetcher, child_ch)
		for result := range child_ch {
			ch <- result
		}
	}
	return
}

func main() {
	cache = map[string]bool{}
	ch := make(chan string)
	go Crawl("https://golang.org/", 4, fetcher, ch)

	for result := range ch {
		fmt.Println(result)
	}

}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// In this exercise you'll use Go's concurrency features to parallelize a web crawler.

// Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.

// Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!
