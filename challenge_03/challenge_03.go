package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
  part1()
}

func part1() {
  inputData:= parseInputData()

  regex:= regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

  matches := regex.FindAllStringSubmatch(inputData, -1)

  multSum := 0

  if (len(matches) > 0) {
    for _, match := range matches {
      mult1, err := strconv.Atoi(match[1])
      mult2, err := strconv.Atoi(match[2])

      if err != nil {
        log.Fatal(err)
      } else {
        multSum += mult1*mult2
      }
    }   
  }

  fmt.Printf("The result of all the matches mults is %v\n", multSum)

}

func parseInputData() string {
	file, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	parsedData := string(file)

	return parsedData

}
