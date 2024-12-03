package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func part1() {
	dataList := parseInputData()
	var safeReportsCount uint16 = 0
	const Asc = "asc"
	const Desc = "Desc"
	for _, row := range dataList {
		firstDiff := (row[0] - row[1])

		if firstDiff != 0 {
			rowInitialDirection := Asc
			if firstDiff < 0 {
				rowInitialDirection = Desc
			}

			sortDirModifier := 1
			if rowInitialDirection == Asc {
				sortDirModifier = -1
			}
			isRowSorted := slices.IsSortedFunc(row, func(a int, b int) int {
				if a > b {
					return 1 * sortDirModifier
				}
				if b > a {
					return -1 * sortDirModifier
				}
				return 0
			})

			if isRowSorted {
				for numIdx, num := range row {
					if (numIdx + 1) == len(row) {
						safeReportsCount++
					} else {
						nextNum := row[numIdx+1]
						diffWithNext := num - nextNum

						if diffWithNext == 0 || IntAbs(diffWithNext) > 3 {
							break
						}
					}

				}
			}
		}
	}

	fmt.Printf("The number of safe reports is: %v ", safeReportsCount)
}

func IntAbs(someint int) int {
	if someint < 0 {
		return -someint
	}

	return someint
}

func parseInputData() [][]int {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	parsedData := make([]([]int), 0)

	scanner := bufio.NewScanner(file)

	const ValueSeparator = " "

	for scanner.Scan() {
		numStrings := strings.Split(scanner.Text(), ValueSeparator)

		parsedNumSlice := make([]int, len(numStrings))

		for numStrIdx, numString := range numStrings {
			parsedNum, err := strconv.Atoi(numString)
			if err != nil {
				log.Fatal(err)
			}
			parsedNumSlice[numStrIdx] = parsedNum
		}

		parsedData = append(parsedData, parsedNumSlice)

	}

	return parsedData

}
