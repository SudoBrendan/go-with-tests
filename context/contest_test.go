package context

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	// mock something slow - appending chars to str; put in data
	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	// race: cancel or finish writing data?
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		// Given
		want := "hello world"
		store := &SpyStore{response: want, t: t}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		// When
		svr.ServeHTTP(response, request)
		got := response.Body.String()

		// Then
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("cancel writes no response", func(t *testing.T) {
		// Given
		want := "hello world"
		store := &SpyStore{response: want, t: t}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := &SpyResponseWriter{}

		// When
		svr.ServeHTTP(response, request)

		// Then
		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
