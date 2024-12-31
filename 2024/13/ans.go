package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type button struct {
	x int64
	y int64
}

type prize struct {
	x int64
	y int64
}

type game struct {
	a     button
	b     button
	prize prize
}

var games []game

const A_COST = 3
const B_COST = 1

func init() {
	input, _ := os.ReadFile("input.txt")

	parseButton :=
		func(line string) button {
			matches := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`).FindStringSubmatch(line)
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			return button{int64(x), int64(y)}
		}

	parsePrize := func(line string) prize {
		matches := regexp.MustCompile(`X=(\d+), Y=(\d+)`).FindStringSubmatch(line)
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		return prize{int64(x), int64(y)}
	}

	blocks := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	for _, block := range blocks {
		lines := strings.Split(block, "\n")
		a := parseButton(lines[0])
		b := parseButton(lines[1])
		prize := parsePrize(lines[2])
		games = append(games, game{a, b, prize})
	}
}

func isInteger(f float64) bool {
	return f == float64(int64(f))
}

/*
solve:

	u*A_X + v*B_X = P_X
	u*3A_Y + vB_Y = P_Y
*/
func getCost(g game) int64 {
	det := float64(g.a.x*g.b.y - g.b.x*g.a.y)
	if det == 0 {
		return 0 // no solution: game can't be won
	}

	// solve using cramer's rule
	u := float64(g.prize.x*g.b.y-g.b.x*g.prize.y) / det
	v := float64(g.a.x*g.prize.y-g.prize.x*g.a.y) / det

	if !isInteger(u) || !isInteger(v) {
		return 0 // non-integer solution: game can't be won
	}

	return 3*int64(u) + int64(v)
}

func main() {
	var total int64 = 0
	for _, g := range games {
		total += getCost(g)
	}
	fmt.Println(total)

	// lucked out on go's 64-bit types
	const OFFSET = int64(10000000000000)
	addOffset := func(g game) game {
		return game{g.a, g.b, prize{g.prize.x + OFFSET, g.prize.y + OFFSET}}
	}

	total = 0
	for _, g := range games {
		total += getCost(addOffset(g))
	}
	fmt.Println(total)
}
