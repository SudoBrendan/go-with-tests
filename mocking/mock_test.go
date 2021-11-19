package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const write = "write"
const sleep = "sleep"

type SpySleeper struct{ Calls int }

func (s *SpySleeper) Sleep() { s.Calls++ }

type SpyCountdownOperations struct{ Calls []string }

func (s *SpyCountdownOperations) Sleep() { s.Calls = append(s.Calls, sleep) }
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) { s.durationSlept = duration }

func TestCountdown(t *testing.T) {

	t.Run("correct output", func(t *testing.T) {
		// Given
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		want := `3
2
1
Go!
`

		// When
		Countdown(buffer, spySleeper)
		got := buffer.String()

		// Then
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleeps", func(t *testing.T) {
		// Given
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		want := 3

		// When
		Countdown(buffer, spySleeper)
		got := spySleeper.Calls

		// Then
		if got != want {
			t.Errorf("want %d sleep calls, got %d", want, got)
		}
	})

	t.Run("correct function call order", func(t *testing.T) {
		// Given
		spySleepPrinter := &SpyCountdownOperations{}
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		// When
		Countdown(spySleepPrinter, spySleepPrinter)
		got := spySleepPrinter.Calls

		// Then
		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted calls %v got %v", want, got)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	// Given
	want := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{want, spyTime.Sleep}

	// When
	sleeper.Sleep()
	got := spyTime.durationSlept

	// Then
	if want != got {
		t.Errorf("wanted to sleep %v, got %v", want, got)
	}
}
