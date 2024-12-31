package main

import (
	"fmt"
	"os"
	"strings"
)

func getInput() []string {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	return lines
}

func getDimensions(input []string) [2]int {
	return [2]int{len(input[0]), len(input)}
}

type node struct {
	x   int
	y   int
	val rune
}

/*
		x
	y	[(0,0),	..., (n, 0),




		(0, n),	..., (n, n)]
*/

func getNodes(input []string) (nodes []node) {
	for y, line := range input {
		for x, r := range line {
			if r != '.' {
				nodes = append(nodes, node{x, y, r})
			}
		}
	}
	return nodes
}

type nodeset = map[node]struct{}

func addNodeToSet(n node, s nodeset) nodeset {
	s[n] = struct{}{}
	return s
}

func getNodeSet(nodes []node) (s nodeset) {
	s = make(nodeset)
	for _, n := range nodes {
		s = addNodeToSet(n, s)
	}
	return s
}

func mergeSets(a, b nodeset) nodeset {
	for n := range b {
		a[n] = struct{}{}
	}
	return a
}

func isInBounds(n node, dims [2]int) bool {
	xMax, yMax := dims[0], dims[1]
	x, y := n.x, n.y
	return 0 <= x && x < xMax && 0 <= y && y < yMax
}

func normalizePair(a, b node) [2]node {
	if a.x > b.x || (a.x == b.x && a.y > b.y) {
		return [2]node{b, a}
	}
	return [2]node{a, b}
}

/*
		x
	y	[(0,0),	..., (n, 0),

		(a-(c-a), b-(d-b)) <- antinode


			(a, b)


				(c, d)


					(c + (c-a), d + (d-b)) <- antinode

		(0, n),	..., (n, n)]
*/

func getAntinodes(a, b node, dims [2]int) (antinodes nodeset) {
	antinodes = make(nodeset)

	pair := normalizePair(a, b)
	n1, n2 := pair[0], pair[1]
	dx, dy := n2.x-n1.x, n2.y-n1.y

	n1 = node{n1.x - dx, n1.y - dy, '#'}
	if isInBounds(n1, dims) {
		antinodes = addNodeToSet(n1, antinodes)

	}

	n2 = node{n2.x + dx, n2.y + dy, '#'}
	if isInBounds(n2, dims) {
		antinodes = addNodeToSet(n2, antinodes)
	}

	return antinodes
}

type pairset = map[[2]node]struct{}

func addPair(a, b node, s pairset) pairset {
	pair := normalizePair(a, b)
	s[pair] = struct{}{}
	return s
}

func hasPair(a, b node, s pairset) bool {
	pair := normalizePair(a, b)
	_, ok := s[pair]
	return ok
}

func getAllAntinodes(s nodeset, dims [2]int) (antinodes nodeset) {
	antinodes = make(nodeset)
	pairs := make(pairset)

	for a := range s {
		for b := range s {
			if (a != b) && !hasPair(a, b, pairs) {
				antinodes = mergeSets(antinodes, getAntinodes(a, b, dims))
				pairs = addPair(a, b, pairs)
			}
		}
	}
	return antinodes
}

type nodegroups = map[rune]nodeset

func groupNodesByValue(s nodeset) (groups nodegroups) {
	groups = make(nodegroups)
	for n := range s {
		_, ok := groups[n.val]
		if !ok {
			groups[n.val] = make(nodeset)
		}
		groups[n.val] = addNodeToSet(n, groups[n.val])
	}
	return groups
}

// don't ask why the signs switched
func getExtendedAntinodes(a, b node, dims [2]int) (antinodes nodeset) {
	antinodes = make(nodeset)

	pair := normalizePair(a, b)
	n1, n2 := pair[0], pair[1]
	dx, dy := n2.x-n1.x, n2.y-n1.y

	n1 = node{n1.x + dx, n1.y + dy, '#'}
	for isInBounds(n1, dims) {
		antinodes = addNodeToSet(n1, antinodes)
		n1 = node{n1.x + dx, n1.y + dy, '#'}
	}

	n2 = node{n2.x - dx, n2.y - dy, '#'}
	for isInBounds(n2, dims) {
		antinodes = addNodeToSet(n2, antinodes)
		n2 = node{n2.x - dx, n2.y - dy, '#'}
	}
	return antinodes
}

func getAllExtendedAntinodes(s nodeset, dims [2]int) (antinodes nodeset) {
	antinodes = make(nodeset)
	pairs := make(pairset)

	for a := range s {
		for b := range s {
			if (a != b) && !hasPair(a, b, pairs) {
				antinodes = mergeSets(antinodes, getExtendedAntinodes(a, b, dims))
				pairs = addPair(a, b, pairs)
			}
		}
	}
	return antinodes
}

func main() {
	input := getInput()
	dims := getDimensions(input)
	nodes := getNodes(input)
	s := getNodeSet(nodes)
	g := groupNodesByValue(s)

	// part 1
	antinodes := make(nodeset)
	for _, s := range g {
		antinodes = mergeSets(antinodes, getAllAntinodes(s, dims))
	}
	fmt.Println(len(antinodes))

	// part 2
	antinodes = make(nodeset)
	for _, s := range g {
		antinodes = mergeSets(antinodes, getAllExtendedAntinodes(s, dims))
	}
	fmt.Println(len(antinodes))

}
