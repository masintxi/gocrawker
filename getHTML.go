package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	response, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch website: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode > 399 {
		return "", fmt.Errorf("bad status code: %d", response.StatusCode)
	}

	urlHeader := response.Header.Get("Content-Type")
	if !strings.Contains(urlHeader, "text/html") {
		return "", fmt.Errorf("not a html page: %s", urlHeader)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(data), nil
}
