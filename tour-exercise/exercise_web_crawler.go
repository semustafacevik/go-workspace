package tourexercise

import (
	"fmt"
	"sync"
	"time"
)

// https://tour.golang.org/concurrency/10
// Exercise: Web Crawler

type fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type safeCache struct {
	m sync.Mutex
	v map[string]bool
}

func (c *safeCache) Any(url string) bool {
	c.m.Lock()
	defer c.m.Unlock()

	_, ok := c.v[url]
	if !ok {
		c.v[url] = true
		return false
	}

	return true
}

var cache = &safeCache{v: make(map[string]bool)}

func crawl(url string, depth int, fetcher fetcher) {
	if depth <= 0 || cache.Any(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(" -  found: %s %q\n", url, body)

	for _, url := range urls {
		go crawl(url, depth-1, fetcher)
	}
}

// Solution of "web crawler" exercise
//
// For detail:
// https://tour.golang.org/concurrency/10
func WebCrawler() {
	crawl("https://golang.org/", 4, _fetcher)
	time.Sleep(time.Second)
}

type fakeResult struct {
	body string
	urls []string
}

type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}

var _fetcher = fakeFetcher{
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
