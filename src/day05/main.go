package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile("src/day05/input.txt")
	if err != nil {
		panic(err)
	}
	alamanac := parseAlamanac(string(inputBytes))
	fmt.Printf("Part 1. %d\n", alamanac.lowestLocation())
	fmt.Printf("Part 2. %d\n", alamanac.lowestLocationUsingSeedRanges())
}

func parseAlamanac(s string) alamanac {
	lines := strings.Split(strings.ReplaceAll(s, "\r", ""), "\n")
	seedString := lines[0][len("seeds: "):]
	var seeds []int
	for _, s := range strings.Split(seedString, " ") {
		seed, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, seed)
	}

	var maps []mapCategory
	var c mapCategory
	newMap := false
	for _, line := range lines[3:] {
		if newMap {
			// Skip map name line.
			newMap = false
			maps = append(maps, c)
			c = mapCategory{}
			continue
		}
		if line == "" {
			newMap = true
			continue
		}
		values := strings.SplitN(line, " ", 3)
		destinationStart, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		sourceStart, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		rangeLength, err := strconv.Atoi(values[2])
		if err != nil {
			panic(err)
		}
		r := mapRange{
			destinationStart: destinationStart,
			sourceStart:      sourceStart,
			rangeLength:      rangeLength,
		}
		c = append(c, r)
	}
	maps = append(maps, c)

	return alamanac{
		seeds:               seeds,
		maps:                maps,
		seedToLocationCache: make(map[int]int),
	}
}

type alamanac struct {
	seeds []int
	// Categories:
	// seedToSoil
	// soilToFertilizer
	// fertilizedToWater
	// waterToLight
	// lightToTemperature
	// temperatureToHumidity
	// humidityToLocation
	maps []mapCategory
	// Turns out this grows way too large.
	// There are certainly ways to make a more intelligent cache (maybe using ranges?).
	// In the end it only took around a minute to execute without a cache so I didn't bother.
	seedToLocationCache map[int]int
}

func (a *alamanac) lowestLocationUsingSeedRanges() int {
	lowest := math.MaxInt32
	for i := 0; i < len(a.seeds)-1; i += 2 {
		seedStart, seedRange := a.seeds[i], a.seeds[i+1]
		for s := seedStart; s <= seedStart+seedRange; s++ {
			if s%1000000 == 0 {
				fmt.Println(s)
			}
			// loc, inCache := a.seedToLocationCache[s]
			// if !inCache {
			loc := a.seedToLocation(s)
			// }
			// a.seedToLocationCache[s] = loc
			if loc < lowest {
				lowest = loc
			}
		}
	}
	return lowest
}

func (a *alamanac) lowestLocation() int {
	lowest := math.MaxInt32
	for _, s := range a.seeds {
		loc, inCache := a.seedToLocationCache[s]
		if !inCache {
			loc = a.seedToLocation(s)
		}
		a.seedToLocationCache[s] = loc
		if loc < lowest {
			lowest = loc
		}
	}
	return lowest
}

func (a *alamanac) seedToLocation(i int) int {
	lastVal := i
	for _, m := range a.maps {
		lastVal = m.convert(lastVal)
	}
	return lastVal
}

type mapCategory []mapRange

func (c mapCategory) convert(v int) int {
	for _, r := range c {
		if v >= r.sourceStart && v <= r.sourceStart+r.rangeLength {
			return r.destinationStart + v - r.sourceStart
		}
	}
	return v
}

type mapRange struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}
