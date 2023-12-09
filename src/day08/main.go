package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile("src/day08/input.txt")
	if err != nil {
		panic(err)
	}
	nodeMap, instructions := parseNodeMapAndInstructions(string(inputBytes))
	fmt.Printf("Part 1. %v\n", nodeMap.traverseUntilZZZ(instructions))
	nodeIndexMap, _ := parseNodeIndexMapAndInstructions(string(inputBytes))
	fmt.Printf("Part 2. %v\n", nodeIndexMap.multitraverseUntilZ(instructions))
}

func parseNodeMapAndInstructions(s string) (nodeMap, string) {
	m := make(nodeMap)
	lines := strings.Split(strings.ReplaceAll(s, "\r", ""), "\n")
	for _, l := range lines[2:] {
		m[l[:3]] = node{
			left:  l[7:10],
			right: l[12:15],
		}
	}
	return m, lines[0]
}

func (m nodeMap) traverseUntilZZZ(instructions string) int {
	curNode := "AAA"
	count := 0
	for curNode != "ZZZ" {
		instruction := instructions[count%len(instructions)]
		switch instruction {
		case 'L':
			curNode = m[curNode].left
		case 'R':
			curNode = m[curNode].right
		default:
			panic("invalid instruction")
		}
		count++
	}
	return count
}

func (m nodeMap) multitraverseUntilZ(instructions string) int {
	var curNodes []string
	for n := range m {
		if strings.HasSuffix(n, "A") {
			curNodes = append(curNodes, n)
		}
	}
	count := 0
	for {
		allNodesEndInZ := true
		for _, n := range curNodes {
			if !strings.HasSuffix(n, "Z") {
				allNodesEndInZ = false
				break
			}
		}
		if allNodesEndInZ {
			break
		}

		instruction := instructions[count%len(instructions)]

		for i, n := range curNodes {
			switch instruction {
			case 'L':
				curNodes[i] = m[n].left
			case 'R':
				curNodes[i] = m[n].right
			default:
				panic("invalid instruction")
			}
		}

		count++
		if count%1000000 == 0 {
			fmt.Println(count)
		}
	}
	return count
}

type nodeMap map[string]node

type node struct {
	left, right string
}

func parseNodeIndexMapAndInstructions(s string) (nodeIndexMap, string) {
	lines := strings.Split(strings.ReplaceAll(s, "\r", ""), "\n")
	m := make([]nodeIndex, len(lines)-2)
	nodeIndexes := make(map[string]int)
	for i, l := range lines[2:] {
		nodeIndexes[l[:3]] = i
	}
	for i, l := range lines[2:] {
		m[i] = nodeIndex{
			name:  l[:3],
			left:  nodeIndexes[l[7:10]],
			right: nodeIndexes[l[12:15]],
		}
	}
	return m, lines[0]
}

type nodeIndexMap []nodeIndex

type nodeIndex struct {
	name        string
	left, right int
}

func (m nodeIndexMap) multitraverseUntilZ(instructions string) int {
	var curNodes []nodeIndex
	for _, n := range m {
		if strings.HasSuffix(n.name, "A") {
			curNodes = append(curNodes, n)
		}
	}
	step := 0
	atZIn := make([]int, len(curNodes))
	for {
		reachedAllZNodes := true
		for _, v := range atZIn {
			if v == 0 {
				reachedAllZNodes = false
				break
			}
		}
		if reachedAllZNodes {
			break
		}

		instruction := instructions[step%len(instructions)]

		for i, n := range curNodes {
			alreadyAtZ := atZIn[i] != 0
			if alreadyAtZ {
				continue
			}
			if strings.HasSuffix(n.name, "Z") {
				atZIn[i] = step
			}
			switch instruction {
			case 'L':
				curNodes[i] = m[n.left]
			case 'R':
				curNodes[i] = m[n.right]
			default:
				panic("invalid instruction")
			}
		}

		step++
	}

	return lcm(atZIn[0], atZIn[1], atZIn[2:]...)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
