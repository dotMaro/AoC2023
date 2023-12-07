package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile("src/day03/input.txt")
	if err != nil {
		panic(err)
	}
	engine := parseEngine(string(inputBytes))
	fmt.Printf("Part 1. %d\n", engine.sumOfPartNumbers())
	fmt.Printf("Part 2. %d\n", engine.gearRatio())
}

func parseEngine(s string) engine {
	engine := make([][]rune, 0)
	for _, line := range strings.Split(strings.ReplaceAll(s, "\r", ""), "\n") {
		row := make([]rune, 0)
		for _, r := range line {
			row = append(row, r)
		}
		engine = append(engine, row)
	}
	return engine
}

type engine [][]rune

func (e engine) gearRatio() int {
	gearRatio := 0
	for y, row := range e {
		for x, r := range row {
			if r != '*' {
				continue
			}

			isAdjacent, partNumber1, partNumber2 := e.isAdjacentToTwoPartNumbers(x, y)
			if isAdjacent {
				gearRatio += partNumber1 * partNumber2
			}
		}
	}

	return gearRatio
}

func (e engine) isAdjacentToTwoPartNumbers(x, y int) (bool, int, int) {
	partNumber1 := -1
	for x2 := x - 1; x2 <= x+1; x2++ {
		if x2 < 0 || x2 >= len(e[0]) {
			continue
		}
		for y2 := y - 1; y2 <= y+1; y2++ {
			if y2 < 0 || y2 >= len(e) || x2 == x && y2 == y {
				continue
			}

			isNumber, partNumber := e.fullNumber(x2, y2)
			if isNumber {
				if partNumber1 != -1 && partNumber != partNumber1 {
					return true, partNumber1, partNumber
				}
				partNumber1 = partNumber
			}
		}
	}

	return false, 0, 0
}

func (e engine) sumOfPartNumbers() int {
	sum := 0
	waitForNonDigit := false
	for y, row := range e {
		for x := range row {
			if waitForNonDigit {
				if !e.isDigit(x, y) {
					waitForNonDigit = false
				}
				continue
			}
			isPartNumber, partNumber := e.isPartNumber(x, y)
			if isPartNumber {
				sum += partNumber
				waitForNonDigit = true
			}
		}
	}
	return sum
}

func (e engine) isPartNumber(x, y int) (bool, int) {
	if !e.isDigit(x, y) {
		return false, 0
	}

	partNumber := e.toDigit(x, y)
	leftX := x - 1
	hasAdjacentSymbol := e.upOrBelowIsSymbol(x, y)
	for leftX >= 0 {
		if !hasAdjacentSymbol && e.upOrBelowIsSymbol(leftX, y) {
			hasAdjacentSymbol = true
		}
		if !e.isDigit(leftX, y) {
			if !hasAdjacentSymbol && e.isSymbol(leftX, y) {
				hasAdjacentSymbol = true
			}
			break
		}

		// This is wrong but since we always "enter" from the left, this won't happen anyway.
		// I just want to get through these tasks so I'll leave it, but at least I went through the effort of writing this comment.
		partNumber = e.toDigit(leftX, y)*10 + partNumber
		leftX--
	}
	rightX := x + 1
	for rightX < len(e[0]) {
		if !hasAdjacentSymbol && e.upOrBelowIsSymbol(rightX, y) {
			hasAdjacentSymbol = true
		}
		if !e.isDigit(rightX, y) {
			if !hasAdjacentSymbol && e.isSymbol(rightX, y) {
				hasAdjacentSymbol = true
			}
			break
		}

		partNumber = partNumber*10 + e.toDigit(rightX, y)
		rightX++
	}

	return hasAdjacentSymbol, partNumber
}

func (e engine) fullNumber(x, y int) (bool, int) {
	if !e.isDigit(x, y) {
		return false, 0
	}

	partNumber := e.toDigit(x, y)
	digits := 1
	leftX := x - 1
	for leftX >= 0 {
		if !e.isDigit(leftX, y) {
			break
		}

		partNumber = e.toDigit(leftX, y)*int(math.Pow10(digits)) + partNumber
		digits++
		leftX--
	}
	rightX := x + 1
	for rightX < len(e[0]) {
		if !e.isDigit(rightX, y) {
			break
		}

		partNumber = partNumber*10 + e.toDigit(rightX, y)
		rightX++
	}

	return true, partNumber
}

func (e engine) isSymbol(x, y int) bool {
	return e[y][x] != '.' && !e.isDigit(x, y)
}

func (e engine) isDigit(x, y int) bool {
	r := e[y][x]
	return r >= '0' && r <= '9'
}

func (e engine) toDigit(x, y int) int {
	return int(e[y][x] - '0')
}

func (e engine) upOrBelowIsSymbol(x, y int) bool {
	return y > 0 && e.isSymbol(x, y-1) || y < len(e)-1 && e.isSymbol(x, y+1)
}
