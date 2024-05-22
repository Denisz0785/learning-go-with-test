package concurrency

import "fmt"

type WebSiteChecker func(string) bool

type res struct {
	string
	bool
}

func CheckWebSite(wc WebSiteChecker, urls []string) map[string]bool {
	result := map[string]bool{}
	resultChannel := make(chan res)
	for _, url := range urls {

		go func(u string) {
			resultChannel <- res{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		result[r.string] = r.bool
	}
	close(resultChannel)
	return result
}

func Hello() {
	fmt.Println("Hello")
	fmt.Println("Good Luck")
}
