package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/sudobrendan/learn-go-with-tests/blogposts"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("STUB FS FAILURE")
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

}

var happyPathFS = fstest.MapFS{
	"hello world.md": {Data: []byte(`Title: Post 1
Description: Description 1
Tags: [ "tdd", "go" ]
---
Hello
World`)},
	"hello-world2.md": {Data: []byte(`Title: Post 2
Description: Description 2
Tags: [ "rust", "borrow-checker" ]
---
B
L
M`)},
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("returns n posts for n files", func(t *testing.T) {
		// Given
		testFS := happyPathFS
		want := []blogposts.Post{
			{
				Title:       "Post 1",
				Description: "Description 1",
				Tags:        []string{"tdd", "go"},
				Body: `Hello
World`,
			},
			{
				Title:       "Post 2",
				Description: "Description 2",
				Tags:        []string{"rust", "borrow-checker"},
				Body: `B
L
M`,
			},
		}

		// When
		got, err := blogposts.NewPostsFromFS(testFS)

		// Then
		if err != nil {
			t.Fatal(err)
		}
		for i, post := range got {
			assertPost(t, post, want[i])
		}
	})

	t.Run("propogates FS errors", func(t *testing.T) {
		// Given
		failingFS := StubFailingFS{}

		// When
		_, err := blogposts.NewPostsFromFS(failingFS)

		// Then
		if err == nil {
			t.Errorf("wanted an error but got none")
		}
	})
}

func BenchmarkNewBlogPosts(b *testing.B) {
	testFS := happyPathFS
	for i := 0; i < b.N; i++ {
		_, err := blogposts.NewPostsFromFS(testFS)
		if err != nil {
			b.Fatal(err)
		}
	}
}
