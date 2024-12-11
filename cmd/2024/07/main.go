package main

import (
	"fmt"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	count := 0
	for _, eq := range input {
		if eq.Solvable() {
			count += eq.target
		}
	}

	fmt.Println("Part 1:", count)
}

