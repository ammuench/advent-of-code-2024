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
	fmt.Println()
	part2()
}

const Asc = "Asc"
const Desc = "Desc"

func part1() {
	dataList := parseInputData()
	var safeReportsCount uint16 = 0
	for _, row := range dataList {
		firstDiff := (row[0] - row[1])

		if firstDiff != 0 {
			rowInitialDirection := Asc
			if firstDiff > 0 {
				rowInitialDirection = Desc
			}

			sortDirModifier := 1
			if rowInitialDirection == Desc {
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

func part2() {
	dataList := parseInputData()
	var safeReportsWithDampenerCount uint16 = 0
	for _, row := range dataList {
		skipFirstCheck := false
		rowErroneousValues := 0
		firstRowIssueIdx := -1
		rowInitialDirection := Asc

		firstDiff := (row[0] - row[1])
		secondDiff := 0

		if firstDiff == 0 {
			rowErroneousValues++
			firstRowIssueIdx = 0
			// Skip doing all the checks on the first number, we know it's bad already
			skipFirstCheck = true
			secondDiff = (row[1] - row[2])
			if secondDiff == 0 {
				// Skip row if first 3 numbers are the same
				continue
			}
		}

		if firstDiff > 0 || secondDiff > 0 {
			rowInitialDirection = Desc
		}

		for numIdx, num := range row {
			if skipFirstCheck && numIdx == 0 {
				continue
			}

			if (numIdx + 1) == len(row) {
				safeReportsWithDampenerCount++
			} else {
				nextNum := row[numIdx+1]
				diffWithNext := num - nextNum

				sequenceChangeValueIsInvalid := diffWithNext == 0 || IntAbs(diffWithNext) > 3

				sequenceChangeDirectionIsInvalid := diffWithNext > 0 && rowInitialDirection == Asc || diffWithNext < 0 && rowInitialDirection == Desc

				if sequenceChangeValueIsInvalid || sequenceChangeDirectionIsInvalid {
					// If the previous number is where the problem started, we try removing the current number and 
					// seeing if things play nice for the rest of the run  so that the current bad number doesn't
					// get double-counted
					if (numIdx != 0 && firstRowIssueIdx == numIdx-1) {
						endSkipDiff := row[numIdx-1] - row[numIdx+1]

						endSkipSequenceChangeValueIsInvalid := endSkipDiff == 0 || IntAbs(endSkipDiff) > 3
						endSkipSequenceChangeDirectionIsInvalid := endSkipDiff > 0 && rowInitialDirection == Asc || endSkipDiff < 0 && rowInitialDirection == Desc

						if endSkipSequenceChangeValueIsInvalid || endSkipSequenceChangeDirectionIsInvalid {
							rowErroneousValues++
						} else {
							fmt.Printf("DoubleChecked Row Passed! %v\n", row)
						}

					} else {

						rowErroneousValues++
					}
				}

				// If we have more than one erroneous number, its invalid
				if rowErroneousValues > 1 {
					break
				} else if (rowErroneousValues == 1 && firstRowIssueIdx == -1) {
					firstRowIssueIdx = numIdx
				}
			}
		}

	}

	fmt.Printf("The number of safe reports with Problem Dampener on is is: %v ", safeReportsWithDampenerCount)
}

func IntAbs(someint int) int {
	if someint < 0 {
		return -someint
	}

	return someint
}

func parseInputData() [][]int {
	file, err := os.Open("./input2.txt")

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
