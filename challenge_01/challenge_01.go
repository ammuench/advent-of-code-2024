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
	part2()
}

func part2() {
    
	leftCol, rightCol := parseInputData()

	totalSimilarity := 0

	for _, leftVal := range leftCol {
        currentValSimilarityVal := 0
        for _, rightVal := range rightCol {
            if leftVal == rightVal {
                 currentValSimilarityVal++
            }
        }
        totalSimilarity += (currentValSimilarityVal * leftVal)
	}

	fmt.Printf("The similarity value is: %v", totalSimilarity)
}

func part1() {
	leftCol, rightCol := parseInputData()

	totalDistance := 0

	for idx, leftVal := range leftCol {
		rightVal := rightCol[idx]

		distanceVal := rightVal - leftVal

		if distanceVal < 0 {
			distanceVal = distanceVal * -1
		}

		totalDistance += distanceVal
	}

	fmt.Printf("The list distance value is: %v", totalDistance)

}

func parseInputData() ([]int, []int) {
	// OPEN FILE
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// CLOSE FILE STREAM WHEN FUNC OVER
	defer file.Close()

	leftCol := make([]int, 0)
	rightCol := make([]int, 0)

	const ColSplitCharacter = "   "

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		colSplits := strings.Split(scanner.Text(), ColSplitCharacter)

		if len(colSplits) > 1 {

			leftColString := colSplits[0]
			rightColString := colSplits[1]

			parsedLeftCol, err := strconv.Atoi(leftColString)
			parsedRightCol, err := strconv.Atoi(rightColString)
			if err != nil {
				log.Fatal(err)
			} else {
				leftCol = append(leftCol, parsedLeftCol)
				rightCol = append(rightCol, parsedRightCol)
			}
		}

	}

	// sort the lists
	slices.SortFunc(leftCol, func(a, b int) int {
		if a < b {
			return -1
		}

		if a > b {
			return 1
		}

		return 0
	})
	slices.SortFunc(rightCol, func(a, b int) int {
		if a < b {
			return -1
		}

		if a > b {
			return 1
		}

		return 0
	})

	return leftCol, rightCol

}
