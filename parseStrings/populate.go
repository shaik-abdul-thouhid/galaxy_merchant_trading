package parse

import (
	"coffee/assessment/utils"
	"strconv"
	"strings"
)

func removeNewLineCharacters(input string) string {
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "\r", "")
	input = strings.ReplaceAll(input, "\n", "")

	return input
}

// populateAliasToRomanLiteral populates the aliasToRomanLiteral map
// by parsing the input string slice. It takes the first string
// as the alias, trims whitespace, and splits on spaces. The second
// string is the Roman numeral, which is also trimmed of whitespace
// and newline characters. The Roman numeral is validated and then
// mapped to the alias. Returns error if input is invalid.
func (p *parseQueries) populateAliasToRomanLiteral(stringSlice []string) error {
	firstString := strings.Split(strings.Trim(stringSlice[0], " "), " ")

	secondString := strings.Trim(stringSlice[1], " ")
	secondString = removeNewLineCharacters(secondString)

	if len(firstString) > 1 {
		return returnInvalidInput()
	}

	if _, err := utils.RomanToInt64(secondString); err == nil {
		p.aliasToRomanLiteral[firstString[0]] = secondString
	} else {
		return returnInvalidInput()
	}
	return nil
}

// populateMetalToCreditMap parses a string slice input to populate the
// metalToCreditMap field of the parseQueries struct. It takes the first
// string as the metal alias, splits it on spaces, and maps each part to
// its Roman numeral equivalent. It takes the second string, removes "Credits",
// trims whitespace, and parses it to an integer for the number of credits.
// It divides the credits by the Roman numeral total to get the exchange rate,
// which is stored in the map with the original metal alias as the key.
// Returns any errors encountered during parsing and conversion.
func (p *parseQueries) populateMetalToCreditMap(stringSlice []string) error {
	replacedString := strings.ReplaceAll(stringSlice[1], "Credits", "")
	replacedString = removeNewLineCharacters(replacedString)

	credit, err := strconv.ParseInt(replacedString, 10, 64)

	if err != nil {
		return returnInvalidInput()
	}
	firstStringSlice := strings.Split(stringSlice[0], " ")

	creditString := ""

	for _, str := range firstStringSlice {
		if value, ok := p.aliasToRomanLiteral[str]; ok {
			creditString += value
			continue
		}

		valueOfRoman, err := utils.RomanToInt64(creditString)

		if err != nil {
			return returnInvalidInput()
		}

		p.metalToCreditMap[str] = float64(credit) / float64(valueOfRoman)
	}

	return nil
}
