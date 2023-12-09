package main

import (
	"testing"
)

const input = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func Test_winnings(t *testing.T) {
	hands := parseHands(input, true)
	res := winnings(hands)
	if res != 5905 {
		t.Errorf("Should return 5905, but returned %d", res)
	}
}
