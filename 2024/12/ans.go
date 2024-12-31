package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x int
	y int
}

var values = make(map[point]rune)
var maxX, maxY int

func init() {
	input, err := os.ReadFile("example2.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for y, l := range lines {
		for x, r := range l {
			values[point{x, y}] = r
		}
	}
	// set bounds
	maxX, maxY = len(lines[0]), len(lines)
}

func isInBounds(p point) bool {
	return 0 <= p.x && p.x < maxX && 0 <= p.y && p.y < maxY
}

func getNeighbors(p point) (neighbors []point) {
	if north := (point{p.x, p.y - 1}); isInBounds(north) {
		neighbors = append(neighbors, north)
	}

	if south := (point{p.x, p.y + 1}); isInBounds(south) {
		neighbors = append(neighbors, south)
	}

	if east := (point{p.x - 1, p.y}); isInBounds(east) {
		neighbors = append(neighbors, east)
	}

	if west := (point{p.x + 1, p.y}); isInBounds(west) {
		neighbors = append(neighbors, west)
	}

	return neighbors
}

func getPointPerimeter(p point) (perimeter int) {
	neighbors := getNeighbors(p)
	for _, n := range neighbors {
		if values[n] != values[p] {
			perimeter += 1
		}
	}
	// Count out of bounds in the perimeter
	return perimeter + 4 - len(neighbors)
}

type pointset = map[point]struct{}

func addToSet(p point, s *pointset) {
	(*s)[p] = struct{}{}
}

func getAreaAndPerimeter(p point) (area, perimeter int, visited pointset) {
	visited = make(pointset)
	addToSet(p, &visited)
	// BFS
	queue := []point{p}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		addToSet(curr, &visited)

		perimeter += getPointPerimeter(curr)

		// Enqueue unvisited neighbors with the same value
		for _, n := range getNeighbors(curr) {
			if _, seen := visited[n]; !seen && values[n] == values[p] {
				queue = append(queue, n)
				addToSet(n, &visited)
			}
		}
	}
	return len(visited), perimeter, visited
}

/*
A "corner" either:
- Has exactly 2 neighbors (2 adj. points out of bounds -- literal corner)
- Has 2 neighbors with the same value
- Shares a non-same-valued neighbor with another point in its region
*/
func getAreaAndCorners(p point) (area, corners int, visited pointset) {
	panic("TODO")
	visited = make(pointset)
	addToSet(p, &visited)
	externalOverlapCounts := make(map[point]int)

	// BFS
	queue := []point{p}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		addToSet(curr, &visited)

		// Enqueue unvisited neighbors with the same value
		neighbors := getNeighbors(curr)
		numSimilar := 0
		for _, n := range neighbors {
			if _, seen := visited[n]; !seen && values[n] == values[p] {
				queue = append(queue, n)
				addToSet(n, &visited)
			}

			if values[n] == values[p] {
				numSimilar++
			} else {
				externalOverlapCounts[n] += 1
			}
		}

		// "Peninsula"
		if len(neighbors) == 4 && numSimilar == 1 {
			corners += 2

		}

		// // General corner
		// if len(neighbors) == 4 && numSimilar == 2 {
		// 	corners += 1
		// }

		// Outer corners
		if len(neighbors) == 2 && numSimilar == 2 {
			corners += 1
		}

		// e.g. Top and bottom fork of example "E"
		if len(neighbors) == 2 && numSimilar == 1 {
			corners += 2
		}

		// e.g. Middle fork of example "E"
		if len(neighbors) == 3 && numSimilar == 1 {
			corners += 2
		}

	}

	// Finds concave corners
	for _, count := range externalOverlapCounts {
		if count == 2 {
			corners += 1
		}
		if count == 3 {
			corners += 2
		}
	}

	return len(visited), corners, visited
}

func main() {
	globalVisited := make(pointset)
	totalPrice := 0
	for p := range values {
		if _, seen := globalVisited[p]; !seen {
			area, perimeter, visited := getAreaAndPerimeter(p)
			totalPrice += area * perimeter
			for v := range visited {
				addToSet(v, &globalVisited)
			}
		}
	}
	fmt.Println(totalPrice)

	// part 2
	globalVisited = make(pointset)
	totalPrice = 0
	for p := range values {
		if _, seen := globalVisited[p]; !seen {
			area, corners, visited := getAreaAndCorners(p)
			totalPrice += area * corners
			for v := range visited {
				addToSet(v, &globalVisited)
			}
			fmt.Println(area)
			fmt.Println(corners)
		}
	}
	fmt.Println(totalPrice)
}
