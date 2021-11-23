package roman_numerals

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

type ArabicCache struct {
	Cache map[string]int
}

func NewArabicCache() ArabicCache {
	return ArabicCache{
		Cache: make(map[string]int),
	}
}

func (c *ArabicCache) PopulateArabicMap() {
	if c.Cache["I"] != 1 {
		for i := 0; i < 3999; i++ {
			roman := ConvertToRoman(i)
			c.Cache[roman] = i
		}
	}
}

var cache = NewArabicCache()

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

func ConvertToArabic(roman string) int {
	// For real, this sounds like it'll be a trash algorithm that performs like trash too!
	// So... brute force it and look it up. We've only got 3999 values.
	cache.PopulateArabicMap()
	return cache.Cache[roman]
}
