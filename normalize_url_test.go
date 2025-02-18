package main

import (
	"strings"
	testing "testing"
)

func TestNormalizeURL(t *testing.T) {
	test := map[string]struct {
		inputURL      string
		expected      string
		error_content string
	}{
		"remove scheme": {
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		"remove final slash": {
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		"lowercase": {
			inputURL: "https://BLOG.boot.dev/PATH",
			expected: "blog.boot.dev/path",
		},
		"lowercase 2": {
			inputURL: "http://BLOG.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		"path only": {
			inputURL: "/path",
			expected: "/path",
		},
		"wrong url": {
			inputURL:      "//wrong url",
			expected:      "",
			error_content: "parse",
		},
		"empty url": {
			inputURL:      "",
			expected:      "",
			error_content: "empty",
		},
	}

	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			actual, err := NormalizeURL(tc.inputURL)
			if err != nil {
				if !strings.Contains(err.Error(), tc.error_content) || tc.error_content == "" {
					t.Errorf("Test %v - %s FAIL: unexpected error: %v", name, tc.inputURL, err)
					return
				}
			} else if tc.error_content != "" {
				t.Errorf("Test %v - %s FAIL: expected error: %s", name, tc.inputURL, tc.error_content)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected %s, got %s", name, tc.inputURL, tc.expected, actual)
				return
			}
		})
	}
}
