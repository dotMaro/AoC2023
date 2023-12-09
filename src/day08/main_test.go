package main

import "testing"

const input1 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

func Test_nodeMap_traverseUntilZZZ(t *testing.T) {
	m, i := parseNodeMapAndInstructions(input1)
	res := m.traverseUntilZZZ(i)
	if res != 6 {
		t.Errorf("Should return 6, but returned %d", res)
	}
}

func Test_nodeMap_multitraverseUntilZ(t *testing.T) {
	const input = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

	m, i := parseNodeMapAndInstructions(input)
	res := m.multitraverseUntilZ(i)
	if res != 6 {
		t.Errorf("Should return 6, but returned %d", res)
	}
}

func Test_nodeIndexMap_multitraverseUntilZ(t *testing.T) {
	const input = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

	m, i := parseNodeIndexMapAndInstructions(input)
	res := m.multitraverseUntilZ(i)
	if res != 6 {
		t.Errorf("Should return 6, but returned %d", res)
	}
}
