package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	inputData := parseInputData()

	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := regex.FindAllStringSubmatch(inputData, -1)

	multSum := 0

	if len(matches) > 0 {
		for _, match := range matches {
			mult1, err := strconv.Atoi(match[1])
			mult2, err := strconv.Atoi(match[2])

			if err != nil {
				log.Fatal(err)
			} else {
				multSum += mult1 * mult2
			}
		}
	}

	fmt.Printf("The result of all the matches mults is %v\n", multSum)
}

func part2() {
	inputData := parseInputData()
	searchData := inputData
	doCmds := ""

	findDoBlock := false // Commands start enabled

	for {
		if findDoBlock {
			splitResults := strings.SplitN(searchData, "do()", 2)
			findDoBlock = false
			if len(splitResults) < 2 {
				break
			} else {
				searchData = splitResults[1]
			}
		} else {
			splitResults := strings.SplitN(searchData, "don't()", 2)
			findDoBlock = true
			doCmds = doCmds + splitResults[0]
			if len(splitResults) < 2 {
				break
			} else {
				searchData = splitResults[1]
			}
		}
	}

	multSum := 0
		secondRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		secondMatches := secondRegex.FindAllStringSubmatch(doCmds, -1)

		if len(secondMatches) > 0 {
			for _, subMatch := range secondMatches{
				mult1, err := strconv.Atoi(subMatch[1])
				mult2, err := strconv.Atoi(subMatch[2])

				if err != nil {
					log.Fatal(err)
				} else {
					multSum += mult1 * mult2
				}
			}
		}

	fmt.Printf("The result of all the do-matches mults is %v\n", multSum)
}

func parseInputData() string {
	file, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	parsedData := string(file)

	return parsedData

}
