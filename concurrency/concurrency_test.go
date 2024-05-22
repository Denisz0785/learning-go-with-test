package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebSIteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebSite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebSite(mockWebSIteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v got %v", want, got)
	}

}

func slowWebSiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "OM A HUM"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebSite(slowWebSiteChecker, urls)
	}
}
