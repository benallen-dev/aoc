package main

import (
	"fmt"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	c := 0
	for _, v := range input {
		if v.FullOverlap() {
			c++
		}
	}

	fmt.Printf("Part 1: %d\n", c)

	c = 0
	for _, v := range input {
		if v.PartialOverlap() {
			c++
		}
	}

	fmt.Printf("Part 2: %d\n", c)
}

