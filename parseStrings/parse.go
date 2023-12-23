package parse

import (
	"coffee/assessment/utils"
	"fmt"
	"strings"
)

type parseQueries struct {
	metalToCreditMap map[string]float64

	aliasToRomanLiteral map[string]string
}

// NewParseQueries initializes a new parseQueries instance with empty maps for
// metal-to-credit and alias-to-roman literal mappings. It returns a pointer
// to the new parseQueries struct. As an exported function that constructs
// and returns a key struct, this serves as a constructor for the parseQueries
// type.
func NewParseQueries() *parseQueries {
	return &parseQueries{
		metalToCreditMap:    make(map[string]float64),
		aliasToRomanLiteral: make(map[string]string),
	}
}

func returnInvalidInput() error {
	return fmt.Errorf("i have no idea what you are talking about")
}

// handleStringsWithCredits parses input strings containing "Credits" and delegates
// to appropriate handling functions. It handles populating metal-to-credit mappings,
// querying total credits for a metal, and populating alias-to-roman literal mappings.
func (p *parseQueries) handleStringsWithCredits(stringSlice []string) (string, error) {
	if strings.Contains(stringSlice[1], "Credits") {
		if err := p.populateMetalToCreditMap(stringSlice); err != nil {
			return "", err
		}

		return "", nil

	} else if strings.Contains(stringSlice[0], "Credits") && strings.Contains(stringSlice[0], "how many Credits") {
		return p.queryHowMany(stringSlice[1])
	}

	if strings.Contains(stringSlice[0], "how much") {
		return p.queryHowMuch(stringSlice[1])
	}

	return "", p.populateAliasToRomanLiteral(stringSlice)

}

// handleStringsStartingWithIs parses input strings starting with "Is" and handles
// comparing roman numeral values. It splits the input on "larger than" and
// "smaller than" to extract the values to compare, and delegates to
// compareRomanValues.
func (p *parseQueries) handleStringsStartingWithIs(str string) (string, error) {

	if strings.Contains(str, " larger than ") {
		inputs := strings.Split(str, " larger than ")

		return p.compareRomanValues(inputs[0], inputs[1])
	} else if strings.Contains(str, " smaller than ") {
		inputs := strings.Split(str, " smaller than ")

		return p.compareRomanValues(inputs[0], inputs[1])
	}

	return "", returnInvalidInput()
}

// handleStringsStartingWithDoes parses input strings starting with "Does" and handles
// comparing credit values. It splits the input on "has more Credits than" and
// "has less Credits than" to extract the values to compare, and delegates to
// compareCredits.
func (p *parseQueries) handleStringsStartingWithDoes(str string) (string, error) {
	if strings.Contains(str, "  has more Credits than ") {
		inputs := strings.Split(str, "  has more Credits than ")

		return p.compareCredits(inputs[0], inputs[1])

	} else if strings.Contains(str, " has less Credits than ") {
		inputs := strings.Split(str, " has less Credits than ")

		return p.compareCredits(inputs[0], inputs[1])
	}

	return "", nil
}

// ParseLines takes an input string, splits it on " is ", and delegates to the
// appropriate handler function based on the input. It handles inputs starting
// with "Is" and "Does" by trimming the prefix and calling the corresponding
// handler. It returns the parsed result string and any error. This is the main
// entry point for parsing user input queries.
func (p *parseQueries) ParseLines(input string) (string, error) {
	stringSlice := strings.Split(input, " is ")

	if len(stringSlice) == 2 {

		return p.handleStringsWithCredits(stringSlice)

	} else if len(stringSlice) == 1 {

		if strings.HasPrefix(stringSlice[0], "Is ") {
			trimIs := utils.PrefixTrimQuestion(stringSlice[0], "Is ")

			return p.handleStringsStartingWithIs(trimIs)

		} else if strings.HasPrefix(stringSlice[0], "Does ") {
			trimDoes := utils.PrefixTrimQuestion(stringSlice[0], "Does ")

			return p.handleStringsStartingWithDoes(trimDoes)
		} else {
			return "", returnInvalidInput()
		}

	}

	return "", returnInvalidInput()
}
