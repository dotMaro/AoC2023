package main

import "testing"

const input = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func Test_engine_sumOfPartNumbers(t *testing.T) {
	engine := parseEngine(input)
	sum := engine.sumOfPartNumbers()
	if sum != 4361 {
		t.Errorf("Should return 4361, but returned %d", sum)
	}
}

func Test_engine_gearRatio(t *testing.T) {
	engine := parseEngine(input)
	gearRatio := engine.gearRatio()
	if gearRatio != 467835 {
		t.Errorf("Should return 467835, but returned %d", gearRatio)
	}
}
