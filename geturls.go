package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	var urlList []string

	for node := range doc.Descendants() {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			// isMalformed := false
			// for _, attr := range node.Attr {
			// 	if attr.Val == "" {
			// 		fmt.Printf("malformed link: %+v\n", node.Attr)
			// 		isMalformed = true
			// 		break
			// 	}
			// }

			// if !isMalformed {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					if strings.TrimSpace(attr.Val) == "" {
						fmt.Println("wrong url")
						continue
					}

					newURL, err := url.Parse(attr.Val)
					if err != nil {
						fmt.Printf("failed to parse URL: %v\n", err)
						continue
					}

					//fmt.Println("This goes in:", newURL.String(), "-->", attr)
					urlList = append(urlList, baseURL.ResolveReference(newURL).String())
				}
			}
			//}
		}
	}

	//fmt.Println(urlList)

	return urlList, nil

}
