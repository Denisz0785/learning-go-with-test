package concurrency

func mockWebSIteChecker(url string) bool {
	if url == "waat://fur.geds" {
		return false
	}
	return true
}
