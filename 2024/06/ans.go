package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func getInput() []string {
	input, _ := os.ReadFile("input.txt")
    lines := strings.Split(strings.TrimSpace(string(input)), "\n")
    return lines
}

type position = [2]int
type positions = map[position]interface{}

func addPosition(pos position, set *positions) *positions {
	(*set)[pos] = struct{}{}
	return set
}

func getEmptyBlockedAndStart(input *[]string) (empty, blocked positions, start position) {
	empty = positions{}
	blocked = positions{}
	start = position{-1, -1}
	for y, line := range *input {
		for x, spot := range line {
		pos := position{x, y}
		switch spot {
				case '.':
					addPosition(pos, &empty)
				case '#':
					addPosition(pos, &blocked)
				case '^':
					start = pos
				}
		}
	}
	return empty, blocked, start
} 

func getSpot(pos position, input *[]string) rune {
	return rune((*input)[pos[1]][pos[0]])
}

var vectors = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
type direction = int


func rotate(dir direction) direction {
	return (dir + 1) % 4
}

func move(pos position, dir direction) position {
	vector := vectors[dir]
	return [2]int{pos[0] + vector[0], pos[1] + vector[1]}
}

func isInBounds(pos position, input *[]string) bool {
	return 0 <= pos[0] && pos[0] < len(*input) && 0 <= pos[1] && pos[1] < len((*input)[0])
}

func isBlocked(pos position, input *[]string) bool {
	return getSpot(pos, input) == '#'
}

func step(pos position, dir direction, input *[]string) (nextpos position, nextdir direction) {
	nextpos = move(pos, dir)
	nextdir = dir
	if isInBounds(nextpos, input) && isBlocked(nextpos, input) {
		nextdir = rotate(dir)
		nextpos = move(pos, nextdir)
	}
	return nextpos, nextdir
}

func getReachable(input *[]string) (reachable positions) {
	reachable = positions{}
	_, _, start := getEmptyBlockedAndStart(input)
	pos := start
	var dir direction = 0 // init to north
	for isInBounds(pos, input) {
		addPosition(pos, &reachable)
		pos, dir = step(pos, dir, input)
	}
	return
}

// part 2 funcs

func stepWithBlock(pos position, dir direction, input *[]string, block position) (nextpos position, nextdir direction) {
	nextpos = move(pos, dir)
	nextdir = dir
	if isInBounds(nextpos, input) && (isBlocked(nextpos, input) || block == nextpos )  {
		nextdir = rotate(dir)
		nextpos = move(pos, nextdir)
	}
	return nextpos, nextdir
}

type state = struct{
	pos position 
	dir direction
}

type states = map[state]interface{}

func hasCycle(input *[]string, block position) bool {
	_, _, start := getEmptyBlockedAndStart(input)
	visited := states{}

	pos := start
	dir := 0

	for isInBounds(pos, input) {
		state := state{pos, dir}
		if visited[state] != nil {
			return true
		}
		visited[state] = struct{}{}
		pos, dir = stepWithBlock(pos, dir, input, block)
	}
	return false
}


func main() {
	input := getInput()

	// part 1
	reachable := getReachable(&input)
	fmt.Println(len(reachable))

	// part 2i
	empty, _, _ := getEmptyBlockedAndStart(&input)
	jobs := make(chan position, len(empty))
	var wg sync.WaitGroup
	var cycles sync.Map
	const workers = 10

	for i := 0; i < workers; i++ {
		wg.Add(1) 
		go func() {
			defer wg.Done() 
			for pos := range jobs { 
				if hasCycle(&input, pos) {
					cycles.Store(pos, struct{}{}) 
				}
			}
		}()
	}

	go func() {
		for pos := range empty {
			jobs <- pos 
		}
		close(jobs) 
	}()

	wg.Wait()
	count := 0
	cycles.Range(func(_, _ interface{}) bool {
		count++
		return true
	})
	fmt.Println(count)
}