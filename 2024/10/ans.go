package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() [][]int {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	var points [][]int
	for _, line := range lines {
		var row []int
		for _, val := range line {
			num, err := strconv.Atoi(string(val))
			if err != nil {
				panic(err)
			}
			row = append(row, num)
		}
		points = append(points, row)
	}
	return points
}

type point = [2]int

func getTrailheads(input *[][]int) (trailheads []point) {
	for y, line := range *input {
		for x, val := range line {
			if val == 0 {
				trailheads = append(trailheads, point{x, y})
			}
		}
	}
	return trailheads
}

type direction = [2]int

var directions = [4]direction{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func getVal(input *[][]int, p point) int {
	return (*input)[p[1]][p[0]]
}

func isInBounds(input *[][]int, p point) bool {
	return 0 <= p[0] && p[0] < len(*input) && 0 <= p[1] && p[1] < len((*input)[0])
}

func move(p point, d direction) point {
	return point{p[0] + d[0], p[1] + d[1]}
}

func isValidMove(input *[][]int, p point, d direction) bool {
	next := move(p, d)
	return isInBounds(input, next) && (getVal(input, next) == (getVal(input, p) + 1))
}

type pointset = map[point]interface{}

func collectEndpoints(input *[][]int, p point, endpoints *pointset) {
	if !isInBounds(input, p) {
		return
	}

	if getVal(input, p) >= 9 {
		(*endpoints)[p] = struct{}{}
		return
	}

	for _, dir := range directions {
		if isValidMove(input, p, dir) {
			collectEndpoints(input, move(p, dir), endpoints)
		}
	}
	return
}

func getScore(input *[][]int, p point) int {
	endpoints := pointset{}
	collectEndpoints(input, p, &endpoints)
	return len(endpoints)
}

// part 2 funcs

type path = [10]point
type pathset = map[path]interface{}

func collectFullPaths(input *[][]int, p point, paths *pathset, current path) {
	if !isInBounds(input, p) {
		return
	}

	val := getVal(input, p)
	current[val] = p

	if getVal(input, p) >= 9 {
		(*paths)[current] = struct{}{}
		return
	}

	for _, dir := range directions {
		if isValidMove(input, p, dir) {
			collectFullPaths(input, move(p, dir), paths, current)
		}
	}
	return
}

func getRating(input *[][]int, p point) int {
	paths := pathset{}
	collectFullPaths(input, p, &paths, path{})
	return len(paths)
}

func main() {
	input := getInput()
	trailheads := getTrailheads(&input)
	totalScore := 0
	for _, trailhead := range trailheads {
		totalScore += getScore(&input, trailhead)

	}
	fmt.Println(totalScore)

	// part 2
	totalRating := 0
	for _, trailhead := range trailheads {
		totalRating += getRating(&input, trailhead)

	}
	fmt.Println(totalRating)
}
