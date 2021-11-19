package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func getDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("faster server wins", func(t *testing.T) {
		// Given
		slowServer := getDelayedServer(1 * time.Millisecond)
		fastServer := getDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		want := fastServer.URL

		// When
		got, err := Racer(slowServer.URL, fastServer.URL)

		// Then
		if err != nil {
			t.Errorf("wasn't expecting an error %v", err)
		}
		if want != got {
			t.Errorf("wanted %q got %q", want, got)
		}
	})

	t.Run("error after timeout", func(t *testing.T) {
		// Given
		server := getDelayedServer(50 * time.Millisecond)
		timeout := 1 * time.Millisecond
		defer server.Close()

		// When
		_, err := ConfigurableRacer(server.URL, server.URL, timeout)

		// Then
		if err == nil {
			t.Error("wanted an error, but got nil")
		}
	})
}
