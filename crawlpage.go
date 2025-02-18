package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	// Concurrency control
	cfg.concurrecyControl <- struct{}{}
	defer func() {
		<-cfg.concurrecyControl
		cfg.wg.Done()
	}()

	// Check if we have reached the max number of pages
	if cfg.checkDepth() {
		return
	}

	// Check if we are out of the scope of the website
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to parse current URL: %v\n", err)
		return
	}
	if currentURL.Host != cfg.baseURL.Host {
		return
	}

	// Normalize URL
	normURL, err := NormalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to normalize URL: %v\n", err)
		return
	}

	// Check if we have already crawled this URL
	if !cfg.isFirstVisit(normURL) {
		return
	}
	fmt.Println("have:", normURL)

	// Fetch the HTML of the current URL
	currentBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("failed to fetch HTML: %v\n", err)
		return
	}

	// Get the URLs from the HTML body with the main baseURL
	nextURLs, err := getURLsFromHTML(currentBody, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("failed to get URLs from HTML: %v\n", err)
		return
	}

	// Crawl recursively all the URLs found in the HTML body
	for _, next := range nextURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(next)
	}
}

func (cfg *config) checkDepth() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	return len(cfg.pages) >= cfg.maxPages
}

func (cfg *config) isFirstVisit(someURL string) bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	// Check if we have already crawled this URL
	if _, ok := cfg.pages[someURL]; ok {
		//fmt.Println("already crawled:", normURL)
		cfg.pages[someURL]++
		return false
	}
	// Else, add it to the map
	cfg.pages[someURL] = 1
	return true
}
