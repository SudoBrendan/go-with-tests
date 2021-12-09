package poker_test

import (
	"testing"

	"github.com/sudobrendan/gowithtests/app/internal/poker"
)

func TestFilesystemStore(t *testing.T) {
	t.Run("returns league sorted by highest score", func(t *testing.T) {
		// Given
		db, rmDb := poker.CreateTempFileStore(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer rmDb()
		store, err := poker.NewFileSystemPlayerStore(db)
		want := []poker.Player{
			{Name: "Chris", Wins: 33},
			{Name: "Cleo", Wins: 10},
		}

		// When
		got := store.GetLeague()

		// Then
		poker.AssertNoError(t, err)
		poker.AssertLeague(t, got, want)
	})

	t.Run("should be capable of reading multiple times", func(t *testing.T) {
		// Given
		db, rmDb := poker.CreateTempFileStore(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer rmDb()
		store, err := poker.NewFileSystemPlayerStore(db)
		want := []poker.Player{
			{Name: "Chris", Wins: 33},
			{Name: "Cleo", Wins: 10},
		}

		// When
		store.GetLeague()
		got := store.GetLeague()

		// Then
		poker.AssertNoError(t, err)
		poker.AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		// Given
		db, rmDb := poker.CreateTempFileStore(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer rmDb()
		store, err := poker.NewFileSystemPlayerStore(db)
		want := 33

		// When
		got := store.GetPlayerScore("Chris")

		// Then
		poker.AssertNoError(t, err)
		poker.AssertScoresEqual(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		// Given
		db, rmDb := poker.CreateTempFileStore(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer rmDb()
		store, err := poker.NewFileSystemPlayerStore(db)
		want := 34

		// When
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")

		// Then
		poker.AssertNoError(t, err)
		poker.AssertScoresEqual(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		// Given
		db, rmDb := poker.CreateTempFileStore(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer rmDb()
		store, err := poker.NewFileSystemPlayerStore(db)
		want := 1

		// When
		store.RecordWin("Charlie")
		got := store.GetPlayerScore("Charlie")

		// Then
		poker.AssertNoError(t, err)
		poker.AssertScoresEqual(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		// Given
		db, rmDb := poker.CreateTempFileStore(t, ``)
		defer rmDb()

		// When
		_, err := poker.NewFileSystemPlayerStore(db)

		// Then
		poker.AssertNoError(t, err)
	})
}
