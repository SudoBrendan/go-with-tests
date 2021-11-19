package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func getGoodWebsites() []string {
	return []string{
		"www.google.com",
		"www.facebook.com",
	}
}

func getBadWebsites() []string {
	return []string{
		"foo.biz.baz.com",
	}
}

func getAllWebsites() []string {
	result := getGoodWebsites()
	result = append(result, getBadWebsites()...)
	return result
}

func getAllWebsitesHash() map[string]bool {
	result := make(map[string]bool)
	for _, u := range getGoodWebsites() {
		result[u] = true
	}
	for _, u := range getBadWebsites() {
		result[u] = false
	}
	return result
}

func mockWebSiteChecker(url string) bool {
	for _, u := range getBadWebsites() {
		if url == u {
			return false
		}
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	t.Run("get expected results", func(t *testing.T) {
		// Given
		websites := getAllWebsites()
		want := getAllWebsitesHash()

		// When
		got := CheckWebsites(mockWebSiteChecker, websites)

		// Then
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("wanted %v got %v", want, got)
		}
	})
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
