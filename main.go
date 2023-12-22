package main

import (
	"bufio"
	parse "coffee/assessment/parseStrings"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("sample_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		err = parse.ParseLines(line)

		if err != nil {
			log.Fatal(err)
		}
	}

	map1, map2 := parse.GetMaps()

	fmt.Println("map1", map1)
	fmt.Println("map2", map2)
}
