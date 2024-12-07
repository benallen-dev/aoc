package main

import (
	"fmt"
)

func getMiddleItem(x []int) int {
	return x[len(x)/2]
}

func main() {
	rules, updates, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	part1 := 0
	part2 := 0
	for _, update := range updates {
		if update.Valid(rules) {
			part1 += getMiddleItem(update.IntSlice())
		} else {
			update.Fix(rules)

			part2 += getMiddleItem(update.IntSlice())
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

