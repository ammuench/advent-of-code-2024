package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var sourceData = parseInputData()
var xWestBoundary = 2
var yNorthBoundary = 2
var xEastBoundary = len(sourceData[0]) - 3
var ySouthBoundary = len(sourceData) - 3

func main() {
	part1()
	part2()
}

func part1() {
	totalMatches := 0
	totalDiag := 0
	totalCard := 0
	for xIdx, row := range sourceData {
		for yIdx, char := range row {
			if char == "X" {
				diagonalMatches := checkDiagonals(xIdx, yIdx)
				cardinalMatches := checkCardinals(xIdx, yIdx)
				totalCard += cardinalMatches
				totalDiag += diagonalMatches
				totalMatches += diagonalMatches + cardinalMatches
			}
		}
	}

	fmt.Printf("Total diag matches:%v || Total cardinalMatches: %v \n", totalDiag, totalCard)

	fmt.Printf("Part One total matches ==%v\n", totalMatches)

}

func part2() {
}

func checkCardinals(x, y int) int {
	cardinalMatches := 0

	if x > xWestBoundary && hasWestMatch(x, y) {
		cardinalMatches++
	}
	if y > yNorthBoundary && hasNorthMatch(x, y) {
		cardinalMatches++
	}

	if x < xEastBoundary && hasEastMatch(x, y) {
		cardinalMatches++
	}

	if y < ySouthBoundary && hasSouthMatch(x, y) {
		cardinalMatches++
	}

	return cardinalMatches
}

func hasNorthMatch(x, y int) bool {
	if sourceData[x][y-1] == "M" && sourceData[x][y-2] == "A" && sourceData[x][y-3] == "S" {
		return true
	}

	return false
}
func hasSouthMatch(x, y int) bool {
	if sourceData[x][y+1] == "M" && sourceData[x][y+2] == "A" && sourceData[x][y+3] == "S" {
		return true
	}

	return false
}
func hasEastMatch(x, y int) bool {
	if sourceData[x+1][y] == "M" && sourceData[x+2][y] == "A" && sourceData[x+3][y] == "S" {
		return true
	}

	return false
}
func hasWestMatch(x, y int) bool {
	if sourceData[x-1][y] == "M" && sourceData[x-2][y] == "A" && sourceData[x-3][y] == "S" {
		return true
	}

	return false
}

func checkDiagonals(x, y int) int {
	diagonalMatches := 0

	if x > xWestBoundary {
		if y > yNorthBoundary && hasDiagonalNWMatch(x, y) {
			diagonalMatches++
		}

		if y < ySouthBoundary && hasDiagonalSWMatch(x, y) {
			diagonalMatches++
		}
	}

	if x < xEastBoundary {
		if y < ySouthBoundary && hasDiagonalSEMatch(x, y) {
			diagonalMatches++
		}
		if y > yNorthBoundary && hasDiagonalNEMatch(x, y) {
			diagonalMatches++
		}
	}

	return diagonalMatches
}

func hasDiagonalNWMatch(x, y int) bool {
	if sourceData[x-1][y-1] == "M" && sourceData[x-2][y-2] == "A" && sourceData[x-3][y-3] == "S" {
		return true
	}

	return false
}

func hasDiagonalSWMatch(x, y int) bool {
	if sourceData[x-1][y+1] == "M" && sourceData[x-2][y+2] == "A" && sourceData[x-3][y+3] == "S" {
		return true
	}

	return false

}
func hasDiagonalSEMatch(x, y int) bool {
	if sourceData[x+1][y+1] == "M" && sourceData[x+2][y+2] == "A" && sourceData[x+3][y+3] == "S" {
		return true
	}

	return false

}
func hasDiagonalNEMatch(x, y int) bool {
	if sourceData[x+1][y-1] == "M" && sourceData[x+2][y-2] == "A" && sourceData[x+3][y-3] == "S" {
		return true
	}

	return false

}

func parseInputData() [][]string {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	parsedData := make([]([]string), 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parsedStrings := strings.Split(scanner.Text(), "")

		parsedData = append(parsedData, parsedStrings)

	}

	return parsedData
}
