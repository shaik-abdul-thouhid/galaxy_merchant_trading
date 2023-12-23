package parse

import (
	"coffee/assessment/utils"
	"fmt"
	"strings"
)

const invalidInput = "invalid input"

// compareCredits compares two credit values represented as strings.
// It splits the input strings into roman numeral and metal type parts,
// converts them to numeric values, calculates the total credits for each input,
// and returns a message indicating which input has more credits.
// Returns an error if the inputs are invalid.
func (p *parseQueries) compareCredits(input1, input2 string) (string, error) {

	input1Array := strings.Split(strings.Trim(input1, " "), " ")
	input2Array := strings.Split(strings.Trim(input2, " "), " ")

	var metal1, metal2 float64

	if value, ok := p.metalToCreditMap[input1Array[len(input1Array)-1]]; !ok {
		return "", fmt.Errorf("requested number is in invalid format")
	} else {
		metal1 = value
	}

	if value, ok := p.metalToCreditMap[input2Array[len(input2Array)-1]]; !ok {
		return "", fmt.Errorf("requested number is in invalid format")
	} else {
		metal2 = value
	}

	input1Array = input1Array[:len(input1Array)-1]
	input2Array = input2Array[:len(input2Array)-1]

	roman1, roman2 := "", ""

	for _, item := range input1Array {
		if value, ok := p.aliasToRomanLiteral[item]; ok {
			roman1 += value
			continue
		}
		return "", returnInvalidInput()
	}

	for _, item := range input2Array {
		if value, ok := p.aliasToRomanLiteral[item]; ok {
			roman2 += value
			continue
		}
		return "", returnInvalidInput()
	}

	input1Value, err1 := utils.RomanToInt64(roman1)
	input2Value, err2 := utils.RomanToInt64(roman2)

	if err1 != nil || err2 != nil {
		return "", returnInvalidInput()
	}

	credits1, credits2 := float64(input1Value)*metal1, float64(input2Value)*metal2

	if credits1 > credits2 {
		return fmt.Sprintf("%s has more Credits than %s", input1, input2), nil
	} else {
		return fmt.Sprintf("%s has less Credits than %s", input1, input2), nil
	}
}

// compareRomanValues converts the given Roman numeral strings into integer values,
// compares them, and returns a message indicating which is larger.
// It splits the input strings into tokens, converts each token to
// its Roman numeral equivalent, concatenates the equivalents into strings,
// converts those to integers, and compares the integer values.
// Returns an error if invalid Roman numeral tokens are encountered.
func (p *parseQueries) compareRomanValues(input1, input2 string) (string, error) {
	input1Array := strings.Split(strings.Trim(input1, " "), " ")
	input2Array := strings.Split(strings.Trim(input2, " "), " ")

	tempInput1, tempInput2 := "", ""

	for _, item := range input1Array {
		var (
			value string
			ok    bool
		)

		if value, ok = p.aliasToRomanLiteral[item]; !ok {
			return "", fmt.Errorf(invalidInput)
		}
		tempInput1 += value
	}

	for _, item := range input2Array {
		var (
			value string
			ok    bool
		)

		if value, ok = p.aliasToRomanLiteral[item]; !ok {
			return "", fmt.Errorf(invalidInput)
		}
		tempInput2 += value
	}

	input1Value, err1 := utils.RomanToInt64(tempInput1)
	input2Value, err2 := utils.RomanToInt64(tempInput2)

	if err1 != nil || err2 != nil {
		return "", fmt.Errorf(invalidInput)
	}

	if input1Value > input2Value {
		return fmt.Sprintf("%s is larger than %s", input1, input2), nil
	}

	return fmt.Sprintf("%s is not larger than %s", input1, input2), nil
}
