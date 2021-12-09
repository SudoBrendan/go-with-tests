package poker_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/sudobrendan/gowithtests/app/internal/poker"
)

func TestCLI(t *testing.T) {
	t.Run("`Chris wins` stores win", func(t *testing.T) {
		// Given
		in := strings.NewReader("Chris wins\n")
		store := &poker.StubPlayerStore{}
		blindAlerter := &poker.SpyBlindAlerter{}
		cli := poker.NewCLI(store, in, blindAlerter)
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
		blindAlerter := &poker.SpyBlindAlerter{}
		cli := poker.NewCLI(store, in, blindAlerter)
		want := []string{"Cleo"}

		// When
		cli.PlayPoker()

		// Then
		poker.AssertPlayerStoreWinCalls(t, *store, want)

	})

	t.Run("it scheduled printing of blinds", func(t *testing.T) {
		// Given
		in := strings.NewReader("Chris wins\n")
		store := &poker.StubPlayerStore{}
		blindAlerter := &poker.SpyBlindAlerter{}
		cli := poker.NewCLI(store, in, blindAlerter)
		cases := []poker.StubBlindAlert{
			poker.NewStubBlindAlert(0*time.Second, 100),
			poker.NewStubBlindAlert(10*time.Minute, 200),
			poker.NewStubBlindAlert(20*time.Minute, 300),
			poker.NewStubBlindAlert(30*time.Minute, 400),
			poker.NewStubBlindAlert(40*time.Minute, 500),
			poker.NewStubBlindAlert(50*time.Minute, 600),
			poker.NewStubBlindAlert(60*time.Minute, 800),
			poker.NewStubBlindAlert(70*time.Minute, 1000),
			poker.NewStubBlindAlert(80*time.Minute, 2000),
			poker.NewStubBlindAlert(90*time.Minute, 4000),
			poker.NewStubBlindAlert(100*time.Minute, 8000),
		}

		// When
		cli.PlayPoker()
		alerts := blindAlerter.GetAlerts()

		for i, c := range cases {
			name := fmt.Sprintf("%d scheduled for %v", c.GetAmount(), c.GetScheduledAt())
			t.Run(name, func(t *testing.T) {
				// Then
				if len(alerts) <= i {
					t.Fatalf("alert %+v was not scheduled: %v", c, alerts)
				}
				poker.AssertBlindsEqual(t, alerts[i], c)
			})
		}
	})
}
