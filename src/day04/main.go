package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile("src/day04/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1. %d\n", sumOfCardPoints(string(inputBytes)))
	fmt.Printf("Part 2. %d\n", totalScratchCardCount(string(inputBytes)))
}

func totalScratchCardCount(s string) int {
	count := 0
	lines := strings.Split(strings.ReplaceAll(s, "\r", ""), "\n")
	obtainedCards := make([]int, len(lines))
	for cardID, line := range lines {
		curCardCount := obtainedCards[cardID] + 1
		winningNumbersCount := cardWinningNumbers(line)
		for i := cardID + 1; i <= cardID+winningNumbersCount && i < len(obtainedCards); i++ {
			obtainedCards[i] += curCardCount
		}
		count += curCardCount
	}
	return count
}

func sumOfCardPoints(s string) int {
	sum := 0
	for _, line := range strings.Split(strings.ReplaceAll(s, "\r", ""), "\n") {
		sum += winningNumbersToPoints(cardWinningNumbers(line))
	}
	return sum
}

func cardWinningNumbers(s string) int {
	winningAndCardNumbers := strings.SplitN(s[len("Card nnn: "):], " | ", 2)
	winningNumbers := make(map[int]struct{})
	for _, w := range strings.Split(winningAndCardNumbers[0], " ") {
		// There are double spaces for single digit numbers.
		if w == "" {
			continue
		}

		winningNumber, err := strconv.Atoi(w)
		if err != nil {
			panic(err)
		}
		winningNumbers[winningNumber] = struct{}{}
	}

	var cardNumbers []int
	for _, w := range strings.Split(winningAndCardNumbers[1], " ") {
		// There are double spaces for single digit numbers.
		if w == "" {
			continue
		}

		cardNumber, err := strconv.Atoi(w)
		if err != nil {
			panic(err)
		}
		cardNumbers = append(cardNumbers, cardNumber)
	}

	amountOfWinningNumbers := 0
	for _, n := range cardNumbers {
		_, isWinningNumber := winningNumbers[n]
		if isWinningNumber {
			amountOfWinningNumbers++
		}
	}
	return amountOfWinningNumbers
}

func winningNumbersToPoints(n int) int {
	if n == 0 {
		return 0
	}
	points := 1
	for i := 1; i < n; i++ {
		points *= 2
	}
	return points
}
