package main

import (
	"fmt"
	"math"
)

type Direction int

const (
	DirUp Direction = iota
	DirDown
)

var directionName = map[Direction]string{
	DirUp:   "decreasing",
	DirDown: "increasing",
}

func (d Direction) String() string {
	return directionName[d]
}

func checkSafe(input []int) bool {
	var dir Direction

	if input[0] > input[1] {
		dir = DirDown
	} else if input[0] < input[1] {
		dir = DirUp
	} else {
		return false
	}

	for i := 1; i < len(input); i++ {
		diff := input[i] - input[i-1]
		mag := math.Abs(float64(diff))
		if mag > 3 || diff == 0 || (dir == DirUp && diff < 0) || (dir == DirDown && diff > 0) {
			return false
		}
	}

	return true
}

// We need to make copies to avoid modifying the input slice - the underlying array is shared
func removeElement(input []int, index int) []int {
	result := make([]int, len(input)-1)
	copy(result, input[:index])
	copy(result[index:], input[index+1:])

	return result
}

func checkSafePart2(input []int) bool {
	// The idea behind part 2 is that one of the levels can be removed to
	// create a valid report - this could be any of the levels

	// The simple implementation for this is to remove each level in turn and
	// check if the report is valid. You could try to include this in the main
	// loop by removing the ith element if the pair of i and i+1 are invalid
	// and then rechecking, but you need edge cases to handle the first and
	// last elements and this is day 2, so brute force it is.
	for i := 0; i < len(input); i++ {
		x := removeElement(input, i)
		if checkSafe(x) {
			return true
		}
	}

	return false
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	c := 0
	d := 0
	for _, x := range input {
		if checkSafe(x) {
			c++
			d++
		} else if checkSafePart2(x) {
			d++
		}
	}

	fmt.Printf("Part 1: %d\n", c)
	fmt.Printf("Part 2: %d\n", d)
}
