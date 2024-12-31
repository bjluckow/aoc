package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() []int {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	var nums []int
	for _, line := range lines {
		num, err := strconv.Atoi(string(line))
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func mix(a int64, b int64) int64 {
	return a ^ b // bitwise xor

}

func prune(num int64) int64 {
	return num % 16777216
}

func evolve(num int64) (next int64) {
	next = num
	next = prune(mix(next*64, next))
	next = prune(mix(next/32, next))
	next = prune(mix(next*2048, next))
	return next
}

func getSecretNum(num int64, iters int) int64 {
	current := num
	for range iters {
		current = evolve(current)
	}
	return current
}

// part 2 funcs
func getAllSecretNums(num int64, iters int) (secretNums []int64) {
	secretNums = []int64{num} // include initial num
	for range iters {
		num = evolve(num)
		secretNums = append(secretNums, num)
	}
	return secretNums
}

func getPrices(secretNums []int64) (prices []int) {
	for _, n := range secretNums {
		prices = append(prices, int(n%10))
	}
	return prices
}

func getPriceDiffs(prices []int) (diffs []int) {
	last := prices[0]
	for i, current := range prices {
		if i > 0 {
			diffs = append(diffs, current-last)
			last = current
		}
	}
	return diffs
}

func main() {
	input := getInput()
	const iters = 2000

	// part 1
	var total int64 = 0
	for _, num := range input {
		total += getSecretNum(int64(num), iters)
	}
	fmt.Println(total)

	// part 2
	/*
		want to find a sequence of 4 price *CHANGES* such that if we sell at the price
			of the 4th change, the total price across buyers is maximized
		the value for a buyer is 0 if that sequence of changes never occurs for that buyer
	*/
	type seqmap = map[[4]int]int

	seqtotals := make(seqmap)
	for _, num := range input {
		secretNums := getAllSecretNums(int64(num), iters)
		prices := getPrices(secretNums)
		diffs := getPriceDiffs(prices)
		if len(diffs) < 4 {
			panic("less than 4 diffs")
		}

		seqsells := make(seqmap)
		for i := range len(diffs) - 3 {
			seq := [4]int{diffs[i], diffs[i+1], diffs[i+2], diffs[i+3]}
			// Sell occurs *immediately* when seen
			if _, ok := seqsells[seq]; !ok {
				seqsells[seq] = prices[i+3+1]
			}
		}

		for seq, price := range seqsells {
			seqtotals[seq] += price
		}
	}

	max := 0
	for _, sum := range seqtotals {
		if sum > max {
			max = sum
		}
	}
	fmt.Println(max)
}
