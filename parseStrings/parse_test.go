package parse

import (
	"coffee/assessment/setup"
	"testing"
)

func initializeTest() parseQueries {
	parser := parseQueries{
		metalToCreditMap:    make(map[string]float64),
		aliasToRomanLiteral: make(map[string]string),
	}

	return parser
}

const validRomanAlias = "Parse Valid Roman Alias"

const validMetalCredits = "Parse Valid Metal Credits with Roman Aliases"

const noIdea = "i have no idea what you are talking about"

func TestParseString(t *testing.T) {
	tests := []setup.TestInput[string]{
		{Name: validRomanAlias, Input: "glob is I", ExpectedResult: "", ExpectedError: ""},
		{Name: "Parse InValid Roman Alias", Input: "glob is I am", ExpectedResult: "", ExpectedError: noIdea},
		{Name: validRomanAlias, Input: "prok is V", ExpectedResult: "", ExpectedError: ""},
		{Name: validRomanAlias, Input: "pish is X", ExpectedResult: "", ExpectedError: ""},
		{Name: validRomanAlias, Input: "tegj is L", ExpectedResult: "", ExpectedError: ""},
		{Name: "Calculating Credits before assigning", Input: "how many Credits is glob prok Silver", ExpectedResult: "", ExpectedError: noIdea},
		{Name: validMetalCredits, Input: "glob glob Silver is 34 Credits", ExpectedResult: "", ExpectedError: ""},
		{Name: validMetalCredits, Input: "glob prok Gold is 57800 Credits", ExpectedResult: "", ExpectedError: ""},
		{Name: validMetalCredits, Input: "glob prok Iron is 3910 Credits", ExpectedResult: "", ExpectedError: ""},
		{Name: "Query valid metal credits", Input: "how many Credits is glob glob Gold ?", ExpectedResult: "glob glob Gold is 28900 Credits", ExpectedError: ""},
		{Name: "Query invalid metal credits", Input: "how many Credits is glob glob Platinum ?", ExpectedResult: "", ExpectedError: "number is in invalid format"},
		{Name: "Query valid roman alias", Input: "how much is pish tegj glob glob ?", ExpectedResult: "pish tegj glob glob is 42", ExpectedError: ""},
		{Name: "Query comparison of two metal's credits", Input: "Does glob glob Gold has less Credits than pish tegj glob glob Iron?", ExpectedResult: "glob glob Gold has less Credits than pish tegj glob glob Iron", ExpectedError: ""},
		{Name: "Invalid Input", Input: "how much wood could a woodchuck chuck if a woodchuck could chuck wood ?", ExpectedResult: "", ExpectedError: noIdea},
	}

	parser := initializeTest()

	for _, test := range tests {
		result, err := parser.ParseLines(test.Input)

		t.Logf("Test Name: %s\n", test.Name)

		if err != nil {
			if test.ExpectedError == "" {
				t.Errorf("Not expected any error, but got %v", err)
			}
		} else if result != test.ExpectedResult {
			t.Errorf("Expected %v, got %v", test.ExpectedResult, result)
		}
	}
}
