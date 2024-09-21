package main

import (
	"fmt"
)

func main() {
	rounds, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	total := 0
	for _, round := range rounds {
		total += round.Score()
	}

	// Part 01, total score if everything happens as described
	// No doubt part 2 will be "oh shit they're trying to trick you" or something
	fmt.Printf("Part 1: %d\n", total)

	// Nope, we bamboozled ourselves!
	// X = lose
	// Y = draw
	// Z = win

	// The whole round datastructure is wrong, and I could have known because
	// why use ABC XYZ when they refer to the same thing?

	strats, err := readInputProperly("input.txt")
	if err != nil {
		panic(err)
	}

	total = 0
	for _, strat := range strats {
		total += strat.Score()
	}
	fmt.Printf("Part 2: %d\n", total)
}
