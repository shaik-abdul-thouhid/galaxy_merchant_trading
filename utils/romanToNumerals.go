package utils

import (
	"fmt"
	"strings"
)

var RomanLiteralsToInt64Mapping = map[string]int64{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// RomanToInt64 converts a roman numeral string to an int64.
// It validates the format, counting occurrences of numerals and
// ensuring they don't exceed 3. It trims whitespace, gets the
// last numeral, and converts to int64. It then iterates from
// end to start, adding/subtracting values based on order per
// roman numeral rules, returning the result int64 and nil error
// or 0 and an error for invalid input.
func RomanToInt64(s string) (int64, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("requested number is in invalid format")
	}

	var romanNumeralCounts = make(map[string]int)

	for _, char := range s {
		romanNumeralCounts[string(char)]++
	}

	for _, count := range romanNumeralCounts {
		if count > 3 {
			return 0, fmt.Errorf("requested number is in invalid format")
		}
	}

	s = strings.Trim(s, " ")

	lastElement := string(s[len(s)-1])

	result, ok := RomanLiteralsToInt64Mapping[lastElement]

	if !ok {
		return 0, fmt.Errorf("invalid roman literal: %s", lastElement)
	}

	for i := len(s) - 1; i > 0; i-- {
		if RomanLiteralsToInt64Mapping[s[i:i+1]] <= RomanLiteralsToInt64Mapping[s[i-1:i]] {
			result += RomanLiteralsToInt64Mapping[s[i-1:i]]
		} else {
			result -= RomanLiteralsToInt64Mapping[s[i-1:i]]
		}
	}

	return result, nil
}

// Int64ToRoman converts an int64 to a roman numeral string.
// It iterates over arrays of roman numeral values and symbols,
// building the roman numeral string for the input number.
func Int64ToRoman(num int64) string {

	values := []int64{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	for i := int64(0); i < int64(len(values)); i++ {
		for num >= values[i] {
			roman += numerals[i]
			num -= values[i]
		}
	}
	return roman
}
