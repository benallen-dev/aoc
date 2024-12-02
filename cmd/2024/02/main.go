package main

import (
	"fmt"
	"math"
)

type Direction int

const (
	DirectionDecreasing Direction = iota
	DirectionIncreasing
)

var directionName = map[Direction]string{
	DirectionDecreasing: "decreasing",
	DirectionIncreasing: "increasing",
}

func (d Direction) String() string {
	return directionName[d]
}

func checkSafe(input []int) bool {
	var dir Direction

	if input[1] > input[0] {
		dir = DirectionIncreasing
	} else if input[1] < input[0] {
		dir = DirectionDecreasing
	} else {
		return false
	}

	for i := 1; i < len(input); i++ {
		diff := input[i] - input[i-1]
		mag := math.Abs(float64(diff))
		if mag > 3 || diff == 0 || (dir == DirectionDecreasing && diff > 0) || (dir == DirectionIncreasing && diff < 0) {
			return false
		}
	}

	return true
}

func checkSafePart2(input []int) bool {
	var dir Direction

	if input[1] > input[0] {
		dir = DirectionIncreasing
	} else if input[1] < input[0] {
		dir = DirectionDecreasing
	} else {
		return false
	}

	fail := 0
	// doesn't take into account the first pair being the wrong way around
	for i := 1; i < len(input); i++ {
		diff := input[i] - input[i-1]
		mag := math.Abs(float64(diff))
		if mag > 3 || diff == 0 || (dir == DirectionDecreasing && diff > 0) || (dir == DirectionIncreasing && diff < 0) {
			fail++
		}
	}

	return true
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	// The levels are either all increasing or all decreasing.
	// Any two adjacent levels differ by at least one and at most three.
	c := 0
	d := 0
	for _, x := range input {
		if checkSafe(x) {
			c++
		}

		if checkSafePart2(x) {
			d++
		}
	}

	fmt.Printf("Part 1: %d\n", c)

	fmt.Printf("Part 2: %d\n", d)
}
