package sync

import (
	"sync"
	"testing"
)

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}

func TestCounter(t *testing.T) {

	t.Run("incrementing the counter to 3 leaves it at 3", func(t *testing.T) {
		// Given
		counter := NewCounter()
		want := 3

		// When
		counter.Inc()
		counter.Inc()
		counter.Inc()

		// Then
		assertCounter(t, counter, want)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		// Given
		counter := NewCounter()
		want := 1000

		// When
		var wg sync.WaitGroup
		wg.Add(want)
		for i := 0; i < want; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		// Then
		assertCounter(t, counter, want)
	})
}
