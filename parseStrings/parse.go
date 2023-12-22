package parse

import (
	"fmt"
	"strings"
)

var metalToCreditMap = make(map[string]float64)

var aliasToRomanLiteral = make(map[string]string)

func returnInvalidInput() error {
	return fmt.Errorf("I have no idea what you are talking about")
}

func GetMaps() (map[string]float64, map[string]string) {
	return metalToCreditMap, aliasToRomanLiteral
}

func ParseLines(input string) error {
	stringSlice := strings.Split(input, " is ")

	if len(stringSlice) == 2 {

		if strings.Contains(stringSlice[1], "Credits") {
			if err := populateMetalToCreditMap(stringSlice); err != nil {
				return err
			}

			return nil
		} else if strings.Contains(stringSlice[0], "Credits") {
		} else {
			return populateAliasToRomanLiteral(stringSlice)
		}

	} else if len(stringSlice) == 1 {

		return nil
	}

	return nil
}
