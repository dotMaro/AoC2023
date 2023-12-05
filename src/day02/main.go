package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile("src/day02/input.txt")
	if err != nil {
		panic(err)
	}
	games := parseGames(string(inputBytes))
	fmt.Printf("Part 1. %d", sumOfPossibleGames(games, drawnSet{drawnCubes{color: "red", count: 12}, drawnCubes{color: "green", count: 13}, drawnCubes{color: "blue", count: 14}}))
}

func parseGames(s string) []game {
	games := make([]game, 0)
	for _, line := range strings.Split(strings.ReplaceAll(s, "\r", ""), "\n") {
		games = append(games, parseGame(line))
	}
	return games
}

func parseGame(s string) game {
	split := strings.SplitN(s, ":", 2)
	id, err := strconv.Atoi(split[0][len("Game "):])
	if err != nil {
		panic(err)
	}

	sets := make([]drawnSet, 0)
	for _, set := range strings.Split(split[1], ";") {
		cubes := make([]drawnCubes, 0)
		for _, cube := range strings.Split(set, ", ") {
			countAndColor := strings.Split(strings.TrimSpace(cube), " ")
			count, err := strconv.Atoi(countAndColor[0])
			if err != nil {
				panic(err)
			}
			c := drawnCubes{
				color: countAndColor[1],
				count: count,
			}
			cubes = append(cubes, c)
		}
		sets = append(sets, cubes)
	}
	return game{
		id:        id,
		drawnSets: sets,
	}
}

type game struct {
	id        int
	drawnSets []drawnSet
}

func sumOfPossibleGames(games []game, d drawnSet) int {
	sum := 0
	for _, g := range games {
		allPossible := true
		for _, s := range g.drawnSets {
			if !s.isPossibleWithDrawnCubes(d) {
				allPossible = false
				break
			}
		}
		if allPossible {
			sum += g.id
		}
	}
	return sum
}

type drawnSet []drawnCubes

func (s drawnSet) isPossibleWithDrawnCubes(o drawnSet) bool {
	for _, c := range s {
		for _, d := range o {
			if c.color == d.color {
				if d.count <= c.count {
					return false
				}
				break
			}
		}
	}
	return true
}

type drawnCubes struct {
	color string
	count int
}
