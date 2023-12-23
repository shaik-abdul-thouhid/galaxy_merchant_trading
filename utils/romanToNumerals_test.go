package utils

import (
	"coffee/assessment/setup"
	"testing"
)

func TestRomanToInt64(t *testing.T) {
	tests := []setup.TestInput[int64]{
		{Name: "valid input", Input: "MMVI", ExpectedResult: 2006, ExpectedError: ""},
		{Name: "empty input", Input: "", ExpectedResult: 0, ExpectedError: "requested number is in invalid format"},
		{Name: "invalid input", Input: "IIIII", ExpectedResult: 0, ExpectedError: "requested number is in invalid format"},
	}

	for _, test := range tests {
		result, err := RomanToInt64(test.Input)

		if test.ExpectedError != "" {
			if err == nil {
				t.Errorf("Expected to return %v for input: %s, but got %s", test.ExpectedError, test.Input, err)
			}
		}

		if result != test.ExpectedResult {
			t.Errorf("Expected output %v, got %v", test.ExpectedResult, result)
		}
	}
}
