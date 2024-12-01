package main

import (
	"fmt"
	"slices"
	"math"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	// Copy the inputs because we need them un-sorted later
	// Doing this in memory is way faster than re-reading from disk
	left := make([]int, len(input[0]))
	right := make([]int, len(input[1]))

	copy(left, input[0])
	copy(right, input[1])

	// sort both lists
	slices.Sort(left)
	slices.Sort(right)

	// For each index take the difference between the two lists
	sum := 0
	for i := range left {
		diff := int(math.Abs(float64(left[i] - right[i]))) // I wonder if we'll get FP weirdness
		sum += diff
	}

	fmt.Printf("Part 1: %d\n", sum)

	// For part 2 we need to take the second list and count how often each element appears
	// In O(n) time we can create a map from the second list

	// Because we sorted in place we need reset our left list
	// The right list is going to be summed anyway so we can leave it as is
	copy(left, input[0])

	counts := make(map[int]int)
	for _, v := range input[1] {
		if _, ok := counts[v]; !ok {
			counts[v] = 1
		} else {
			counts[v]++
		}
	}

	// Now we can iterate over the left list and multiply by the map values
	sum2 := 0
	for _, l := range input[0] {
		sum2 += l * counts[l]
	}

	fmt.Printf("Part 2: %d\n", sum2)
}

