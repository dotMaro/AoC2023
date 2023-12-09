package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = `Time:        34     90     89     86
Distance:   204   1713   1210   1780`

func main() {
	contests := parseContests(input)
	fmt.Printf("Part 1. %d\n", productOfRanges(contests))
	contest := contest{
		time:           34908986,
		recordDistance: 204171312101780,
	}
	fmt.Printf("Part 2. %d\n", contest.rangeBeatingRecord())
}

func productOfRanges(contests []contest) int {
	product := 1
	for _, c := range contests {
		product *= c.rangeBeatingRecord()
	}
	return product
}

func parseContests(s string) []contest {
	lines := strings.SplitN(strings.ReplaceAll(s, "\r", ""), "\n", 2)
	var times []int
	for _, t := range strings.Split(lines[0][len("Time:"):], " ") {
		if t == "" {
			continue
		}
		time, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		times = append(times, time)
	}
	var distances []int
	for _, d := range strings.Split(lines[1][len("Distance:"):], " ") {
		if d == "" {
			continue
		}
		distance, err := strconv.Atoi(d)
		if err != nil {
			panic(err)
		}
		distances = append(distances, distance)
	}
	var contests []contest
	for i := range times {
		contest := contest{
			time:           times[i],
			recordDistance: distances[i],
		}
		contests = append(contests, contest)
	}
	return contests
}

type contest struct {
	time           int
	recordDistance int
}

func (c contest) rangeBeatingRecord() int {
	count := 0
	// You can definitely either calculate this artithmetically or at least use binary search or the like.
	// But this is fast enough so once again, leaving it.
	for i := 1; i < c.time; i++ {
		distance := i * (c.time - i)
		if distance > c.recordDistance {
			count++
		}
	}
	return count
}
