package main

import (
	"fmt"
)

func getMiddleItem(x []int) int {
	return x[len(x)/2]
}

func fixUpdate(u update, rules []rule) update {
	// TODO
	return u
}


func main() {
	rules, updates, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	// Part 1
	sum := 0
	for _, update := range updates {
		if update.Valid(rules) {
			sum += getMiddleItem(update.IntSlice())
		}
	}

	fmt.Println("Part 1:", sum)
}

