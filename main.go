package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages             map[string]int
	baseURL           *url.URL
	mu                *sync.Mutex
	concurrecyControl chan struct{}
	wg                *sync.WaitGroup
	maxPages          int
}

func parseArg(arg string, min int, name string) (int, error) {
	argRead, err := strconv.Atoi(arg)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %s: %w", name, err)
	}
	if argRead < min {
		return 0, fmt.Errorf("%s must be greater than %d", name, min)
	}
	return argRead, nil
}

func main() {
	args := os.Args
	maxConcurrency := 5
	maxPages := 10
	var err error

	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	rawBaseURL, err := url.Parse(args[1])
	if err != nil {
		fmt.Printf("failed to parse base URL: %v\n", err)
		os.Exit(1)
	}

	if len(args) > 2 {
		maxConcurrency, err = parseArg(args[2], 1, "concurrency")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if len(args) > 3 {
		maxPages, err = parseArg(args[3], 1, "max pages")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	cfg := config{
		pages:             make(map[string]int),
		baseURL:           rawBaseURL,
		mu:                &sync.Mutex{},
		concurrecyControl: make(chan struct{}, maxConcurrency),
		wg:                &sync.WaitGroup{},
		maxPages:          maxPages,
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(cfg.baseURL.String())
	cfg.wg.Wait()

	cfg.printReport()
}
