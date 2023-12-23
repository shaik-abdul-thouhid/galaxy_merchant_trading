package parse

import (
	"coffee/assessment/utils"
	"fmt"
	"strings"
)

// queryHowMuch parses the user's query to determine the value of the
// roman numeral, converts it to an integer, and returns a string
// indicating the query text and corresponding integer value. It handles
// cleaning the input by removing newlines, carriage returns, and
// trimming whitespace.
func (p *parseQueries) queryHowMuch(input string) (string, error) {
	query := strings.ReplaceAll(input, "?", "")
	query = strings.ReplaceAll(query, "\n", "")
	query = strings.ReplaceAll(query, "\r", "")
	query = strings.Trim(query, " ")

	array := strings.Split(query, " ")

	romanChar := ""

	for _, str := range array {

		if value, ok := p.aliasToRomanLiteral[str]; ok {
			romanChar += value
		}
	}
	value, err := utils.RomanToInt64(romanChar)

	if err != nil {
		return "", returnInvalidInput()
	}

	return fmt.Sprintf("%s is %v", query, value), nil
}

// queryHowMany parses the user query to extract roman numerals, converts them to
// an integer value, looks up the metal type, and returns a string with the
// query text and corresponding total credits value. It handles cleaning the
// input by removing newlines, carriage returns, and trimming whitespace.
func (p *parseQueries) queryHowMany(input string) (string, error) {
	romanLiterals := ""

	input = strings.ReplaceAll(input, "?", "")
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\r", "")
	input = strings.Trim(input, " ")

	array := strings.Split(input, " ")

	var (
		metal float64
		ok    bool
	)

	if metal, ok = p.metalToCreditMap[array[len(array)-1]]; !ok {
		return "", fmt.Errorf("requested number is in invalid format")
	}

	for _, str := range array {
		if value, ok := p.aliasToRomanLiteral[str]; ok {
			romanLiterals += value
		} else {
			break
		}
	}

	credits, err := utils.RomanToInt64(romanLiterals)

	if err != nil {
		return "", fmt.Errorf("requested number is in invalid format")
	}

	return fmt.Sprintf("%s is %v Credits", input, float64(credits)*metal), nil
}
