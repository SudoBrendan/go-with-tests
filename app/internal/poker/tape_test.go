package poker

import (
	"io/ioutil"
	"testing"
)

func TestTape_Write(t *testing.T) {
	t.Run("should be capable of writing shorter content later", func(t *testing.T) {
		// Given
		file, clean := CreateTempFileStore(t, "12345")
		defer clean()
		tape := &tape{file}
		want := "abc"

		// When
		tape.Write([]byte(want))
		file.Seek(0, 0)
		newFileContents, _ := ioutil.ReadAll(file)
		got := string(newFileContents)

		// Then
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
