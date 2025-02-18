package main

import (
	"strings"
	testing "testing"
)

func TestGetURLs(t *testing.T) {
	test := map[string]struct {
		inputURL      string
		inputBody     string
		expected      []string
		error_content string
	}{
		"absolute and relative URLs": {
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<html>
					<body>
						<a href="/path/one">
							<span>Boot.dev</span>
						</a>
						<a href="https://other.com/path/one">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
				`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		"real blog.boot.dev": {
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<li>
					<a href="/computer-science/18-months-with-gpt-4/" class="border border-gray-600 cursor-pointer flex flex-col sm:flex-row p-4 my-4 hover:border-yellow-500">
						
						<div class="flex-2 mr-0 sm:mr-4 mb-2 sm:mb-0 flex justify-center">
						<img width="288" height="192" class="rounded object-cover hidden sm:block" src="/img/800/golemnumber2.png.webp">
						<img class="rounded object-cover block sm:hidden" src="/img/800/golemnumber2.png.webp">
						</div>
						

						<div class="flex-1">
						<h2 class="text-2xl text-white mb-2">18 Months with GPT-4: Now Can I Fire my Developers?</h2>

						
						<span class="text-sm text-gray-400 mb-4">
							Jan 17, 2025 by Lane Wagner
						</span>
						

						<p class="hidden sm:block">As the founder of a company where my largest static expense is engineering salaries, I’m over here just chomping at the bit, eagerly awaiting the moment I can fire everyone and line my pockets with all those juicy savings.</p>
						</div>
					</a>
				</li>
				`,
			expected: []string{"https://blog.boot.dev/computer-science/18-months-with-gpt-4/"},
		},
		"pkg.go.dev html#example-Parse": {
			inputURL: "https://pkg.go.dev/golang.org/x/net/html",
			inputBody: `
				<li>
					<a href="#example-Parse" class="js-exampleHref">Parse</a>
				</li>
				`,
			expected: []string{"https://pkg.go.dev/golang.org/x/net/html#example-Parse"},
		},
		"rare references": {
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<div class="navigation">
					<a href="../tutorials/getting-started">Back to Tutorials</a>
					<a href="./chapter2">Next Chapter</a>
					<a href="images/diagram.png">View Diagram</a>
					<a href="?sort=newest">Sort by Newest</a>
					<a href="#section-3">Jump to Section 3</a>
				</div>
				`,
			expected: []string{
				"https://blog.boot.dev/tutorials/getting-started",
				"https://blog.boot.dev/chapter2",
				"https://blog.boot.dev/images/diagram.png",
				"https://blog.boot.dev?sort=newest",
				"https://blog.boot.dev#section-3",
			},
		},
		"empty hrefs": {
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<a>No href at all</a>
				<a href="">Empty href</a>
				<a href="  ">Whitespace href</a>
			`,
			expected: nil,
		},
		// "invalid link": {
		// 	inputURL: "https://blog.boot.dev",
		// 	inputBody: `
		// 		<a href="valid-link">Good Link</a>
		// 		<a href="broken-link" <malformed>Bad Link</a>
		// 	`,
		// 	expected: []string{"https://blog.boot.dev/valid-link"},
		// },
		// "deeply nested invalid link": {
		// 	inputURL: "https://blog.boot.dev",
		// 	inputBody: `
		// 		<div>
		// 			<a href="valid-link-1">First Good Link</a>
		// 			<div>
		// 				<span>
		// 					<a href="valid-link-2">Second Good Link</a>
		// 					<div>
		// 						<a href="broken-link" <malformed>Bad Link</a>
		// 						<span>
		// 							<a href="valid-link-3">Third Good Link</a>
		// 						</span>
		// 					</div>
		// 					<a href="valid-link-4">Fourth Good Link</a>
		// 				</span>
		// 			</div>
		// 		</div>
		// 	`,
		// 	expected: []string{
		// 		"https://blog.boot.dev/valid-link-1",
		// 		"https://blog.boot.dev/valid-link-2",
		// 		"https://blog.boot.dev/valid-link-3",
		// 		"https://blog.boot.dev/valid-link-4",
		// 	},
		// },
	}

	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				if !strings.Contains(err.Error(), tc.error_content) || tc.error_content == "" {
					t.Errorf("Test %v - %s FAIL: unexpected error: %v", name, tc.inputURL, err)
					return
				}
			} else if tc.error_content != "" {
				t.Errorf("Test %v - %s FAIL: expected error: %s", name, tc.inputURL, tc.error_content)
				return
			}

			if len(actual) != len(tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected %d urls, got %d", name, tc.inputURL, len(tc.expected), len(actual))
				return
			}

			for i, url := range actual {
				if url != tc.expected[i] {
					t.Errorf("Test %v - %s FAIL: expected %s, got %s", name, tc.inputURL, tc.expected[i], url)
					return
				}
			}
		})
	}

}
