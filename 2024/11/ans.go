package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getInput() (nums []int64) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	for _, numStr := range strings.Split(strings.TrimSpace(string(input)), " ") {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, int64(num))
	}
	return nums
}

func getNumDigits(n int64) int {
	return 1 + int(math.Floor(math.Log10(float64(n))))
}

func change(n int64) (result []int64) {
	if n == 0 {
		return []int64{1}
	}

	if numDigits := getNumDigits(n); numDigits%2 == 0 {
		p := int64(math.Pow10(numDigits / 2))
		left, right := int64(n%p), int64(n/p)
		return []int64{left, right}
	}

	return []int64{n * 2024}
}

func blink(nums []int64) (result []int64) {
	for _, num := range nums {
		for _, rnum := range change(num) {
			result = append(result, rnum)
		}
	}
	return result
}

type entry struct {
	n    int64
	iter int
}

type cache = map[entry]int64

func countResults(n int64, iter, iters int, cache *cache) (count int64) {
	// Base Case
	if iter >= iters {
		return 1
	}

	if cachedValue, isCached := (*cache)[entry{n, iter}]; isCached {
		return cachedValue
	}

	// 0 -> 1
	if n == 0 {
		return countResults(1, iter+1, iters, cache)
	}

	if numDigits := getNumDigits(n); numDigits%2 == 0 {
		p := int64(math.Pow10(numDigits / 2))
		left, right := int64(n%p), int64(n/p)
		count = countResults(left, iter+1, iters, cache) + countResults(right, iter+1, iters, cache)
	} else {
		count = countResults(int64(n*2024), iter+1, iters, cache)
	}

	(*cache)[entry{n, iter}] = count
	return count
}

func main() {
	input := getInput()

	stones := input
	for range 25 {
		stones = blink(stones)
	}
	fmt.Println(len(stones))

	// part 2
	cache := make(cache)
	total := int64(0)
	for _, n := range input {
		total += countResults(n, 0, 75, &cache)
	}
	fmt.Println(total)
}
