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

	// in-place sort both lists
	slices.Sort(input[0])
	slices.Sort(input[1])

	// For each index take the difference between the two lists
	sum := 0
	for i := range input[0] {
		diff := int(math.Abs(float64(input[0][i] - input[1][i]))) // I wonder if we'll get FP weirdness
		sum += diff
	}

	fmt.Printf("Part 1: %d\n", sum)

	// For part 2 we need to take the 'right' list and count how often each element appears
	// In O(n) time we can create a map from the sorted list by iterating over list items
	counts := make(map[int]int)
	for _, v := range input[1] {
		if _, ok := counts[v]; !ok {
			counts[v] = 1
		} else {
			counts[v]++
		}
	}

	// Now we can iterate over the original left list and multiply by the map values
	sum2 := 0
	for _, l := range input[0] {
		sum2 += l * counts[l]
	}

	fmt.Printf("Part 2: %d\n", sum2)
}

