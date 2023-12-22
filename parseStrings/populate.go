package parse

import (
	"coffee/assessment/utils"
	"strconv"
	"strings"
)

func populateAliasToRomanLiteral(stringSlice []string) error {
	firstString := strings.Split(strings.Trim(stringSlice[0], " "), " ")

	secondString := strings.Trim(stringSlice[1], " ")
	secondString = strings.ReplaceAll(secondString, " ", "")
	secondString = strings.ReplaceAll(secondString, "\r", "")
	secondString = strings.ReplaceAll(secondString, "\n", "")

	if len(firstString) > 1 {
		return returnInvalidInput()
	}

	if _, err := utils.RomanToInt64(secondString); err == nil {
		aliasToRomanLiteral[firstString[0]] = secondString
	} else {
		return returnInvalidInput()
	}
	return nil
}

func populateMetalToCreditMap(stringSlice []string) error {
	replacedString := strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(strings.ReplaceAll(stringSlice[1], "Credits", ""), " ", ""),
			"\r", "",
		),
		"\n", "",
	)

	credit, err := strconv.ParseInt(replacedString, 10, 64)

	if err != nil {
		return returnInvalidInput()
	}
	firstStringSlice := strings.Split(stringSlice[0], " ")

	creditString := ""

	for _, str := range firstStringSlice {
		if value, ok := aliasToRomanLiteral[str]; ok {
			creditString += value
		} else {
			valueOfRoman, err := utils.RomanToInt64(creditString)

			if err != nil {
				return returnInvalidInput()
			}

			metalToCreditMap[str] = float64(credit) / float64(valueOfRoman)
		}
	}

	return nil
}
