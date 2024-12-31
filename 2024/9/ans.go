package main

import (
	"fmt"
	"os"
	"strconv"
)


func getInput() string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err) 
	}
	return string(input)
}

func getBlocks(input string) (blocks []int) {
	idCounter := 0
	for i, r := range input {
		isFile := bool(i % 2 == 0)
		n, _ := strconv.Atoi(string(r))
		if isFile {
			// fmt.Println(fmt.Sprintf("%d:%d", n, idCounter))
			for range n {
				blocks = append(blocks, idCounter)
			}
			idCounter++
		} else {
			for range n {
				blocks = append(blocks, -1)
			}	
		}
	}
	return blocks
}


func moveBlocks(blocks []int) (result []int) {
	forwardPtr := 0
	backwardPtr := len(blocks)-1

	for forwardPtr <= backwardPtr {
		if blocks[forwardPtr] >= 0 {
			result = append(result, blocks[forwardPtr])
		} else {
			for blocks[backwardPtr] < 0 {
				backwardPtr--
				if backwardPtr <= forwardPtr {
					break;
				}
			}
			result = append(result, blocks[backwardPtr])
			backwardPtr--
		}
		forwardPtr++
	}
	return result
}

func getChecksum(compact []int) (result int64) {
	for i, n := range compact {
		if n >= 0 {
			result += int64(i * n)
		}
	}
	return result
}

/*
	Iterate backwards over each window of contiguous non-negative elements ("CNNE")
	For each CNNE window, iterate forwards over each window of contiguous negative elements ("CNE")
		until the current CNNE window is reached or a CNE is found such that len(CNE) >= len(CNNE)
	If such a CNE window is found, populate it with each value from the CNNE window from left to right,
		setting each value in the CNNE window to -1
	Return the mutated array
*/

func moveWholeBlocks(blocks []int) ([]int) {
	seenIDs := make(map[int]struct{})
	br := len(blocks)-1 // Right index of "backward" window
	bl := br // Left index of backward window


	// Iterate until backwards window is out of bounds
	for bl >= 0 && br >= 0 {
		// Move right index to first element of CNNEs
		for br >= 0 && blocks[br] < 0 {
			br--
		}

		// Exit early if right index has gone out of bounds
		if br < 0 {
			break;
		}

		// Start left index at first element of CNNEs, move to last equal element
		bl=br-1
		for bl >= 0 && blocks[bl] == blocks[br] {
			bl--
		}
		bl++ // Bump left index back into the CNNE
		

		// Backward window has captured CNNEs
		// Check if this ID has already been seen (may have been moved)
		_, seenID := seenIDs[blocks[br]]
		if seenID {
			continue
		}

		// Begin trying each forward window


		fl, fr := 0, 0 // Init indices of "forward" window
		// Iterate until forward window overlaps backward window
		for fr < bl {
			// Move fl to first element of CNEs
			for fl < bl && blocks[fl] >= 0 {
				fl++
			}

			// If we've fit the backward window, there are no CNEs
			if fl >= bl {
				break
			}

			// Assuming we've hit CNEs, start right index at left index 
			fr=fl+1
			// Move right index to end of CNEs
			for fr < bl && blocks[fr] == blocks[fl] {
				fr++
			}
			fr-- // Bump index back into CNEs

			// Forward window has been captured
			// See if transferring elements is possible

			forwardLen := fr-fl+1
			backwardLen := br-bl+1
			if forwardLen >= backwardLen {
				// Move each CNNE in backward window to leftmost CNE in forward window
				// There may still be space in the CNE after this process
				for i := range backwardLen {
					blocks[fl+i] = blocks[br-i]
					blocks[br-i] = -1  // Last CNNE will be index br-(br-bl) = bl
				}

				// Backward window has been transferred; can move onto next
				break
			}

			// No suitable forward window found; reset indices to be adjusted next iter
			fl=fr+1
			fr=fl
		}

		// Mark ID as seen
		seenIDs[blocks[br]] = struct{}{}

		// Reset backwards window indices to be adjusted next iter.
		br=bl-1
		bl=br
	}
	return blocks
}

func main() {
	input := getInput()
	checksum := getChecksum(moveBlocks(getBlocks(input)))
	fmt.Println(checksum)

	checksumNonContiguous := getChecksum(moveWholeBlocks(getBlocks(input)))
	fmt.Println(checksumNonContiguous)
}