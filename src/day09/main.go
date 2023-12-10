package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile("src/day09/input.txt")
	if err != nil {
		panic(err)
	}
	sensorHistories := parseSensorHistories(string(inputBytes))
	fmt.Printf("Part 1. %d\n", sumOfExtrapolations(sensorHistories))
	fmt.Printf("Part 2. %d\n", sumOfBackwardsExtrapolations(sensorHistories))
}

func sumOfExtrapolations(histories []sensorHistory) int {
	sum := 0
	for _, h := range histories {
		sum += h.extrapolate()
	}
	return sum
}

func sumOfBackwardsExtrapolations(histories []sensorHistory) int {
	sum := 0
	for _, h := range histories {
		sum += h.extrapolateBackwards()
	}
	return sum
}

func parseSensorHistories(s string) []sensorHistory {
	var h []sensorHistory
	for _, l := range strings.Split(strings.ReplaceAll(s, "\r", ""), "\n") {
		h = append(h, parseSensorHistory(l))
	}
	return h
}

func parseSensorHistory(s string) sensorHistory {
	var h []int
	for _, v := range strings.Split(s, " ") {
		v2, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		h = append(h, v2)
	}
	return h
}

type sensorHistory []int

func (h sensorHistory) makeDiffStack() [][]int {
	stack := make([][]int, 1)
	stack[0] = make([]int, len(h))
	copy(stack[0], h)

	allZeroes := false
	for !allZeroes {
		diff := make([]int, 0, len(h)-1)
		lastStack := stack[len(stack)-1]
		for i := 1; i < len(lastStack); i++ {
			diff = append(diff, lastStack[i]-lastStack[i-1])
		}
		stack = append(stack, diff)

		allZeroes = true
		for _, d := range diff {
			if d != 0 {
				allZeroes = false
				break
			}
		}
	}
	return stack
}

func (h sensorHistory) extrapolate() int {
	stack := h.makeDiffStack()

	for i := len(stack) - 2; i >= 0; i-- {
		curStack := stack[i]
		lastStack := stack[i+1]
		stack[i] = append(stack[i], curStack[len(curStack)-1]+lastStack[len(lastStack)-1])
	}

	return stack[0][len(stack[0])-1]
}

func (h sensorHistory) extrapolateBackwards() int {
	stack := h.makeDiffStack()

	for i := len(stack) - 2; i >= 0; i-- {
		curStack := stack[i]
		lastStack := stack[i+1]
		stack[i] = append([]int{curStack[0] - lastStack[0]}, stack[i]...)
	}

	return stack[0][0]
}
