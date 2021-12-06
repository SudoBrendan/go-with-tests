package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFilesystemStore(t *testing.T) {
	t.Run("returns league sorted by highest score", func(t *testing.T) {
		// Given
		db, rmDb := createTempFileStore(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer rmDb()
		store, err := NewFileSystemPlayerStore(db)
		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		// When
		got := store.GetLeague()

		// Then
		assertNoError(t, err)
		assertLeague(t, got, want)
	})

	t.Run("should be capable of reading multiple times", func(t *testing.T) {
		// Given
		db, rmDb := createTempFileStore(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer rmDb()
		store, err := NewFileSystemPlayerStore(db)
		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		// When
		store.GetLeague()
		got := store.GetLeague()

		// Then
		assertNoError(t, err)
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		// Given
		db, rmDb := createTempFileStore(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer rmDb()
		store, err := NewFileSystemPlayerStore(db)
		want := 33

		// When
		got := store.GetPlayerScore("Chris")

		// Then
		assertNoError(t, err)
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		// Given
		db, rmDb := createTempFileStore(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer rmDb()
		store, err := NewFileSystemPlayerStore(db)
		want := 34

		// When
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")

		// Then
		assertNoError(t, err)
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		// Given
		db, rmDb := createTempFileStore(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer rmDb()
		store, err := NewFileSystemPlayerStore(db)
		want := 1

		// When
		store.RecordWin("Charlie")
		got := store.GetPlayerScore("Charlie")

		// Then
		assertNoError(t, err)
		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		// Given
		db, rmDb := createTempFileStore(t, ``)
		defer rmDb()

		// When
		_, err := NewFileSystemPlayerStore(db)

		// Then
		assertNoError(t, err)
	})
}

func createTempFileStore(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))
	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
