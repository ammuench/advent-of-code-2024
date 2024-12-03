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
		rowInitialDirection := Asc

		firstDiff := (row[0] - row[1])
		secondDiff := 0

		if firstDiff == 0 {
			rowErroneousValues++
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
				fmt.Println(row)
				safeReportsWithDampenerCount++
			} else {
				nextNum := row[numIdx+1]
				diffWithNext := num - nextNum

				sequenceChangeValueIsInvalid := diffWithNext == 0 || IntAbs(diffWithNext) > 3

				sequenceChangeDirectionIsInvalid := diffWithNext > 0 && rowInitialDirection == Asc || diffWithNext < 0 && rowInitialDirection == Desc

				if sequenceChangeValueIsInvalid || sequenceChangeDirectionIsInvalid {
					if numIdx == len(row)-2 {
						if rowErroneousValues == 0 {
							// If only error occurs on second to last number of the set, we know we can drop the last and it will be good
							continue
						} else {
							break
						}
					} else {
						// If we triggered a failure, see if skipping and trying the next option fixes it.  If it doesn't we know the set is bad
						skipDiff := num - row[numIdx+2]
						skipSequenceChangeValueIsInvalid := skipDiff == 0 || IntAbs(skipDiff) > 3
						skipSequenceChangeDirectionIsInvalid := skipDiff > 0 && rowInitialDirection == Asc || skipDiff < 0 && rowInitialDirection == Desc

						if skipSequenceChangeValueIsInvalid || skipSequenceChangeDirectionIsInvalid {
							break
						} else {
							rowErroneousValues++
						}
					}

				}

				// If we have more than one erroneous number, its invalid
				if rowErroneousValues > 1 {
					break
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
