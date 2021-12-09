package poker_test

import (
	"strings"
	"testing"

	"github.com/sudobrendan/gowithtests/app/internal/poker"
)

func TestCLI(t *testing.T) {
	t.Run("`Chris wins` stores win", func(t *testing.T) {
		// Given
		in := strings.NewReader("Chris wins\n")
		store := &poker.StubPlayerStore{}
		cli := poker.NewCLI(store, in)
		want := []string{"Chris"}

		// When
		cli.PlayPoker()

		// Then
		poker.AssertPlayerStoreWinCalls(t, *store, want)

	})

	t.Run("`Cleo wins` stores win", func(t *testing.T) {
		// Given
		in := strings.NewReader("Cleo wins\n")
		store := &poker.StubPlayerStore{}
		cli := poker.NewCLI(store, in)
		want := []string{"Cleo"}

		// When
		cli.PlayPoker()

		// Then
		poker.AssertPlayerStoreWinCalls(t, *store, want)

	})
}
