package main

import (
	"aoc/internal/set"
	"fmt"
)

func main() {
	lab, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	guardPositions := set.Set[position]{}
	for lab.GuardPresent() {
		guardPositions.Add(lab.guard.pos)
		lab.Tick()
	}

	fmt.Println("Part 1:", guardPositions.Length())
}

