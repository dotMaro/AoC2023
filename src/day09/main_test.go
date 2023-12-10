package main

import "testing"

const input = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func Test_sumOfExtrapolations(t *testing.T) {
	h := parseSensorHistories(input)
	res := sumOfExtrapolations(h)
	if res != 114 {
		t.Errorf("Should return 114, but returned %d", res)
	}
}
