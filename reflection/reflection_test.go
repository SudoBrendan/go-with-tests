package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	City string
	Age  int
}

func assertContains(t testing.TB, array []string, str string) {
	t.Helper()
	for _, x := range array {
		if x == str {
			return
		}
	}
	t.Errorf("expected %+v to contain %q", array, str)
}

func TestWalk(t *testing.T) {
	// Given
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with multiple string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non-string fields",
			struct {
				Name string
				City string
				Age  int
			}{"Chris", "London", 33},
			[]string{"Chris", "London"},
		},
		{
			"struct with nested string fields",
			Person{
				"Chris",
				Profile{"London", 33},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointers",
			&Person{
				"Chris",
				Profile{"London", 33},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{"London", 33},
				{"New York", 28},
			},
			[]string{"London", "New York"},
		},
		{
			"arrays",
			[2]Profile{
				{"London", 33},
				{"New York", 28},
			},
			[]string{"London", "New York"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			// When
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			// Then
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("maps", func(t *testing.T) {
		// Given
		m := map[string]string{
			"foo": "bar",
			"biz": "baz",
		}

		// When
		var got []string
		walk(m, func(input string) {
			got = append(got, input)
		})

		// Then
		assertContains(t, got, "bar")
		assertContains(t, got, "baz")
	})

	t.Run("channels", func(t *testing.T) {
		// Given
		c := make(chan Profile)
		go func() {
			c <- Profile{"London", 33}
			c <- Profile{"New York", 28}
			close(c)
		}()
		want := []string{"London", "New York"}

		// When
		var got []string
		walk(c, func(input string) {
			got = append(got, input)
		})

		// Then
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("functions", func(t *testing.T) {
		// Given
		f := func() (Profile, Profile) {
			return Profile{"London", 33}, Profile{"New York", 28}
		}
		want := []string{"London", "New York"}

		// When
		var got []string
		walk(f, func(input string) {
			got = append(got, input)
		})

		// Then
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
