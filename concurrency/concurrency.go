package concurrency

import "fmt"

type WebSiteChecker func(string) bool

func CheckWebSite(wc WebSiteChecker, urls []string) map[string]bool {
	result := map[string]bool{}

	for _, url := range urls {
		result[url] = wc(url)
	}
	return result
}

func Hello() {
	fmt.Println("Hello,World!")
}
