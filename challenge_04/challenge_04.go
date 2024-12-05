package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var sourceData = parseInputData()
var xWestBoundaryPart1 = 2
var yNorthBoundaryPart1 = 2
var xEastBoundaryPart1 = len(sourceData[0]) - 3
var ySouthBoundaryPart1 = len(sourceData) - 3
var xWestBoundaryPart2 = 0
var yNorthBoundaryPart2 = 0
var xEastBoundaryPart2 = len(sourceData[0]) - 1
var ySouthBoundaryPart2 = len(sourceData) - 1

func main() {
	part1()
	part2()
}

func part2() {
	totalMatches := 0
	for xIdx, row := range sourceData {
		for yIdx, char := range row {
			if char == "A" {
				crossMasMatches := checkMasDiagonals(xIdx, yIdx)
				totalMatches += crossMasMatches
			}
		}
	}

	fmt.Printf("Part Two total matches ==%v\n", totalMatches)
}

func checkMasDiagonals(x, y int) int {
	diagonalMatches := 0

	if x > xWestBoundaryPart2 && x < xEastBoundaryPart2 && y < ySouthBoundaryPart2 && y > yNorthBoundaryPart2 {
		nsMatch := hasMasNSMatch(x, y)
		snMatch := hasMasSNMatch(x, y)

		if nsMatch && snMatch {
			diagonalMatches++
		}
	}

	return diagonalMatches
}

func hasMasNSMatch(x, y int) bool {
	if sourceData[x-1][y-1] == "M" && sourceData[x+1][y+1] == "S" {
		return true
	}

	if sourceData[x-1][y-1] == "S" && sourceData[x+1][y+1] == "M" {
		return true
	}

	return false
}

func hasMasSNMatch(x, y int) bool {
	if sourceData[x-1][y+1] == "M" && sourceData[x+1][y-1] == "S" {
		return true
	}

	if sourceData[x-1][y+1] == "S" && sourceData[x+1][y-1] == "M" {
		return true
	}

	return false
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

func checkCardinals(x, y int) int {
	cardinalMatches := 0

	if x > xWestBoundaryPart1 && hasWestMatch(x, y) {
		cardinalMatches++
	}
	if y > yNorthBoundaryPart1 && hasNorthMatch(x, y) {
		cardinalMatches++
	}

	if x < xEastBoundaryPart1 && hasEastMatch(x, y) {
		cardinalMatches++
	}

	if y < ySouthBoundaryPart1 && hasSouthMatch(x, y) {
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

	if x > xWestBoundaryPart1 {
		if y > yNorthBoundaryPart1 && hasDiagonalNWMatch(x, y) {
			diagonalMatches++
		}

		if y < ySouthBoundaryPart1 && hasDiagonalSWMatch(x, y) {
			diagonalMatches++
		}
	}

	if x < xEastBoundaryPart1 {
		if y < ySouthBoundaryPart1 && hasDiagonalSEMatch(x, y) {
			diagonalMatches++
		}
		if y > yNorthBoundaryPart1 && hasDiagonalNEMatch(x, y) {
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
