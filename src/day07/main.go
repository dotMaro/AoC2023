package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile("src/day07/input.txt")
	if err != nil {
		panic(err)
	}
	hands := parseHands(string(inputBytes), false)
	fmt.Printf("Part 1. %d\n", winnings(hands))
	handsWithJokerWildcards := parseHands(string(inputBytes), true)
	fmt.Printf("Part 2. %d\n", winnings(handsWithJokerWildcards))
}

func winnings(hands []hand) int {
	sort.SliceStable(hands, func(i, j int) bool {
		return hands[j].greaterThan(hands[i])
	})
	winnings := 0
	for i, h := range hands {
		rank := i + 1
		winnings += rank * h.bid
	}
	return winnings
}

func parseHands(s string, jokerWildcards bool) []hand {
	var hands []hand
	for _, line := range strings.Split(strings.ReplaceAll(s, "\r", ""), "\n") {
		hands = append(hands, parseHand(line, jokerWildcards))
	}
	return hands
}

func parseHand(s string, jokerWildcards bool) hand {
	var cards cards
	for i, r := range s[:5] {
		cards[i] = newCardLabel(r)
	}
	bid, err := strconv.Atoi(s[6:])
	if err != nil {
		panic(err)
	}
	return hand{
		cards:         cards,
		bid:           bid,
		handType:      cards.handType(jokerWildcards),
		jokerWildcard: jokerWildcards,
	}
}

type hand struct {
	cards         cards
	bid           int
	handType      handType
	jokerWildcard bool
}

func (h hand) String() string {
	var b strings.Builder
	for _, c := range h.cards {
		b.WriteRune(rune(c))
	}
	return b.String()
}

func (h hand) greaterThan(o hand) bool {
	if h.handType > o.handType {
		return true
	}
	if h.handType < o.handType {
		return false
	}

	for i, c := range h.cards {
		if c.greaterThan(o.cards[i], h.jokerWildcard) {
			return true
		}
		if o.cards[i].greaterThan(c, h.jokerWildcard) {
			return false
		}
	}
	panic("Hands shouldn't be equal")
}

type cards [5]cardLabel

func (c cards) handType(jokerWildcards bool) handType {
	distr := make(map[cardLabel]uint8, 5)
	for _, card := range c {
		distr[card]++
	}
	jokerCount, hasJoker := distr[labelJ]
	if jokerWildcards && hasJoker && jokerCount < 5 {
		var highestCount uint8
		var highestCard cardLabel
		for card, count := range distr {
			if card != labelJ && count > highestCount {
				highestCount = count
				highestCard = card
			}
		}
		distr[highestCard] += jokerCount
		delete(distr, labelJ)
	}
	switch len(distr) {
	case 1:
		return fiveOfAKind
	case 2:
		for _, d := range distr {
			switch d {
			case 4, 1:
				return fourOfAKind
			case 3, 2:
				return fullHouse
			}
		}
	case 3:
		for _, d := range distr {
			if d == 3 {
				return threeOfAKind
			}
		}
		return twoPair
	case 4:
		return onePair
	case 5:
		return highCard
	}
	panic("Invalid distribution")
}

type handType uint8

const (
	highCard handType = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func newCardLabel(r rune) cardLabel {
	switch r {
	case 'A':
		return labelA
	case 'K':
		return labelK
	case 'Q':
		return labelQ
	case 'J':
		return labelJ
	case 'T':
		return labelT
	case '9':
		return label9
	case '8':
		return label8
	case '7':
		return label7
	case '6':
		return label6
	case '5':
		return label5
	case '4':
		return label4
	case '3':
		return label3
	case '2':
		return label2
	default:
		panic("Invalid card label rune")
	}
}

type cardLabel uint8

const (
	labelA cardLabel = iota
	labelK
	labelQ
	labelJ
	labelT
	label9
	label8
	label7
	label6
	label5
	label4
	label3
	label2
)

func (l cardLabel) greaterThan(o cardLabel, jokerWildcard bool) bool {
	if jokerWildcard {
		if l == labelJ {
			return false
		}
		if o == labelJ {
			return true
		}
	}
	return l < o
}
