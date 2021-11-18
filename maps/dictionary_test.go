package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("happy", func(t *testing.T) {
		// given
		want := "this is just a test"

		// when
		got, err := dict.Search("test")

		// then
		assertStrings(t, want, got)
		assertNoError(t, err)
	})

	t.Run("unknown word", func(t *testing.T) {
		// given
		want := ErrWordUnknown

		// when
		_, got := dict.Search("basdpfoihbas")

		// then
		assertError(t, want, got)
	})
}

func TestAdd(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Add("test", "this is just a test")
		assertNoError(t, err)
		assertDefinition(t, dict, "test", "this is just a test")
	})

	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{"test": "first value"}
		err := dict.Add("test", "second value")
		assertError(t, ErrWordAlreadyDefined, err)
		assertDefinition(t, dict, "test", "first value")
	})
}

func assertDefinition(t testing.TB, dict Dictionary, word, def string) {
	t.Helper()
	got, err := dict.Search(word)
	assertStrings(t, got, def)
	assertNoError(t, err)
}

func assertStrings(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}
}

func assertError(t testing.TB, want error, got error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("wanted no error but got one")
	}
}
