package main

import (
	"fmt"
	"net/url"
	"strings"
)

func NormalizeURL(rawURL string) (string, error) {
	newURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %w", err)
	}

	if newURL.String() == "" {
		return "", fmt.Errorf("empty URL")
	}

	finalURL := newURL.Host + newURL.Path

	finalURL = strings.ToLower(finalURL)

	finalURL = strings.TrimSuffix(finalURL, "/")

	return finalURL, nil
}
