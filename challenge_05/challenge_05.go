package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var rulesData, pageData = parseInputData()

func main() {
	// for _, rule := range rulesData {
	//   fmt.Println(rule)
	// }
	//
	// fmt.Println()
	//
	// for _, pages := range pageData {
	//   fmt.Println(pages)
	// }
	part1()
}

func part1() {
	middlePageSum := 0
	ruleKeys := slices.Collect(maps.Keys(rulesData))
	sort.Strings(ruleKeys)

	for _, pageToCheck := range pageData {
		pageFail := false
		fmt.Printf("Checking page: %v\n", pageToCheck)

		for _, rule := range ruleKeys {
			if pageFail {
				break
			}

			testIdx := slices.Index(pageToCheck, rule)
			fmt.Printf("Rule value: %v || TestIdx on currentpage: %v\n", rule, testIdx)



			// If the first number isn't in the list, skip remaining checks and go next
			if testIdx != -1 {
				for _, ruleCheckVal := range rulesData[rule] {
					evalIdx := slices.Index(pageToCheck, ruleCheckVal)

					if evalIdx < testIdx && evalIdx != -1 {
						fmt.Printf("Rules failed on ruleCheckVal %v, %v || evaluatedIdx was %v\n", rule, ruleCheckVal, evalIdx)
						pageFail = true
						break
					}
				}
			}
		}

		if !pageFail {
			fmt.Printf("Valid page: %v\n", pageToCheck)
			valToAdd, err := strconv.Atoi(pageToCheck[(len(pageToCheck)-1)/2])
			fmt.Printf("Valtoadd: %v\n", valToAdd)
			if err != nil {
				log.Panic(err)
			} else {
				middlePageSum += valToAdd
			}
		} else {
			fmt.Printf("Invalid page: %v\n", pageToCheck)
		}

	}

	fmt.Printf("The sum of valid middle pages is: %v\n", middlePageSum)
}

func parseInputData() (map[string][]string, [][]string) {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	rulesData := make([]([]string), 0)
	pageData := make([]([]string), 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// If rule
		if strings.Contains(scanner.Text(), "|") {
			rulesSplitData := strings.Split(scanner.Text(), "|")
			rulesData = append(rulesData, rulesSplitData)
			// Elif Pages
		} else if strings.Contains(scanner.Text(), ",") {

			pageSplitData := strings.Split(scanner.Text(), ",")
			pageData = append(pageData, pageSplitData)
		}
	}

	rulesMap := make(map[string]([]string))

	for _, ruleTuple := range rulesData {
		if rulesMap[ruleTuple[0]] == nil {
			rulesMap[ruleTuple[0]] = make([]string, 0)

			rulesMap[ruleTuple[0]] = append(rulesMap[ruleTuple[0]], ruleTuple[1])
		} else {
			rulesMap[ruleTuple[0]] = append(rulesMap[ruleTuple[0]], ruleTuple[1])
		}
	}

	return rulesMap, pageData
}
