package roman_numerals

import (
	"fmt"
	"testing"
)

var test_cases = []struct {
	Arabic int
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{19, "XIX"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{49, "XLIX"},
	{50, "L"},
	{90, "XC"},
	{99, "XCIX"},
	{100, "C"},
	{399, "CCCXCIX"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
}

func TestConvertToRoman(t *testing.T) {
	for _, test := range test_cases {
		name := fmt.Sprintf("%d is %q", test.Arabic, test.Roman)
		t.Run(name, func(t *testing.T) {
			// Given
			want := test.Roman

			// When
			got := ConvertToRoman(test.Arabic)

			// Then
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range test_cases {
		name := fmt.Sprintf("%q is %d", test.Roman, test.Arabic)
		t.Run(name, func(t *testing.T) {
			// Given
			want := test.Arabic

			// When
			got := ConvertToArabic(test.Roman)

			// Then
			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}
