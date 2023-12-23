package main

import (
	"bufio"
	parse "coffee/assessment/parseStrings"
	"fmt"
	"log"
	"os"
)

func isValidFile(path string) bool {

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func main() {

	arguments := os.Args[1:]

	if len(arguments) != 1 {
		log.Fatal("Expected path to the input text file")
		os.Exit(1)
	} else if !isValidFile(arguments[0]) {
		log.Fatal("invalid file path")
		os.Exit(1)
	}

	file, err := os.Open(arguments[0])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	parser := parse.NewParseQueries()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		result, err := parser.ParseLines(line)

		if result != "" {
			fmt.Println(result)
		} else if err != nil {
			fmt.Println(err)
		}
	}
}
